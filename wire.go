//go:build wireinject
// +build wireinject

package main

import (
	"bitbucket.org/bridce/ms-clean-code/configs"
	"bitbucket.org/bridce/ms-clean-code/infras/database"
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/repository"
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/service"
	"bitbucket.org/bridce/ms-clean-code/internal/handler/user"
	"bitbucket.org/bridce/ms-clean-code/transport/http/router"
	"github.com/google/wire"
)

var Configs = wire.NewSet(
	configs.LoadConfig,
)

var Infra = wire.NewSet(
	database.ConnectDatabase,
)

var UserDomain = wire.NewSet(
	//service
	service.UserInterface,
	//repo
	repository.UserRepoImpl,
)

var Domains = wire.NewSet(
	UserDomain,
)

// set wire to external service
// var External = wire.NewSet(
// call External
// )
var Routing = wire.NewSet(
	wire.Struct(new(user.Handler), "*"),
	user.ProvideUserHandler,
	router.ProvideRoute,
)

func InitializeService() *router.Http {
	wire.Build(
		Configs,
		Infra,
		Domains,
		Routing,
	)

	return &router.Http{}
}
