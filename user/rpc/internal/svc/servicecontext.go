package svc

import (
	"GloriaCloudDisk/user/model"
	"GloriaCloudDisk/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	IdentityModel model.IdentityModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(sqlx.NewMysql(c.DataSource)),
		IdentityModel: model.NewIdentityModel(sqlx.NewMysql(c.DataSource)),
	}
}
