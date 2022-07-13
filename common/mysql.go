package common

import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pwd string `json:"pwd"`
	DataBase string `json:"data_base"`
	Port int64 `json:"port"`
}

// GetMysqlFormConsul 获取mysql的配置
func GetMysqlFormConsul(config config.Config, path ...string) *MysqlConfig{
	mysqlConfig := &MysqlConfig{}
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}