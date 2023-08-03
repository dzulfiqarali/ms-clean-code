package http

import (
	"bitbucket.org/bridce/ms-clean-code/configs"
	"bitbucket.org/bridce/ms-clean-code/internal/handler"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Http struct {
	Config  *configs.Config
	Gin     *gin.Engine
	Handler handler.Handler
}

func ProvideRoute(Config *configs.Config, handler handler.Handler) *Http {
	return &Http{
		Config:  Config,
		Handler: handler,
	}
}

func (h *Http) Serve() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	//	router.Use(apmgin.Middleware(router))

	h.Handler.SetupRoute(router)

	addr := h.Config.Service.Host + ":" + h.Config.Service.Port

	err := router.Run(addr)
	if err != nil {
		fmt.Println(err, "error when start server")
	}

	fmt.Println("Your service is up and running at " + addr)
}
