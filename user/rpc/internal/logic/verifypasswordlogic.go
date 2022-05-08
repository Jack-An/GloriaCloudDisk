package logic

import (
	"GloriaCloudDisk/common"
	"context"

	"GloriaCloudDisk/user/rpc/internal/svc"
	"GloriaCloudDisk/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyPasswordLogic {
	return &VerifyPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyPasswordLogic) VerifyPassword(in *user.VerifyReq) (*user.VerifyResp, error) {
	identity, err := l.svcCtx.IdentityModel.FindOne(l.ctx, in.Id)

	encrypt := common.CryptContext{Schema: l.svcCtx.Config.EncryptSchema}

	if err != nil {
		return &user.VerifyResp{Ok: false}, err
	}

	status, err := encrypt.Verify(in.Password, identity.Password)
	if err != nil {
		return &user.VerifyResp{Ok: false}, err
	}

	return &user.VerifyResp{Ok: status}, nil
}
