package svc

import (
	"activity/answer/api/internal/config"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config  config.Config
	KvStore kv.Store
	DbConn  sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		KvStore: kv.NewStore(c.KvConf),
		DbConn:  sqlx.NewSqlConn("mysql", c.Mysql.DataSource),
	}
}
