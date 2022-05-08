package logic

import (
	"context"

	"GloriaCloudDisk/user/rpc/internal/svc"
	"GloriaCloudDisk/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByEmailLogic {
	return &GetUserByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByEmailLogic) GetUserByEmail(in *user.GetByEmailReq) (*user.GetUserResp, error) {
	find, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return &user.GetUserResp{}, nil
	}
	return &user.GetUserResp{
		Id:        find.Id,
		Name:      find.Name,
		Phone:     find.Phone.String,
		Email:     find.Email.String,
		Active:    find.Active != 0,
		Source:    find.Source,
		CreatedAt: find.CreatedAt.String(),
	}, nil
}
