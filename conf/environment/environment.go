package environment

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Environment struct {
	ServerPort uint
}

func ConfigAppEnv() Environment {

	viper.AutomaticEnv()

	env := viper.GetString("ENVIRONMENT")
	logrus.Info("Current environment:", env)

	viper.AddConfigPath("./resources")

	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Info("config file not found")
	}

	viper.SetConfigName("config-" + env)

	err := viper.MergeInConfig()
	if err != nil {
		fmt.Println("environment config file not found")
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	return Environment{ServerPort: viper.GetUint("SERVER_PORT")}
}
