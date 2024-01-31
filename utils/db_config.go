package utils

import "fmt"

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
}

func (config DbConfig) GetConnConfigs() string {
	connConfigs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Db)
	return connConfigs
}
