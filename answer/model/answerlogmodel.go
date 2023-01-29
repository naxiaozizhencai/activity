package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AnswerLogModel = (*customAnswerLogModel)(nil)

type (
	// AnswerLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAnswerLogModel.
	AnswerLogModel interface {
		answerLogModel
	}

	customAnswerLogModel struct {
		*defaultAnswerLogModel
	}
)

// NewAnswerLogModel returns a model for the database table.
func NewAnswerLogModel(conn sqlx.SqlConn) AnswerLogModel {
	return &customAnswerLogModel{
		defaultAnswerLogModel: newAnswerLogModel(conn),
	}
}
