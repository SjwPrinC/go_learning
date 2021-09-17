package biz

import "context"

type UserService struct {
	ur UserRepo
}

type UserRepo interface {
	Register(ctx context.Context) error

	ChangeUserName(ctx context.Context) error
}

func (us *UserService) Register() {
	us.ur.Register(context.Background())
}

func (us *UserService) ChangeUserName() {
	us.ur.Register(context.Background())
}

func NewUserService() {

}
