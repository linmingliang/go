package main

import (
	"demo/global"
	"demo/pkg/config"
)

func init() {
	err := setupConfig()
	if err != nil {
		//TODO 打印错误日志

	}
}
func main() {

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
