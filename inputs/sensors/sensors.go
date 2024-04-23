package sensors

import (
	"github.com/broemp/growBro/config"
	"github.com/spf13/viper"
	"github.com/stianeikeland/go-rpio/v4"
	"go.uber.org/zap"
)

var SensorConfig *viper.Viper

func Init() {
	SensorConfig = viper.New()

	SensorConfig.SetConfigType("json")
	SensorConfig.SetConfigFile("sensors")
	SensorConfig.AddConfigPath(config.Env.ConfigPath)

	err := rpio.Open()
	if err != nil {
		zap.L().Error("cannot open rpio", zap.Error(err))
	}
}
