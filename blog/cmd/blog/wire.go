//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-maple/nop-kratos-demo/blog/internal/biz"
	"github.com/go-maple/nop-kratos-demo/blog/internal/conf"
	"github.com/go-maple/nop-kratos-demo/blog/internal/data"
	"github.com/go-maple/nop-kratos-demo/blog/internal/server"
	"github.com/go-maple/nop-kratos-demo/blog/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(string, *conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
