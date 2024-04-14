package sensors

import (
	"github.com/broemp/cannaBro/config"
	"github.com/spf13/viper"
)

var SensorConfig *viper.Viper

func Init() {
	SensorConfig = viper.New()

	SensorConfig.SetConfigType("json")
	SensorConfig.SetConfigFile("sensors")
	SensorConfig.AddConfigPath(config.Env.ConfigPath)
}
