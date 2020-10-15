package main

import (
	"demo/global"
	"demo/internal/model"
	"demo/internal/routers"
	"demo/pkg/config"
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
		Addr:           "127.0.0.1" + global.ServerConfig.Get("Port").(string),
		Handler:        routers.NewRouter(),
		ReadTimeout:    global.ServerConfig.Get("ReadTimeout").(time.Duration) * time.Second,
		WriteTimeout:   global.ServerConfig.Get("WriteTimeout").(time.Duration) * time.Second,
		MaxHeaderBytes: global.ServerConfig.Get("MaxHeaderBytes").(int),
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
	global.ServerConfig = global.ConfigSetting.Get("Server").(config.Configure)
	global.AppConfig = global.ConfigSetting.Get("App").(config.Configure)
	global.DataBaseConfig = global.DataBaseConfig.Get("DataBase").(config.Configure)
	return nil
}

func setupDataSource() error {
	var err error
	global.DBEngine, err = model.NewDB(global.DataBaseConfig, global.ServerConfig.Get("RunMode").(string))
	if err != nil {
		return err
	}
	return nil
}
