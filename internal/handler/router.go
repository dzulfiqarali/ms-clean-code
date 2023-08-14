package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ms-clean-code/internal/handler/action"
	"github.com/ms-clean-code/internal/handler/user"
	"net/http"
	"time"
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

	v1.Use()
	// user
	v1.POST("/user/register", h.User.InsertDataUser)

	// use duration timeout from config remote; don't hard code duration in param function
	v1.GET("/users", processTimeout(h.User.ResolveListUser, 10*time.Second))

	//action
	v1.POST("action/register", h.Action.Create)
}

func processTimeout(h gin.HandlerFunc, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), duration)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		processDone := make(chan bool)
		go func() {
			h(c)
			processDone <- true
		}()

		select {
		case <-ctx.Done():
			c.JSON(http.StatusBadRequest, gin.H{"error": "process timeout"})
		case <-processDone:
		}
	}
}
