package config

import "github.com/caarlos0/env"

type ServerConfig struct {
	Port       string `json:"PORT,omitempty" env:"PORT,required"`
	DriverName string `json:"DRIVER,omitempty" env:"DRIVER,required"`
	DataSource string `json:"DATASOURCE,omitempty" env:"DATASOURCE,required"`
}

func LoadEnvConfig() (*ServerConfig, error) {
	var cfg ServerConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
