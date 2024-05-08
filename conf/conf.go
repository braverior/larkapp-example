package conf

import (
	"github.com/spf13/viper"
	"larkapp-example/internal/larkapp"
)

type (
	ServerName string
)

type Config struct {
	LarkAppConfig      *larkapp.LarkAppConfig
	TestCardTemplateId string
}

var GlobalConfig Config

func ConfigInit(configPath string) (err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(&GlobalConfig)
}
