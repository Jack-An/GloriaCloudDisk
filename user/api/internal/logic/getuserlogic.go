package logic

import (
	"GloriaCloudDisk/user/rpc/user"
	"context"
	"strconv"

	"GloriaCloudDisk/user/api/internal/svc"
	"GloriaCloudDisk/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserResp, err error) {
	// todo: add your logic here and delete this line
	userId, err := strconv.Atoi(req.Id)
	if err != nil {
		res := &types.GetUserResp{}
		res.Code = 400
		res.Err = "id not valid"
		return res, nil
	}
	r, err := l.svcCtx.User.GetUser(l.ctx, &user.GetUserReq{Id: int64(userId)})
	if err != nil {
		res := &types.GetUserResp{}
		res.Code = 404
		res.Err = "not found"
		return res, nil
	}
	res := &types.GetUserResp{Data: types.UserInfo{
		Id:        r.Id,
		Name:      r.Name,
		Email:     r.Email,
		Phone:     r.Phone,
		Active:    r.Active,
		Source:    r.Source,
		CreatedAt: r.CreatedAt}}
	res.Code = 0
	res.Err = ""
	return res, nil
}
