package user

import "bitbucket.org/bridce/ms-clean-code/internal/domain/user/service"

type UserHandler struct {
	UserService service.UserServiceInterface
}

func ProvideUserHandler(userService service.UserServiceInterface) UserHandler {
	return UserHandler{
		UserService: userService,
	}
}
