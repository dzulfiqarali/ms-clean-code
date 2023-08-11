//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ms-clean-code/configs"
	"github.com/ms-clean-code/external"
	"github.com/ms-clean-code/external/fakeapi"
	"github.com/ms-clean-code/infras/database"
	"github.com/ms-clean-code/infras/log"
	"github.com/ms-clean-code/internal/domain/user/repository"
	"github.com/ms-clean-code/internal/domain/user/service"
	"github.com/ms-clean-code/internal/handler"
	"github.com/ms-clean-code/internal/handler/action"
	"github.com/ms-clean-code/internal/handler/user"
	"github.com/ms-clean-code/transport/http"
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
