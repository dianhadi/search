package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Server   Server   `yaml:"server"`
	Consumer Consumer `yaml:"consumer"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Consumer struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func GetConfig(file string) (AppConfig, error) {
	var config AppConfig

	data, err := os.ReadFile(file)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
