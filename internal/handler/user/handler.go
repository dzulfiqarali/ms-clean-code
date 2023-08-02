package user

import (
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService service.UserServiceInterface
}

type Router struct {
	Handler Handler
}

func ProvideUserHandler(userService service.UserService) Handler {
	return Handler{
		UserService: userService,
	}
}

func (h *Handler) SetupRoute(router *gin.Engine) {

	api := router.Group("api")
	v1 := api.Group("/v1.0/service")

	v1.POST("/regist-user", h.InsertDataUser)

}
