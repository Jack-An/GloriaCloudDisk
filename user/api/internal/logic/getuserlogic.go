package logic

import (
	"GloriaCloudDisk/common"
	"GloriaCloudDisk/user/rpc/user"
	"context"
	"fmt"
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
	userId, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, common.NewCodeError(common.INVALID_ARGUMENT, "id not valid")
	}

	jwtTokenUserId := l.ctx.Value("userId")
	if fmt.Sprintf("%v", jwtTokenUserId) != req.Id {
		return nil, common.NewDefaultMgsError(common.PERMISSION_DENIED)
	}

	r, err := l.svcCtx.User.GetUser(l.ctx, &user.GetUserReq{Id: int64(userId)})
	if err != nil {
		return nil, common.NewDefaultMgsError(common.NOT_FOUND)
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
