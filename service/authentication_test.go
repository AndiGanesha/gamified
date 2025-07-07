package service

import (
	"testing"

	"github.com/AndiGanesha/gamified/mock"
	"github.com/AndiGanesha/gamified/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockIAuthenticationRepository(ctrl)
	// Prepare test data
	data := model.User{
		Id:       1,
		Username: "asd",
		Password: "asd",
		Phone:    "+62222222",
	}

	service := &AuthService{
		authRepo: mockRepo,
	}

	mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)

	// Call the function being tested
	err := service.CreateUser(data)

	// Assert the results
	assert.Nil(t, err)
}

func TestVerifyUserFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockIAuthenticationRepository(ctrl)
	// Prepare test data
	data := model.User{
		Id:       1,
		Username: "asd",
		Password: "asd",
		Phone:    "+62222222",
	}

	data2 := model.User{
		Id:       2,
		Username: "asd",
		Password: "asdffffff",
		Phone:    "+62222222",
	}

	service := &AuthService{
		authRepo: mockRepo,
	}

	t.Run("should OK", func(t *testing.T) {
		mockRepo.EXPECT().GetUser(gomock.Any()).Return(data, nil)

		// Call the function being tested
		val, err := service.VerifyUserFromDB(data)

		// Assert the results
		assert.Nil(t, err)
		assert.Equal(t, true, val)
	})

	t.Run("password wrong", func(t *testing.T) {
		mockRepo.EXPECT().GetUser(gomock.Any()).Return(data2, nil)

		// Call the function being tested
		val, err := service.VerifyUserFromDB(data)

		// Assert the results
		assert.Error(t, err)
		assert.Equal(t, true, val)
	})
}
