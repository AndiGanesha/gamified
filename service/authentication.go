package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AndiGanesha/authentication/application"
	"github.com/AndiGanesha/authentication/configuration"
	"github.com/AndiGanesha/authentication/model"
	"github.com/AndiGanesha/authentication/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
)

// define interface
type IAuthService interface {
	VerifyUserFromDB(model.User) (bool, error)
	CreateUser(model.User) error
	GenerateToken(user string, pass string) (string, error)
	SetRedisToken(token string, user model.User) error
}

// define a scallable struct if needed in the future
type AuthService struct {
	config   *configuration.Configuration
	authRepo repository.IAuthenticationRepository
	redis    redis.Client
	context  context.Context
}

// create stock service func
func NewAuthenticationService(app *application.App) IAuthService {
	return &AuthService{
		authRepo: repository.NewAuthenticationRepository(app),
		redis:    app.Redis,
		config:   app.Configuration,
		context:  app.Context,
	}
}

func (s *AuthService) CreateUser(user model.User) error {
	return s.authRepo.CreateUser(user)
}

func (s *AuthService) VerifyUserFromDB(user model.User) (bool, error) {
	userDB, err := s.authRepo.GetUser(user.Username)
	if err != nil {
		return false, err
	}

	if userDB.Username == "" {
		return false, nil
	}

	if user.Password != userDB.Password {
		return true, errors.New("user or password wrong")
	}

	return true, nil
}

func (s *AuthService) GenerateToken(user string, pass string) (string, error) {
	token, err := s.createToken(user, pass)
	if err != nil {
		return "", err
	}

	verifiedToken, err := verifyToken(token, pass)
	if err != nil {
		return "", err
	}

	return verifiedToken.Raw, nil
}

func (s *AuthService) SetRedisToken(token string, user model.User) error {
	err := s.redis.Set(s.context, token, user, time.Duration(s.config.Redis.ExpiryTime)*time.Second).Err()
	if err != nil {
		fmt.Printf("Failed to set key '%s' in Redis: %v\n", token, err)
		return err
	}
	return nil
}

// Generate the token
func (s *AuthService) createToken(username, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Duration(s.config.Token.AuthExpiry)).Unix(),
	})

	// Sign the token with the password as the secret key
	return token.SignedString([]byte(password))
}

// Verify and parse the token
func verifyToken(tokenString, password string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		// Use the password as the secret key to verify the token
		return []byte(password), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
