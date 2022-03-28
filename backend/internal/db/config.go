package db

import "fmt"

type Config struct {
	Host     string `yaml:"host" config:"required"`
	Port     string `yaml:"port" config:"required"`
	Username string `yaml:"username" config:"required"`
	Password string `yaml:"password" config:"required"`
	Name     string `yaml:"database" config:"required"`
}

func (config *Config) uri() string {
	ret := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
	)
	return ret
}
