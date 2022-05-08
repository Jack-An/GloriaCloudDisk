package logic

import (
	"GloriaCloudDisk/common"
	"GloriaCloudDisk/user/rpc/user"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"GloriaCloudDisk/user/api/internal/svc"
	"GloriaCloudDisk/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func makeLoginFailedResp(message string) *types.LoginResp {
	resp := types.LoginResp{}
	resp.Code = 1000
	resp.Err = message
	return &resp
}

func (l *LoginUserLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func makeLoginSuccessResp(l *LoginUserLogic, userResp *user.GetUserResp) *types.LoginResp {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userResp.Id)
	if err != nil {
		return makeLoginFailedResp("generate jwt token fail")
	}
	resp := types.LoginResp{Data: types.LoginClaims{
		Id:           userResp.Id,
		Name:         userResp.Name,
		AccessToken:  jwtToken,
		AccessExpire: accessExpire,
		RefreshAfter: accessExpire / 2,
	}}
	return &resp
}

func (l *LoginUserLogic) LoginUser(req *types.LoginReq) (resp *types.LoginResp, err error) {
	var res *user.GetUserResp
	var errno error

	if common.VerifyMobileFormat(req.Identity) {
		res, errno = l.svcCtx.User.GetUserByPhone(l.ctx, &user.GetByPhoneReq{Phone: req.Identity})
	} else if common.VerifyEmailFormat(req.Identity) {
		res, errno = l.svcCtx.User.GetUserByEmail(l.ctx, &user.GetByEmailReq{Email: req.Identity})
	} else {
		return makeLoginFailedResp("identity not match phone or email"), nil
	}

	if errno != nil {
		return makeLoginFailedResp("user not found"), nil
	}

	verifyResult, errno := l.svcCtx.User.VerifyPassword(l.ctx, &user.VerifyReq{Id: res.Id, Password: req.Password})
	if errno != nil {
		return makeLoginFailedResp(errno.Error()), nil
	}
	if verifyResult.Ok {
		return makeLoginSuccessResp(l, res), nil
	} else {
		return makeLoginFailedResp("password is wrong"), nil
	}
}
