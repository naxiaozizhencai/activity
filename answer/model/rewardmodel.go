package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ RewardModel = (*customRewardModel)(nil)

type (
	// RewardModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRewardModel.
	RewardModel interface {
		rewardModel
		FindUserAward(ctx context.Context, gameUid string, answerId int64, ItemId string) (*Reward, error)
		FindUserFragmentRewards(ctx context.Context, gameUid string, ItemIds []string) ([]Reward, error)
		UpdateStatus(ctx context.Context, id int64, status int) error
	}

	customRewardModel struct {
		*defaultRewardModel
	}
)

// NewRewardModel returns a model for the database table.
func NewRewardModel(conn sqlx.SqlConn, c cache.CacheConf) RewardModel {
	return &customRewardModel{
		defaultRewardModel: newRewardModel(conn, c),
	}
}

func (c *customRewardModel) FindUserAward(ctx context.Context, gameUid string, answerId int64, ItemId string) (*Reward, error) {
	var data Reward
	query := fmt.Sprintf("select %s from %s where game_uid = ? and answer_id = ? and item_id = ?",
		rewardRows, c.defaultRewardModel.table)
	err := c.defaultRewardModel.QueryRowNoCacheCtx(ctx, &data, query, gameUid, answerId, ItemId)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (c *customRewardModel) FindUserFragmentRewards(ctx context.Context, gameUid string, ItemIds []string) ([]Reward, error) {
	var data []Reward
	query := fmt.Sprintf("select %s from %s where game_uid = ? and item_id IN ('%s') ",
		rewardRows, c.defaultRewardModel.table, strings.Join(ItemIds, "','"))
	err := c.defaultRewardModel.QueryRowsNoCacheCtx(ctx, &data, query, gameUid)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (m *defaultRewardModel) UpdateStatus(ctx context.Context, id int64, status int) error {
	rewardIdKey := fmt.Sprintf("%s%v", cacheRewardIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set status = ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, status, id)
	}, rewardIdKey)
	return err
}
