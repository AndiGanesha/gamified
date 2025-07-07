package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AndiGanesha/gamified/application"
	"github.com/AndiGanesha/gamified/model"
)

// define interface
type IAuthenticationRepository interface {
	GetUser(string) (model.User, error)
	CreateUser(user model.User) error
}

// define a scallable struct if needed in the future
type AuthenticationRepository struct {
	DB *sql.DB
}

// create stock service func
func NewAuthenticationRepository(app *application.App) IAuthenticationRepository {
	return &AuthenticationRepository{
		DB: app.DB,
	}
}

func (r *AuthenticationRepository) GetUser(username string) (model.User, error) {
	var user model.User

	rows := r.DB.QueryRow("SELECT id, user, pass, experience FROM user WHERE user = ?", username)
	if rows.Err() != nil {
		return user, fmt.Errorf("user by username %q: %v", username, rows.Err())
	}

	if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Experience); err != nil {
		log.Println(fmt.Errorf("error scan user by username %q: %v", username, rows.Err()))
		return user, nil
	}

	return user, nil
}

func (r *AuthenticationRepository) CreateUser(user model.User) error {
	query := `
		INSERT INTO user (user, pass, phone, experience)
		VALUES (?, ?, ?, 0)
	`
	_, err := r.DB.Exec(query, user.Username, user.Password, user.Username)
	if err != nil {
		return err
	}

	return nil
}
