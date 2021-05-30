package app

import "github.com/spf13/viper"

type Config struct {
	// The port for the server to run on.
	Server *Server
}

type Server struct {
	Port int
}

func InitConfig() (*Config, error) {
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}