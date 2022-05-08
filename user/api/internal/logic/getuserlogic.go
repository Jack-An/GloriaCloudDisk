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

func makeGetUserFailedResp(message string) *types.GetUserResp {
	res := &types.GetUserResp{}
	res.Code = 1002
	res.Err = message
	return res
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserResp, err error) {
	userId, err := strconv.Atoi(req.Id)
	if err != nil {
		return makeGetUserFailedResp("id not valid"), nil
	}

	jwtTokenUserId := l.ctx.Value("userId")
	if jwtTokenUserId != userId {
		return makeGetUserFailedResp("unauthorized"), nil
	}

	r, err := l.svcCtx.User.GetUser(l.ctx, &user.GetUserReq{Id: int64(userId)})
	if err != nil {
		return makeGetUserFailedResp("user not found"), nil
	}
	res := &types.GetUserResp{Data: types.UserInfo{
		Id:        r.Id,
		Name:      r.Name,
		Email:     r.Email,
		Phone:     r.Phone,
		Active:    r.Active,
		Source:    r.Source,
		CreatedAt: r.CreatedAt}}
	return res, nil
}
