package config

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
)

// NewRegistry 注册nacos
// 这里有一个问题，这个现成的plugin有个问题没有办法设置namespace
func NewRegistry() registry.Registry {
	r := consul.NewRegistry(func(options *registry.Options) {
		// nacos注册中心地址
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	return r
}

//func Register() {
//	serverConfig := []constant.ServerConfig{
//		{
//			IpAddr: "127.0.0.1",
//			Port:   8848,
//		},
//	}
//
//	// 创建clientConfig
//	clientConfig := constant.ClientConfig{
//		NamespaceId:         "", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
//		TimeoutMs:           50000,
//		NotLoadCacheAtStart: true,
//		LogLevel:            "debug",
//	}
//
//	// 创建服务发现客户端的另一种方式 (推荐)
//	namingClient, err := clients.NewNamingClient(
//		vo.NacosClientParam{
//			ClientConfig:  &clientConfig,
//			ServerConfigs: serverConfig,
//		},
//	)
//	if err != nil {
//		log.Fatalf("初始化nacos失败: %s", err.Error())
//	}
//	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
//		Ip:          "127.0.0.1",
//		Port:        uint64(cfg.Port),
//		ServiceName: "emailservice",
//		Weight:      10,
//		Enable:      true,
//		Healthy:     true,
//		Ephemeral:   true,
//		Metadata:    map[string]string{"name": "test"},
//		ClusterName: "DEFAULT",       // 默认值DEFAULT
//		GroupName:   "DEFAULT_GROUP", // 默认值DEFAULT_GROUP
//	})
//	if err != nil {
//		log.Fatalf("注册服务失败: %s", err.Error())
//	}
//
//	log.Println("success: ", success)
//	log.Printf("服务启动成功;PORT:%d\n", cfg.Port)
//}
