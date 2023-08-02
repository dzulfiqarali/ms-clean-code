package user

import (
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService service.UserServiceInterface
	//MenuService service.MenuService
}

type Router struct {
	Handler Handler
}

func ProvideHandler(handler Handler) Router {
	return Router{
		Handler: handler,
	}
}

func (h *Handler) SetupRoute(router *gin.Engine) {

	api := router.Group("api")
	v1 := api.Group("/v1.0/service")

	// user
	v1.POST("/user/regist", h.InsertDataUser)

	//menu
}
