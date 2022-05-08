// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"GloriaCloudDisk/user/rpc/internal/logic"
	"GloriaCloudDisk/user/rpc/internal/svc"
	"GloriaCloudDisk/user/rpc/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.GetUserReq) (*user.GetUserResp, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

func (s *UserServer) CreateUser(ctx context.Context, in *user.CreateUserReq) (*user.CreateUserResp, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *UserServer) GetUserByPhone(ctx context.Context, in *user.GetByPhoneReq) (*user.GetUserResp, error) {
	l := logic.NewGetUserByPhoneLogic(ctx, s.svcCtx)
	return l.GetUserByPhone(in)
}

func (s *UserServer) GetUserByEmail(ctx context.Context, in *user.GetByEmailReq) (*user.GetUserResp, error) {
	l := logic.NewGetUserByEmailLogic(ctx, s.svcCtx)
	return l.GetUserByEmail(in)
}
