// Code generated by goctl. DO NOT EDIT.

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
	pageLogFieldNames          = builder.RawFieldNames(&PageLog{})
	pageLogRows                = strings.Join(pageLogFieldNames, ",")
	pageLogRowsExpectAutoSet   = strings.Join(stringx.Remove(pageLogFieldNames, "`id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")
	pageLogRowsWithPlaceHolder = strings.Join(stringx.Remove(pageLogFieldNames, "`id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), "=?,") + "=?"
)

type (
	pageLogModel interface {
		Insert(ctx context.Context, data *PageLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PageLog, error)
		Update(ctx context.Context, data *PageLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPageLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PageLog struct {
		Id      int64     `db:"id"`
		Uid     string    `db:"uid"`
		PageUrl string    `db:"page_url"`
		Ip      string    `db:"ip"`
		LogTime time.Time `db:"log_time"`
	}
)

func newPageLogModel(conn sqlx.SqlConn) *defaultPageLogModel {
	return &defaultPageLogModel{
		conn:  conn,
		table: "`page_log`",
	}
}

func (m *defaultPageLogModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPageLogModel) FindOne(ctx context.Context, id int64) (*PageLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pageLogRows, m.table)
	var resp PageLog
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

func (m *defaultPageLogModel) Insert(ctx context.Context, data *PageLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, pageLogRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uid, data.PageUrl, data.Ip, data.LogTime)
	return ret, err
}

func (m *defaultPageLogModel) Update(ctx context.Context, data *PageLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pageLogRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Uid, data.PageUrl, data.Ip, data.LogTime, data.Id)
	return err
}

func (m *defaultPageLogModel) tableName() string {
	return m.table
}
