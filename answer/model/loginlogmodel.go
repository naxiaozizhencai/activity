package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LoginLogModel = (*customLoginLogModel)(nil)

type (
	// LoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLoginLogModel.
	LoginLogModel interface {
		loginLogModel
		FindAllGroupByLanguage(context.Context, string, string) (interface{}, error)
	}

	customLoginLogModel struct {
		*defaultLoginLogModel
	}
)

// NewLoginLogModel returns a model for the database table.
func NewLoginLogModel(conn sqlx.SqlConn) LoginLogModel {
	return &customLoginLogModel{
		defaultLoginLogModel: newLoginLogModel(conn),
	}
}

func (c *customLoginLogModel) FindAllGroupByLanguage(ctx context.Context, start string, end string) (interface{}, error) {
	return nil, nil
}
