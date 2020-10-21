package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName     string         `json:"app_name"`
	AppMode     string         `json:"app_mode"`
	AppHost     string         `json:"app_host"`
	AppPort     string         `json:"app_port"`
	AliSms      AliSms         `json:"ali_sms"`
	Database    DatabaseConfig `json:"database"`
	RedisConfig RedisConfig    `json:"redis_config"`
	PasswordKey string         `json:"password_key"`
}

type AliSms struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
	RegionId     string `json:"region_id"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

var _cfg *Config = nil

/**
读取配置文件
*/
func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}

func GetConfig() *Config {
	return _cfg
}
