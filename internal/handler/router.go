package handler

import (
	"bitbucket.org/bridce/ms-clean-code/internal/handler/action"
	"bitbucket.org/bridce/ms-clean-code/internal/handler/user"
	"github.com/gin-gonic/gin"
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
	v1 := api.Group("/v1.0/service")

	// user
	v1.POST("/user/register", h.User.InsertDataUser)

	//action
	v1.POST("action/register", h.Action.Create)
}
