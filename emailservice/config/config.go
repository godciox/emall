package config

import (
	"fmt"
	"github.com/pkg/errors"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/env"
)

type Config struct {
	Port    int
	Address string
	Tracing TracingConfig
}

type TracingConfig struct {
	Enable bool
	Jaeger JaegerConfig
}

type JaegerConfig struct {
	URL string
}

var cfg = &Config{
	Port:    9555,
	Address: "127.0.0.1",
}

func Address() string {
	return fmt.Sprintf(":%d", cfg.Port)
}

func Tracing() TracingConfig {
	return cfg.Tracing
}

func Load() error {
	configor, err := config.NewConfig(config.WithSource(env.NewSource()))
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

//func InitSetting() {
//	clientConfig := constant.ClientConfig{
//		NamespaceId:         "9dbcec67-96b5-45c9-8e80-737342fac52d",
//		TimeoutMs:           5000,
//		NotLoadCacheAtStart: true,
//		RotateTime:          "1h",
//		LogDir:              "./log",
//		MaxAge:              3,
//		LogLevel:            "debug",
//	}
//	// At least one ServerConfig
//	serverConfigs := []constant.ServerConfig{
//		{
//			IpAddr:      cfg.Address,
//			ContextPath: "/nacos",
//			Port:        8848,
//			Scheme:      "http",
//		},
//	}
//
//}
