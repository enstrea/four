// +build wireinject

package main

import (
	"four/app/book/internal/biz"
	"four/app/book/internal/conf"
	"four/app/book/internal/data"
	"four/app/book/internal/server"
	"four/app/book/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(biz.ProviderSet, data.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
}
