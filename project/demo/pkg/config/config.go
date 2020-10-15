package config

import (
	"github.com/spf13/viper"
)

type Info struct {
	Type string
	Name string
	Path string
}

type Configure map[string]interface{}

func (c Configure) Get(key string) interface{} {
	value := c[key]
	switch value.(type) {
	case map[string]interface{}:
		return Configure(value.(map[string]interface{}))
	case bool:
		return value.(bool)
	case int:
		return value.(int)
	case string:
		return value.(string)
	default:
		return value
	}
}

func ReadConfig(info *Info) (Configure, error) {
	v := viper.New()
	v.SetConfigType(info.Type)
	v.SetConfigName(info.Name)
	v.AddConfigPath(info.Path)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v.AllSettings(), nil
}
