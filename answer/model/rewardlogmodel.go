package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RewardLogModel = (*customRewardLogModel)(nil)

type (
	// RewardLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRewardLogModel.
	RewardLogModel interface {
		rewardLogModel
	}

	customRewardLogModel struct {
		*defaultRewardLogModel
	}
)

// NewRewardLogModel returns a model for the database table.
func NewRewardLogModel(conn sqlx.SqlConn) RewardLogModel {
	return &customRewardLogModel{
		defaultRewardLogModel: newRewardLogModel(conn),
	}
}
