package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UploadModel = (*customUploadModel)(nil)

type (
	// UploadModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUploadModel.
	UploadModel interface {
		uploadModel
	}

	customUploadModel struct {
		*defaultUploadModel
	}
)

// NewUploadModel returns a model for the database table.
func NewUploadModel(conn sqlx.SqlConn) UploadModel {
	return &customUploadModel{
		defaultUploadModel: newUploadModel(conn),
	}
}
