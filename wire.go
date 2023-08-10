//go:build wireinject
// +build wireinject

package main

import (
	"bitbucket.org/bridce/ms-clean-code/configs"
	"bitbucket.org/bridce/ms-clean-code/external"
	"bitbucket.org/bridce/ms-clean-code/external/fakeapi"
	"bitbucket.org/bridce/ms-clean-code/infras/database"
	"bitbucket.org/bridce/ms-clean-code/infras/log"
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/repository"
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/service"
	"bitbucket.org/bridce/ms-clean-code/internal/handler"
	"bitbucket.org/bridce/ms-clean-code/internal/handler/action"
	"bitbucket.org/bridce/ms-clean-code/internal/handler/user"
	"bitbucket.org/bridce/ms-clean-code/transport/http"
	"github.com/google/wire"
)

var Configs = wire.NewSet(
	configs.LoadConfig,
)

var Infra = wire.NewSet(
	database.ProvideConn,
	log.ProvideConnElk,
)

var ExternalService = wire.NewSet(
	fakeapi.NewClientRequest,
	wire.Bind(new(fakeapi.FakeApiImpl), new(*fakeapi.ClientImpl)),
	external.ProvideFakeApi,
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

var Handler = wire.NewSet(
	action.ProvideActionHandler,
	user.ProvideUserHandler,
)

var Routing = wire.NewSet(
	handler.ProvideHandler,
	http.ProvideRoute,
)

func InitializeService() *http.Http {
	wire.Build(
		Configs,
		Infra,
		ExternalService,
		Domains,
		Handler,
		Routing,
	)

	return &http.Http{}
}
