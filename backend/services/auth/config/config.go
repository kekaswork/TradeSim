package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/viper"
)

type Config struct {
	Env string `yaml:"env" env-required:"true"`

	Database DatabaseConfig `yaml:"database" env-required:"true"`

	JWT JWTConfig `yaml:"jwt" env-required:"true"`
}

type DatabaseConfig struct {
	Host           string
	Port           int
	User           string
	Password       string
	Name           string
	MaxConnections int `yaml:"max_connections"`
}

type JWTConfig struct {
	Secret            string
	ExpirationMinutes int `yaml:"expiration_minutes"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is not set")
	}

	return MustLoadByPath(path)
}

func fetchConfigPath() (res string) {
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = viper.GetString("CONFIG_PATH")
	}

	return
}

func MustLoadByPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	cfg.Database.Host = viper.GetString("DB_HOST")
	cfg.Database.Port = viper.GetInt("DB_PORT")
	cfg.Database.User = viper.GetString("DB_USER")
	cfg.Database.Password = viper.GetString("DB_PASSWORD")
	cfg.Database.Name = viper.GetString("DB_NAME")

	cfg.JWT.Secret = viper.GetString("JWT_SECRET")

	return &cfg
}
