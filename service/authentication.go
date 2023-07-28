package service

import (
	"github.com/AndiGanesha/authentication/application"
	"github.com/AndiGanesha/authentication/model"
	"github.com/AndiGanesha/authentication/repository"
)

// define interface
type IAuthService interface {
	SignUp(model.User) (token string, err error)
	SignIn(model.User) (token string, err error)
}

// define a scallable struct if needed in the future
type AuthService struct {
	authRepo repository.IAuthenticationRepository
}

// create stock service func
func NewAuthenticationService(app *application.App) IAuthService {
	return &AuthService{
		authRepo: repository.NewAuthenticationRepository(app),
	}
}

func (s *AuthService) SignUp(model.User) (token string, err error) {
	return "", nil
}

func (s *AuthService) SignIn(model.User) (token string, err error) {
	return "", nil
}
