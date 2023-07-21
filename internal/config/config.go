package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

var (
	cfg = &Config{}
)

type (
	Config struct {
		App       AppConfig       `mapstructure:"app" validate:"required"`
		API       APIConfig       `mapstructure:"api" validate:"required"`
		Transport TransportConfig `mapstructure:"transport" validate:"required"`
		Storage   StorageConfig   `mapstructure:"database" validate:"required"`
	}

	AppConfig struct {
	}

	APIConfig struct {
		Scheme         string `mapstructure:"scheme" validate:"required"`
		Host           string `mapstructure:"host" validate:"required"`
		SwaggerEnabled bool   `mapstructure:"swaggerEnabled" validate:"required"`
	}

	TransportConfig struct {
		HTTP HTTPConfig `mapstructure:"http" validate:"required"`
	}

	HTTPConfig struct {
		Host               string         `mapstructure:"host" validate:"required"`
		Port               string         `mapstructure:"port" validate:"required"`
		Timeouts           TimeoutsConfig `mapstructure:"timeouts" validate:"required"`
		MaxHeaderMegabytes int            `mapstructure:"maxHeaderMegabytes" validate:"required"`
	}

	TimeoutsConfig struct {
		Read  time.Duration `mapstructure:"read" validate:"required"`
		Write time.Duration `mapstructure:"write" validate:"required"`
	}

	StorageConfig struct {
		Postgres PostgresConfig `mapstructure:"postgres" validate:"required"`
	}

	PostgresConfig struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     string `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		DBName   string `mapstructure:"dbName" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		SSLMode  string `mapstructure:"sslMode" validate:"required"`
	}
)

func Init(cfgFolder, cfgFile string) error {
	if err := parseConfigFile(cfgFolder, cfgFile); err != nil {
		return fmt.Errorf("parseConfigFile: %w", err)
	}

	if err := unmarshall(cfg); err != nil {
		return fmt.Errorf("unmarshall: %w", err)
	}

	parseEnvFile(cfg)
	return nil
}

func Get() *Config {
	return cfg
}

func parseConfigFile(folder, file string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName(file)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func unmarshall(cfg *Config) error {
	if err := viper.UnmarshalKey("api", &cfg.API); err != nil {
		return fmt.Errorf("viper.UnmarshalKey[api]: %s", err)
	}
	if err := viper.UnmarshalKey("transport", &cfg.Transport); err != nil {
		return fmt.Errorf("viper.UnmarshalKey[transport]: %s", err)
	}
	if err := viper.UnmarshalKey("storage", &cfg.Storage); err != nil {
		return fmt.Errorf("viper.UnmarshalKey[storage]: %s", err)
	}
	return nil
}

func parseEnvFile(cfg *Config) {
	cfg.Storage.Postgres.Password = os.Getenv("DATABASE_PASSWORD")

}

func (p PostgresConfig) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		p.Host, p.User, p.Password, p.DBName, p.Port, p.SSLMode)
}
