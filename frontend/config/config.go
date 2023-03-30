package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/env"
)

type Config struct {
	Address          string `mapstructure:"SERVER_ADDRESS"`
	Tracing          TracingConfig
	EmailService     string `mapstructure:"EMAIL_SERVICE_NAME"`
	OrderService     string `mapstructure:"ORDER_SERVICE_NAME"`
	UserService      string `mapstructure:"USER_SERVICE_NAME"`
	ProductService   string `mapstructure:"PRODUCT_SERVICE_NAME"`
	CartService      string `mapstructure:"CART_SERVICE_NAME"`
	InventoryService string `mapstructure:"INVENTORY_SERVICE_NAME"`
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
	return cfg.Address
}

func Tracing() TracingConfig {
	return cfg.Tracing
}

func Load() error {
	configor, err := config.NewConfig(config.WithSource(env.NewSource()))

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
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
