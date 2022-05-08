package logic

import (
	"GloriaCloudDisk/common"
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"GloriaCloudDisk/user/rpc/internal/svc"
	"GloriaCloudDisk/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserReq) (*user.CreateUserResp, error) {
	conn := sqlx.NewMysql(l.svcCtx.Config.DataSource)
	encrypt := common.CryptContext{Schema: l.svcCtx.Config.EncryptSchema}
	password, errno := encrypt.Encrypt(in.Password)
	if errno != nil {
		logx.Errorf("encrypt password fail: %s", errno)
		return &user.CreateUserResp{Ok: false}, errno
	}

	userSql := `insert into user(name, phone, email, source) values (?, ?, ?, ?)`
	identitySql := `insert into identity(id, password) values (?, ?)`
	err := conn.Transact(func(session sqlx.Session) error {
		stmt, err := conn.Prepare(userSql)
		stmt2, err2 := conn.Prepare(identitySql)
		if err != nil || err2 != nil {
			return err
		}
		defer stmt.Close()
		defer stmt2.Close()

		res, err := stmt.Exec(in.Name, in.Phone, in.Email, in.Source)
		if err != nil {
			logx.Errorf("insert user stmt exec: %s", err)
			return err
		}

		userId, _ := res.LastInsertId()

		if _, err := stmt2.Exec(userId, password); err != nil {
			logx.Errorf("insert identity stmt exec: %s", err)
			return err
		}

		return nil
	})

	if err != nil {
		return &user.CreateUserResp{Ok: false}, err
	}
	return &user.CreateUserResp{Ok: true}, nil
}
