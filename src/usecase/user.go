package usecase

import (
	"context"
	"fmt"

	"github.com/fajrulaulia/minder/helper"
	"github.com/fajrulaulia/minder/src/model"
	"github.com/fajrulaulia/minder/src/repository"
	User "github.com/fajrulaulia/minder/src/usecase/user"
)

type UserUsecaseIface interface {
	Signup(ctx context.Context, params *User.UserRequest) (string, error)
	Login(ctx context.Context, params *User.LoginRequest) (string, error)
}

type UserUsecaseStruct struct {
	UserRepo repository.UserRepositoryIface
}

func NewUserUsecase(user repository.UserRepositoryIface) UserUsecaseIface {
	return &UserUsecaseStruct{
		UserRepo: user,
	}
}

func (c *UserUsecaseStruct) Signup(ctx context.Context, params *User.UserRequest) (string, error) {
	var (
		err         error
		passwordStr string
	)

	if len(params.Password) == 0 {
		return "", fmt.Errorf("password empty")
	}

	if len(params.Emal) == 0 {
		return "", fmt.Errorf("email empty")
	}

	if len(params.Username) == 0 {
		return "", fmt.Errorf("email username")
	}

	if passwordStr, err = helper.HashPassword(params.Password); err != nil {
		return "", err
	}

	if err = c.UserRepo.CreateUser(ctx, model.User{
		Username: params.Username,
		Email:    params.Emal,
		Password: passwordStr,
	}); err != nil {
		return "", err
	}

	token, err := helper.CreateToken(&helper.User{
		Username: params.Username,
		Email:    params.Emal,
	})

	if err != nil {
		return "", err
	}

	return token, nil
}

func (c *UserUsecaseStruct) Login(ctx context.Context, params *User.LoginRequest) (string, error) {
	var (
		err error
	)

	data, err := c.UserRepo.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return "", err
	}

	if data == nil {
		return "", fmt.Errorf("username or password wrong")
	}

	err = helper.CheckPassword(params.Password, data.Password)
	if err != nil {
		return "", fmt.Errorf("username or password wrong")
	}

	token, err := helper.CreateToken(&helper.User{
		Email:    params.Email,
		Username: data.Username,
	})

	if err != nil {
		return "", err
	}

	return token, nil
}
