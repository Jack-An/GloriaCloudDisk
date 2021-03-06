// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	resourceRelationFieldNames          = builder.RawFieldNames(&ResourceRelation{})
	resourceRelationRows                = strings.Join(resourceRelationFieldNames, ",")
	resourceRelationRowsExpectAutoSet   = strings.Join(stringx.Remove(resourceRelationFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	resourceRelationRowsWithPlaceHolder = strings.Join(stringx.Remove(resourceRelationFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	resourceRelationModel interface {
		Insert(ctx context.Context, data *ResourceRelation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ResourceRelation, error)
		Update(ctx context.Context, data *ResourceRelation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultResourceRelationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ResourceRelation struct {
		Id        int64     `db:"id"`
		UserId    int64     `db:"user_id"`
		UploadId  int64     `db:"upload_id"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		Deleted   int64     `db:"deleted"`
	}
)

func newResourceRelationModel(conn sqlx.SqlConn) *defaultResourceRelationModel {
	return &defaultResourceRelationModel{
		conn:  conn,
		table: "`resource_relation`",
	}
}

func (m *defaultResourceRelationModel) Insert(ctx context.Context, data *ResourceRelation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, resourceRelationRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.UploadId, data.CreatedAt, data.UpdatedAt, data.Deleted)
	return ret, err
}

func (m *defaultResourceRelationModel) FindOne(ctx context.Context, id int64) (*ResourceRelation, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", resourceRelationRows, m.table)
	var resp ResourceRelation
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultResourceRelationModel) Update(ctx context.Context, data *ResourceRelation) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, resourceRelationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.UploadId, data.CreatedAt, data.UpdatedAt, data.Deleted, data.Id)
	return err
}

func (m *defaultResourceRelationModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultResourceRelationModel) tableName() string {
	return m.table
}
