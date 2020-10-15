package model

import (
	"demo/pkg/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB(dataSource config.Configure, runMode string) (*gorm.DB, error) {
	driver := dataSource.Get("driver").(string)
	username := dataSource.Get("username").(string)
	password := dataSource.Get("password").(string)
	schema := dataSource.Get("schema").(string)
	addr := dataSource.Get("addr").(string)
	charset := dataSource.Get("charset").(string)
	parseTime := dataSource.Get("parsetime").(bool)
	loc := dataSource.Get("loc").(string)
	maxIdleConns := dataSource.Get("maxidleconns").(int)
	maxOpenConns := dataSource.Get("maxopenconns").(int)
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&Loc=%s",
		username, password, addr, schema, charset, parseTime, loc)
	db, err := gorm.Open(driver, conn)
	if err != nil {
		return nil, err
	}

	if runMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxOpenConns)
	return db, nil
}
