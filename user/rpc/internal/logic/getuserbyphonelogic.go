package logic

import (
	"GloriaCloudDisk/user/rpc/internal/svc"
	"GloriaCloudDisk/user/rpc/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByPhoneLogic {
	return &GetUserByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByPhoneLogic) GetUserByPhone(in *user.GetByPhoneReq) (*user.GetUserResp, error) {
	find, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
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
