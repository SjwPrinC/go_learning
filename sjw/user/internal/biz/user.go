package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

type UserRepo interface {
	Register(ctx context.Context, user *User)

	Login(ctx context.Context, user *User)
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUseCase) Register(ctx context.Context, user *User) error {
	uc.repo.Register(ctx, user)

	return nil
}

func (uc *UserUseCase) Login(ctx context.Context, user *User) error {
	uc.repo.Login(ctx, user)

	return nil
}

type User struct {
	UserId        int64
	UserName      string
	Mobile        string
	Password      string
	Email         string
	CreateTime    time.Time
	LastVisitTime time.Time
	State         int8
	Is_Vip        bool
}
