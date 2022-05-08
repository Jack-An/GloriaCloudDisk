package logic

import (
	"context"

	"GloriaCloudDisk/user/rpc/internal/svc"
	"GloriaCloudDisk/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	r, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &user.GetUserResp{Name: r.Name,
		Email:     r.Email.String,
		Phone:     r.Phone.String,
		Active:    r.Active != 0,
		Source:    r.Source,
		CreatedAt: r.CreatedAt.String(),
	}, nil
}
