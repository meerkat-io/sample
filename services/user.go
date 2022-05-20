package services

import (
	"context"

	"sample/dtos"
	"sample/entities"
	"sample/repositories"
)

type UserService interface {
	CreateUser(ctx context.Context, user *dtos.UserData) error
	ReadUser(ctx context.Context, userId string) (*dtos.UserData, error)
	UpdateUser(ctx context.Context, user *dtos.UserData) error
	DeleteUser(ctx context.Context, userId string) error
}

type userService struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) CreateUser(ctx context.Context, user *dtos.UserData) error {
	return u.repo.Create(ctx, &entities.User{
		UserId:    user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
	})
}

func (u *userService) ReadUser(ctx context.Context, userId string) (*dtos.UserData, error) {
	user, err := u.repo.Read(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &dtos.UserData{
		UserId:    user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
	}, nil
}

func (u *userService) UpdateUser(ctx context.Context, user *dtos.UserData) error {
	return u.repo.Update(ctx, &entities.User{
		UserId:    user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
	})
}

func (u *userService) DeleteUser(ctx context.Context, userId string) error {
	return u.repo.Delele(ctx, userId)
}
