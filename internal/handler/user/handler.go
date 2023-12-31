package user

import (
	"github.com/ms-clean-code/infras/log"
	"github.com/ms-clean-code/internal/domain/user/service"
)

type UserHandler struct {
	UserService service.UserServiceInterface
	l           *log.LogCustom
}

func ProvideUserHandler(userService service.UserServiceInterface, l *log.LogCustom) UserHandler {
	return UserHandler{
		UserService: userService,
		l:           l,
	}
}
