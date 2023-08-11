package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ms-clean-code/internal/handler/action"
	"github.com/ms-clean-code/internal/handler/user"
)

// Handler : populate all domain handler
type Handler struct {
	User   user.UserHandler
	Action action.ActionHandler
}

func ProvideHandler(
	user user.UserHandler,
	action action.ActionHandler,
) Handler {
	return Handler{
		User:   user,
		Action: action,
	}
}

func (h *Handler) SetupRoute(router *gin.Engine) {

	api := router.Group("api")
	v1 := api.Group("/v1/service")

	// user
	v1.POST("/user/register", h.User.InsertDataUser)
	v1.GET("/users", h.User.ResolveListUser)

	//action
	v1.POST("action/register", h.Action.Create)
}
