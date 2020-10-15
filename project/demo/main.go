package main

import (
	"demo/global"
	"demo/internal/model"
	"demo/internal/routers"
	"demo/pkg/config"
	"fmt"
	"net/http"
	"time"
)

func init() {
	err := setupConfig()
	if err != nil {
		//TODO 打印错误日志
	}
	err = setupDataSource()
	if err != nil {
		//TODO 打印错误日志
	}
}
func main() {
	server := http.Server{
		Addr:           fmt.Sprintf("%s:%d", "127.0.0.1", global.ServerConfig.Get("port").(int)),
		Handler:        routers.NewRouter(),
		ReadTimeout:    time.Duration(global.ServerConfig.Get("readtimeout").(int)) * time.Second,
		WriteTimeout:   time.Duration(global.ServerConfig.Get("writetimeout").(int)) * time.Second,
		MaxHeaderBytes: global.ServerConfig.Get("maxheaderbytes").(int),
	}
	err := server.ListenAndServe()
	if err != nil {
		//TODO 打印错误日志
	}

}

func setupConfig() error {
	var err error
	configInfo := config.Info{
		Type: "yaml",
		Name: "config",
		Path: "configs/",
	}
	global.ConfigSetting, err = config.ReadConfig(&configInfo)
	if err != nil {
		return err
	}
	global.ServerConfig = global.ConfigSetting.Get("server").(config.Configure)
	global.DataBaseConfig = global.ConfigSetting.Get("datasource").(config.Configure)
	return nil
}

func setupDataSource() error {
	var err error
	global.DBEngine, err = model.NewDB(global.DataBaseConfig, global.ServerConfig.Get("runmode").(string))
	if err != nil {
		return err
	}
	return nil
}
