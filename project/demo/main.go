package main

import (
	"demo/global"
	"demo/internal/routers"
	"demo/pkg/config"
	"net/http"
)

func init() {
	err := setupConfig()
	if err != nil {
		//TODO 打印错误日志

	}
}
func main() {
	server := http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        routers.NewRouter(),
		ReadTimeout:    60,
		WriteTimeout:   60,
		MaxHeaderBytes: 1 << 20,
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
	global.Config, err = config.ReadConfig(&configInfo)
	if err != nil {
		return err
	}
	return nil
}
