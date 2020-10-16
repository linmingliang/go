package model

import (
	"demo/pkg/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB(dataSource config.Configure, runMode string) (*gorm.DB, error) {
	driver := dataSource.MustGet("driver").(string)
	username := dataSource.MustGet("username").(string)
	password := dataSource.MustGet("password").(string)
	schema := dataSource.MustGet("schema").(string)
	addr := dataSource.MustGet("addr").(string)
	charset := dataSource.MustGet("charset").(string)
	parseTime := dataSource.MustGet("parsetime").(bool)
	loc := dataSource.MustGet("loc").(string)
	maxIdleConns := dataSource.MustGet("maxidleconns").(int)
	maxOpenConns := dataSource.MustGet("maxopenconns").(int)
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
