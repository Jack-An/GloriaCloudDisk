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

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {

	if len(req.Phone) == 0 && len(req.Email) == 0 {
		res := &types.CreateUserResp{}
		res.Code = 400
		res.Err = "phone not email must have one"
		return res, nil
	}

	_, err1 := l.svcCtx.User.CreateUser(l.ctx, &user.CreateUserReq{
		Password: req.Password,
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Source:   req.Source})

	if err1 != nil {
		res := &types.CreateUserResp{}
		res.Code = 500
		res.Err = err1.Error()
		return res, nil
	}

	return &types.CreateUserResp{}, nil
}
