package services_test

import (
	"context"
	"sample/dtos"
	"sample/repositories"
	"sample/services"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	ctx := context.Background()
	ctl := gomock.NewController(t)
	userRepo := repositories.NewMockUserRepo(ctl)
	userRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	userService := services.NewUserService(userRepo)
	err := userService.CreateUser(ctx, &dtos.UserData{
		UserId: "test_user_0",
	})
	assert.Nil(t, err)
}
