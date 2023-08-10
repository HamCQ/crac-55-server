package model

import (
	"context"
	"crac55/common/db"

	"gorm.io/gorm"
)

// Dao 数据操作层
type Dao struct {
	DB *gorm.DB
}

// NewDao new
func NewDao() *Dao {
	return &Dao{
		DB: db.DB,
	}
}

// BeginTran 开启事务
func (d *Dao) BeginTran(c context.Context) *gorm.DB {
	return d.DB.Begin().WithContext(c)
}
