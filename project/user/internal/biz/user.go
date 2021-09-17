package biz

import "context"

type UserBiz struct {
	ur UserRepo
}

type UserRepo interface {
	Register(ctx context.Context) error

	ChangeUserName(ctx context.Context) error
}

func (ub *UserBiz) Register() {
	ub.ur.Register(context.Background())
}

func (ub *UserBiz) ChangeUserName() {
	ub.ur.ChangeUserName(context.Background())
}

func NewUserBiz() {

}
