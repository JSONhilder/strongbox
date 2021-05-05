package utils

import "github.com/spf13/viper"

type Config struct {
	Version  string `mapstructure:"VERSION"`
	FilePath string `mapstructure:"FILE_PATH"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return config, nil
}
