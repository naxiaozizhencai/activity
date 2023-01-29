package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AnwserResultModel = (*customAnwserResultModel)(nil)

type (
	// AnwserResultModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAnwserResultModel.
	AnwserResultModel interface {
		anwserResultModel
	}

	customAnwserResultModel struct {
		*defaultAnwserResultModel
	}
)

// NewAnwserResultModel returns a model for the database table.
func NewAnwserResultModel(conn sqlx.SqlConn, c cache.CacheConf) AnwserResultModel {
	return &customAnwserResultModel{
		defaultAnwserResultModel: newAnwserResultModel(conn, c),
	}
}
