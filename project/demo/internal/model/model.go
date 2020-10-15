package model

import (
	"demo/pkg/config"
	"fmt"
	"github.com/jinzhu/gorm"
)

func NewDB(dataSource config.Configure, runMode string) (*gorm.DB, error) {
	driver := dataSource.Get("Driver").(string)
	username := dataSource.Get("Username").(string)
	password := dataSource.Get("Password").(string)
	schema := dataSource.Get("Schema").(string)
	addr := dataSource.Get("Addr").(string)
	charset := dataSource.Get("Charset").(string)
	parseTime := dataSource.Get("ParseTime").(bool)
	loc := dataSource.Get("Loc").(string)
	maxIdleConns := dataSource.Get("MaxIdleConns").(int)
	maxOpenConns := dataSource.Get("MaxOpenConns").(int)
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
