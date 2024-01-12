package server

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-maple/nop-kratos-demo/server/internal/conf"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistrar)

func NewRegistrar(serverName string, registry *conf.Registry) registry.Registrar {
	nc := registry.GetNacos()
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(nc.Address, nc.Port),
	}

	cc := constant.ClientConfig{
		AppName:             serverName,
		NamespaceId:         nc.NamespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              nc.LogDir,
		CacheDir:            nc.CacheDir,
		LogLevel:            "error",
		Username:            nc.Username,
		Password:            nc.Password,
	}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  &cc,
		},
	)

	if err != nil {
		panic(err)
	}

	r := nacos.New(client)
	return r
}
