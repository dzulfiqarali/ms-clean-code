package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ms-clean-code/internal/handler/action"
	"github.com/ms-clean-code/internal/handler/user"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
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

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	api := router.Group("api")
	v1 := api.Group("/v1")

	v1.Use(prometheusMiddleware())

	//user
	v1.Group("/user")
	{
		v1.POST("/user/register", h.User.InsertDataUser)
		// use duration timeout from config remote; don't hard code duration in param function
		v1.GET("/users", processTimeout(h.User.ResolveListUser, 10*time.Second))
		v1.GET("/user/:nama", processTimeout(h.User.ResolveUserByName, 10*time.Second))
	}

	//action
	v1.Group("/action")
	{
		v1.POST("action/register", h.Action.Create)
	}
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

var totalRequest = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "Number of get request",
	},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_temperature_celsius",
	Help: "Current temperature of the CPU.",
})

func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		totalRequest.Inc()
		responseStatus.WithLabelValues(status).Inc()
	}
}

func init() {
	prometheus.MustRegister(totalRequest)
	prometheus.MustRegister(responseStatus)
	prometheus.MustRegister(cpuTemp)
}
