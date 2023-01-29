package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PageLogModel = (*customPageLogModel)(nil)

type (
	// PageLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPageLogModel.
	PageLogModel interface {
		pageLogModel
	}

	customPageLogModel struct {
		*defaultPageLogModel
	}
)

// NewPageLogModel returns a model for the database table.
func NewPageLogModel(conn sqlx.SqlConn) PageLogModel {
	return &customPageLogModel{
		defaultPageLogModel: newPageLogModel(conn),
	}
}
