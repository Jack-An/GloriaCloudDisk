package logic

import (
	"GloriaCloudDisk/common"
	"GloriaCloudDisk/user/rpc/user"
	"context"

	"GloriaCloudDisk/user/api/internal/svc"
	"GloriaCloudDisk/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {
	err = l.svcCtx.Redis.SetexCtx(l.ctx, "test", "value", 60)
	if err != nil {
		return nil, common.NewCodeError(common.UNKNOWN, err.Error())
	}
	switch req.Source {
	case "Phone":
		{
			if len(req.Phone) == 0 {
				return nil, common.NewCodeError(common.INVALID_ARGUMENT, "Phone cannot be null")
			}
			res, _ := l.svcCtx.User.GetUserByPhone(l.ctx, &user.GetByPhoneReq{Phone: req.Phone})
			if res.Id != 0 {
				return nil, common.NewDefaultMgsError(common.ALREADY_EXISTS)
			}
		}
	case "Email":
		{
			if len(req.Email) == 0 {
				return nil, common.NewCodeError(common.INVALID_ARGUMENT, "Email cannot be null")
			}
			res, _ := l.svcCtx.User.GetUserByEmail(l.ctx, &user.GetByEmailReq{Email: req.Email})
			if res.Id != 0 {
				return nil, common.NewDefaultMgsError(common.ALREADY_EXISTS)
			}
		}
	default:
		return nil, common.NewCodeError(common.INVALID_ARGUMENT, "Only support Phone or Email")

	}

	_, errno := l.svcCtx.User.CreateUser(l.ctx, &user.CreateUserReq{
		Password: req.Password,
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Source:   req.Source})

	if errno != nil {
		logx.Errorf("create user fail: %s", errno)
		return nil, common.NewDefaultMgsError(common.UNKNOWN)
	}

	return &types.CreateUserResp{}, nil
}
