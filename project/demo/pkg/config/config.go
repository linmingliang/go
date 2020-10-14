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
	case Configure:
		v := value.(Configure)
		return v
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
