package logic

import (
	"context"

	"GloriaCloudDisk/verify/rpc/internal/svc"
	"GloriaCloudDisk/verify/rpc/verify"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailCodeLogic {
	return &SendEmailCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailCodeLogic) SendEmailCode(in *verify.SendEmailCodeReq) (*verify.SendEmailCodeResp, error) {
	// todo: add your logic here and delete this line

	return &verify.SendEmailCodeResp{}, nil
}
