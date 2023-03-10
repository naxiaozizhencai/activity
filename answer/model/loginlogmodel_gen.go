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
	loginLogFieldNames          = builder.RawFieldNames(&LoginLog{})
	loginLogRows                = strings.Join(loginLogFieldNames, ",")
	loginLogRowsExpectAutoSet   = strings.Join(stringx.Remove(loginLogFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	loginLogRowsWithPlaceHolder = strings.Join(stringx.Remove(loginLogFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"
)

type (
	loginLogModel interface {
		Insert(ctx context.Context, data *LoginLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*LoginLog, error)
		Update(ctx context.Context, data *LoginLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLoginLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	LoginLog struct {
		Id        int64     `db:"id"`
		Language  string    `db:"language"`
		Uid       string    `db:"uid"`
		LoginTime time.Time `db:"login_time"`
		LogTime   time.Time `db:"log_time"`
	}
)

func newLoginLogModel(conn sqlx.SqlConn) *defaultLoginLogModel {
	return &defaultLoginLogModel{
		conn:  conn,
		table: "`login_log`",
	}
}

func (m *defaultLoginLogModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultLoginLogModel) FindOne(ctx context.Context, id int64) (*LoginLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", loginLogRows, m.table)
	var resp LoginLog
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

func (m *defaultLoginLogModel) Insert(ctx context.Context, data *LoginLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, loginLogRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Language, data.Uid, data.LoginTime, data.LogTime)
	return ret, err
}

func (m *defaultLoginLogModel) Update(ctx context.Context, data *LoginLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, loginLogRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Language, data.Uid, data.LoginTime, data.LogTime, data.Id)
	return err
}

func (m *defaultLoginLogModel) tableName() string {
	return m.table
}
