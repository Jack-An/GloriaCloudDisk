package logic

import (
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

func makeExistsResp() *types.CreateUserResp {
	res := &types.CreateUserResp{}
	res.Code = 401
	res.Err = "already exist user"
	return res
}

func makeCreateParamsNotValidResp() *types.CreateUserResp {
	res := &types.CreateUserResp{}
	res.Code = 401
	res.Err = "phone not email must have one"
	return res
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {

	switch req.Source {
	case "Phone":
		{
			if len(req.Phone) == 0 {
				return makeCreateParamsNotValidResp(), nil
			}
			res, _ := l.svcCtx.User.GetUserByPhone(l.ctx, &user.GetByPhoneReq{Phone: req.Phone})
			if res.Id != 0 {
				return makeExistsResp(), nil
			}
		}
	case "Email":
		{
			if len(req.Email) == 0 {
				return makeCreateParamsNotValidResp(), nil
			}
			res, _ := l.svcCtx.User.GetUserByEmail(l.ctx, &user.GetByEmailReq{Email: req.Email})
			if res.Id != 0 {
				return makeExistsResp(), nil
			}
		}
	default:
		return makeCreateParamsNotValidResp(), nil

	}

	_, errno := l.svcCtx.User.CreateUser(l.ctx, &user.CreateUserReq{
		Password: req.Password,
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Source:   req.Source})

	if errno != nil {
		logx.Errorf("create user fail: %s", errno)
		res := &types.CreateUserResp{}
		res.Code = 500
		res.Err = "create fail"
		return res, nil
	}

	return &types.CreateUserResp{}, nil
}
