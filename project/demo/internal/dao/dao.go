package dao

import (
	"demo/pkg/config"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	db *gorm.DB
}

func NewDao(dataSource config.Configure) (*Dao, error) {
	return &Dao{}, nil
}
