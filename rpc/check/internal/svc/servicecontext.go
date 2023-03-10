package svc

import (
	"github.com/geek377148474/bookstore/rpc/check/internal/config"
	"github.com/geek377148474/bookstore/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.BookModel // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动代码
	}
}
