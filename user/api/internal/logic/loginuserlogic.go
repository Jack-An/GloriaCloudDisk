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

func (l *LoginUserLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginUserLogic) LoginUser(req *types.LoginReq) (resp *types.LoginResp, err error) {
	var res *user.GetUserResp
	var errno error

	if common.VerifyMobileFormat(req.Identity) {
		res, errno = l.svcCtx.User.GetUserByPhone(l.ctx, &user.GetByPhoneReq{Phone: req.Identity})
	} else if common.VerifyEmailFormat(req.Identity) {
		res, errno = l.svcCtx.User.GetUserByEmail(l.ctx, &user.GetByEmailReq{Email: req.Identity})
	} else {
		return nil, common.NewCodeError(common.INVALID_ARGUMENT, "identity not match phone or email")
	}

	if errno != nil {
		return nil, common.NewDefaultMgsError(common.NOT_FOUND)
	}

	verifyResult, errno := l.svcCtx.User.VerifyPassword(l.ctx, &user.VerifyReq{Id: res.Id, Password: req.Password})
	if errno != nil {
		logx.Errorf("verify fail: %s", err)
		return nil, common.NewCodeError(common.UNKNOWN, "verify fail")
	}
	if verifyResult.Ok {
		now := time.Now().Unix()
		accessExpire := l.svcCtx.Config.Auth.AccessExpire
		jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, res.Id)
		if err != nil {
			return nil, common.NewCodeError(common.UNKNOWN, "generate token fail")
		}
		return &types.LoginResp{Data: types.LoginClaims{
			Id:           res.Id,
			Name:         res.Name,
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		}}, nil
	} else {
		return nil, common.NewCodeError(common.INVALID_ARGUMENT, "password is wrong")
	}
}
