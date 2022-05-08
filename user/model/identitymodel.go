package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ IdentityModel = (*customIdentityModel)(nil)

type (
	// IdentityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIdentityModel.
	IdentityModel interface {
		identityModel
	}

	customIdentityModel struct {
		*defaultIdentityModel
	}
)

// NewIdentityModel returns a model for the database table.
func NewIdentityModel(conn sqlx.SqlConn) IdentityModel {
	return &customIdentityModel{
		defaultIdentityModel: newIdentityModel(conn),
	}
}
