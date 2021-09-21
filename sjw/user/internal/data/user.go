package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/sjw/user/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (ur *userRepo) Register(ctx context.Context, user *biz.User) {
	//todo:
}

func (ur *userRepo) Login(ctx context.Context, user *biz.User) {
	//todo:
}
