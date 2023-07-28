package repository

import (
	"github.com/AndiGanesha/authentication/application"
	"github.com/AndiGanesha/authentication/configuration"
	"github.com/AndiGanesha/authentication/model"
)

// define interface
type IAuthenticationRepository interface {
	Verify(model.User) (bool, error)
}

// define a scallable struct if needed in the future
type AuthenticationRepository struct {
	Config *configuration.Configuration
}

// create stock service func
func NewAuthenticationRepository(app *application.App) IAuthenticationRepository {
	return &AuthenticationRepository{
		Config: app.Configuration,
	}
}

func (r *AuthenticationRepository) Verify(model.User) (bool, error) {
	return true, nil
}
