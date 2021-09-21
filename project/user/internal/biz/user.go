package biz

import (
	"context"

	do "github.com/sjw/user/do"
	"github.com/sjw/user/model"
)

type UserUseCase struct {
	ur UserRepo
}

type UserRepo interface {
	Register(ctx context.Context, user *model.User) error

	ChangeUserName(ctx context.Context, userName string, newName string) error
}

func (ub *UserUseCase) Register(do *do.RegisterDo) error {
	ub.ur.Register(context.Background(),
		&model.User{
			User_Name: do.UserName,
			Mobile:    do.Mobile,
			Password:  do.Password,
		})
	return nil
}

func (ub *UserUseCase) ChangeUserName(do *do.ChangeUserNameDo) error {

	ub.ur.ChangeUserName(context.Background(), do.UserName, do.NewName)

	return nil
}

func NewUserUseCase(ur UserRepo) *UserUseCase {
	return &UserUseCase{
		ur: ur,
	}
}
