package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/sjw/user/api/v1"
	"github.com/sjw/user/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(useCase *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  useCase,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{
			Message: "注册成功",
		}, s.uc.Register(ctx, &biz.User{
			UserName: req.Name,
			Password: req.Password,
			Mobile:   req.Mobile,
		})
}
func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	return &pb.LoginReply{
			Message: "登录成功",
		}, s.uc.Login(ctx, &biz.User{
			Mobile:   req.Mobile,
			Password: req.Password,
		})
}
