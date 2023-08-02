package router

import (
	"fmt"

	"bitbucket.org/bridce/ms-clean-code/configs"
	"bitbucket.org/bridce/ms-clean-code/internal/handler/user"
	"github.com/gin-gonic/gin"
)

type Http struct {
	Config *configs.Config
	Gin    *gin.Engine
	Router user.Router
}

func ProvideRoute(Config *configs.Config, router user.Router) *Http {
	return &Http{
		Config: Config,
		Router: router,
	}
}

func (h *Http) Serve() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
//	router.Use(apmgin.Middleware(router))

	h.Router.Handler.SetupRoute(router)
	
	addr := h.Config.Service.Host + ":" + h.Config.Service.Port

	err := router.Run(addr)
	if err != nil {
		fmt.Println(err, "error when start server")
	}

	fmt.Println("Your service is up and running at " + addr)
}
