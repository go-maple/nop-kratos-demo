//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-maple/nop-kratos-demo/server/internal/biz"
	"github.com/go-maple/nop-kratos-demo/server/internal/conf"
	"github.com/go-maple/nop-kratos-demo/server/internal/server"
	"github.com/go-maple/nop-kratos-demo/server/internal/service"
	"github.com/google/wire"
)

func wireApp(string, *conf.Server, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
