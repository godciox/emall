package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/env"
)

type Config struct {
	Tracing       TracingConfig
	ServiceName   string `mapstructure:"SERVICE_NAME"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

type TracingConfig struct {
	Enable bool
	Jaeger JaegerConfig
}

type JaegerConfig struct {
	URL string
}

var cfg = &Config{}

func Get() Config {
	return *cfg
}

func Address() string {
	return cfg.ServerAddress
}

func Tracing() TracingConfig {
	return cfg.Tracing
}

func DBDriver() string {
	return cfg.DBDriver
}

func DBSource() string {
	return cfg.DBSource
}

func Load() error {
	configor, err := config.NewConfig(config.WithSource(env.NewSource()))

	viper.AddConfigPath(".")
	viper.SetConfigName("orderservice")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return errors.Wrap(err, "viper.Load")
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return errors.Wrap(err, "configor.New")
	}
	if err := configor.Load(); err != nil {
		return errors.Wrap(err, "configor.Load")
	}
	if err := configor.Scan(cfg); err != nil {
		return errors.Wrap(err, "configor.Scan")
	}
	return nil
}
