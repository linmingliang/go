package main

import (
	"demo/global"
	"demo/internal/model"
	"demo/internal/routers"
	"demo/pkg/config"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"net/http"
	"os"
	"path"
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
	setupLogger()
}
func main() {
	server := http.Server{
		Addr:           fmt.Sprintf("%s:%d", "127.0.0.1", global.ServerConfig.MustGet("port").(int)),
		Handler:        routers.NewRouter(),
		ReadTimeout:    time.Duration(global.ServerConfig.MustGet("readtimeout").(int)) * time.Second,
		WriteTimeout:   time.Duration(global.ServerConfig.MustGet("writetimeout").(int)) * time.Second,
		MaxHeaderBytes: global.ServerConfig.MustGet("maxheaderbytes").(int),
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
	global.ServerConfig = global.ConfigSetting.MustGet("server").(config.Configure)
	global.DataBaseConfig = global.ConfigSetting.MustGet("datasource").(config.Configure)
	return nil
}

func setupDataSource() error {
	var err error
	global.DBEngine, err = model.NewDB(global.DataBaseConfig, global.ServerConfig.MustGet("runmode").(string))
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() {
	configure, _ := global.ServerConfig.MustGet("logger").(config.Configure)
	r, b := global.ServerConfig.Get("runmode")
	runMode := "debug"
	if b {
		runMode = r.(string)
	}
	var i io.Writer = os.Stdout
	if runMode == "release" {
		i = &lumberjack.Logger{
			Filename:  path.Join(configure.MustGet("path").(string), configure.MustGet("file").(string)),
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: false,
		}
	}
	global.Logger = log.New(i, "", 0)
}
