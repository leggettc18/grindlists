package app

import "github.com/spf13/viper"

type Config struct {
	// The port for the server to run on.
	Server    *Server
	SecretKey string `mapstructure:"secret_key"`
}

type Server struct {
	Port   int
	DbConn string `mapstructure:"db_conn"`
}

func InitConfig() (*Config, error) {
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
