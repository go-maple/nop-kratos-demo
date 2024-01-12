// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-maple/nop-kratos-demo/blog/internal/biz"
	"github.com/go-maple/nop-kratos-demo/blog/internal/conf"
	"github.com/go-maple/nop-kratos-demo/blog/internal/data"
	"github.com/go-maple/nop-kratos-demo/blog/internal/server"
	"github.com/go-maple/nop-kratos-demo/blog/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(string2 string, confServer *conf.Server, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confServer, confData, logger)
	if err != nil {
		return nil, nil, err
	}
	articleRepo := data.NewArticleRepo(dataData, logger)
	discovery := server.NewDiscovery(registry)
	viteRepo := data.NewViteRepo(confServer, confData, discovery, logger)
	articleUsecase := biz.NewArticleUsecase(articleRepo, viteRepo, logger)
	blogService := service.NewBlogService(articleUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, logger, blogService)
	grpcServer := server.NewGRPCServer(confServer, logger, blogService)
	registrar := server.NewRegistrar(string2, registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
