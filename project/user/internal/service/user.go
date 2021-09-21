package service

import (
	"context"
	"log"

	"github.com/sjw/user/do"
	"github.com/sjw/user/internal/biz"
	napi "github.com/sjw/user/newapi/v1"
)

type UserService struct {
	ub *biz.UserUseCase
}

func (us *UserService) Register(ctx *context.Context, rr *napi.RegisterRequest) (*napi.CommonReply, error) {
	err := us.ub.Register(&do.RegisterDo{
		UserName: rr.Name,
		Password: rr.Password,
		Mobile:   rr.Mobile,
	})

	if err != nil {
		log.Fatal(err)
	} else {
		return &napi.CommonReply{
			Success: true,
			Msg:     "resister success",
		}, nil
	}

	return nil, nil
}

func (us *UserService) ChangeUserName(ctx context.Context, rr *napi.ChangeUserNameRequest) (*napi.CommonReply, error) {
	err := us.ub.ChangeUserName(&do.ChangeUserNameDo{
		UserName: rr.Uid,
		NewName:  rr.NewName,
	})

	if err != nil {
		log.Fatal(err)
	} else {
		return &napi.CommonReply{
			Success: true,
			Msg:     "resister success",
		}, nil
	}

	return nil, nil
}

func NewUserService(userUseCase *biz.UserUseCase) *UserService {
	return &UserService{ub: userUseCase}
}
