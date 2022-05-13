package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ResourceRelationModel = (*customResourceRelationModel)(nil)

type (
	// ResourceRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customResourceRelationModel.
	ResourceRelationModel interface {
		resourceRelationModel
	}

	customResourceRelationModel struct {
		*defaultResourceRelationModel
	}
)

// NewResourceRelationModel returns a model for the database table.
func NewResourceRelationModel(conn sqlx.SqlConn) ResourceRelationModel {
	return &customResourceRelationModel{
		defaultResourceRelationModel: newResourceRelationModel(conn),
	}
}
