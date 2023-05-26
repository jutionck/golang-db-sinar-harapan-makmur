package config

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"strconv"
	"time"

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

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type Config struct {
	DbConfig
	ApiConfig
	FileConfig
	TokenConfig
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

	tokenExpire, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRE"))
	accessTokenLifeTime := time.Duration(tokenExpire) * time.Minute
	if err != nil {
		return errors.New("failed to convert token expire")
	}
	c.TokenConfig = TokenConfig{
		ApplicationName:     os.Getenv("TOKEN_APP_NAME"),
		JwtSignatureKey:     os.Getenv("TOKEN_SECRET"),
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: accessTokenLifeTime,
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
