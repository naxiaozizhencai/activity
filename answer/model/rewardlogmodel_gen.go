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
	rewardLogFieldNames          = builder.RawFieldNames(&RewardLog{})
	rewardLogRows                = strings.Join(rewardLogFieldNames, ",")
	rewardLogRowsExpectAutoSet   = strings.Join(stringx.Remove(rewardLogFieldNames, "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), ",")
	rewardLogRowsWithPlaceHolder = strings.Join(stringx.Remove(rewardLogFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), "=?,") + "=?"
)

type (
	rewardLogModel interface {
		Insert(ctx context.Context, data *RewardLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RewardLog, error)
		Update(ctx context.Context, data *RewardLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRewardLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RewardLog struct {
		Id       int64     `db:"id"`
		AnwserId int64     `db:"anwser_id"`
		ItemId   string    `db:"item_id"`
		ItemNum  int64     `db:"item_num"`
		Uid      string    `db:"uid"`
		LogTime  time.Time `db:"log_time"`
	}
)

func newRewardLogModel(conn sqlx.SqlConn) *defaultRewardLogModel {
	return &defaultRewardLogModel{
		conn:  conn,
		table: "`reward_log`",
	}
}

func (m *defaultRewardLogModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRewardLogModel) FindOne(ctx context.Context, id int64) (*RewardLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", rewardLogRows, m.table)
	var resp RewardLog
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

func (m *defaultRewardLogModel) Insert(ctx context.Context, data *RewardLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, rewardLogRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.AnwserId, data.ItemId, data.ItemNum, data.Uid, data.LogTime)
	return ret, err
}

func (m *defaultRewardLogModel) Update(ctx context.Context, data *RewardLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rewardLogRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.AnwserId, data.ItemId, data.ItemNum, data.Uid, data.LogTime, data.Id)
	return err
}

func (m *defaultRewardLogModel) tableName() string {
	return m.table
}
