package svc

import (
	"GloriaCloudDisk/user/api/internal/config"
	"GloriaCloudDisk/user/rpc/user"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	User   user.User
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   user.NewUser(zrpc.MustNewClient(c.User)),
		Redis:  c.Redis.NewRedis(),
	}
}
