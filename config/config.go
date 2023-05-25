package config

import (
	"errors"
	"os"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/utils/common"
)

type ApiConfig struct {
	ApiPort string
	ApiHost string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type FileConfig struct {
	LogFilePath string
	Env         string
}

type Config struct {
	DbConfig
	ApiConfig
	FileConfig
}

func (c *Config) ReadConfigFile() error {
	err := common.LoadFileEnv(".env")
	if err != nil {
		return err
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.ApiConfig = ApiConfig{
		ApiHost: os.Getenv("API_HOST"),
		ApiPort: os.Getenv("API_PORT"),
	}

	c.FileConfig = FileConfig{
		LogFilePath: os.Getenv("REQUEST_FILE_PATH"),
	}

	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.Name == "" ||
		c.DbConfig.User == "" || c.DbConfig.Password == "" || c.DbConfig.Driver == "" || c.ApiConfig.ApiHost == "" ||
		c.ApiConfig.ApiPort == "" {
		return errors.New("missing required environment variables")
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfigFile()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
