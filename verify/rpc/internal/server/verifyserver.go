// Code generated by goctl. DO NOT EDIT!
// Source: verify.proto

package server

import (
	"context"

	"GloriaCloudDisk/verify/rpc/internal/logic"
	"GloriaCloudDisk/verify/rpc/internal/svc"
	"GloriaCloudDisk/verify/rpc/verify"
)

type VerifyServer struct {
	svcCtx *svc.ServiceContext
	verify.UnimplementedVerifyServer
}

func NewVerifyServer(svcCtx *svc.ServiceContext) *VerifyServer {
	return &VerifyServer{
		svcCtx: svcCtx,
	}
}

func (s *VerifyServer) SendEmailCode(ctx context.Context, in *verify.SendEmailCodeReq) (*verify.SendEmailCodeResp, error) {
	l := logic.NewSendEmailCodeLogic(ctx, s.svcCtx)
	return l.SendEmailCode(in)
}
