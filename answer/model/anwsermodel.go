package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AnwserModel = (*customAnwserModel)(nil)

type (
	// AnwserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAnwserModel.
	AnwserModel interface {
		anwserModel
		FindAll(ctx context.Context) ([]Anwser, error)
	}

	customAnwserModel struct {
		*defaultAnwserModel
	}
)

// NewAnwserModel returns a model for the database table.
func NewAnwserModel(conn sqlx.SqlConn, c cache.CacheConf) AnwserModel {
	return &customAnwserModel{
		defaultAnwserModel: newAnwserModel(conn, c),
	}
}

func (c customAnwserModel) FindAll(ctx context.Context) ([]Anwser, error) {
	data := make([]Anwser, 0)
	query := fmt.Sprintf("select %s from %s", anwserRows, c.defaultAnwserModel.table)
	err := c.defaultAnwserModel.QueryRowsNoCacheCtx(ctx, &data, query)
	if err != nil {
		return nil, err
	}
	return data, nil
}
