package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ms-clean-code/configs"
	"github.com/ms-clean-code/internal/handler"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ServerState is an indicator if this server's state.
type ServerState int

const (
	// ServerStateReady indicates that the server is ready to serve.
	ServerStateReady ServerState = iota + 1
	// ServerStateInGracePeriod indicates that the server is in its grace
	// period and will shut down after it is done cleaning up.
	ServerStateInGracePeriod
	// ServerStateInCleanupPeriod indicates that the server no longer
	// responds to any requests, is cleaning up its internal state, and
	// will shut down shortly.
	ServerStateInCleanupPeriod
)

type Http struct {
	Config  *configs.Config
	Gin     *gin.Engine
	Handler handler.Handler
	State   ServerState
}

func ProvideRoute(Config *configs.Config, handler handler.Handler) *Http {
	return &Http{
		Config:  Config,
		Handler: handler,
	}
}

func (h *Http) Serve() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	//	router.Use(apmgin.Middleware(router))

	h.Handler.SetupRoute(router)

	h.setupGracefulShutdown()
	h.State = ServerStateReady

	addr := h.Config.Service.Host + ":" + h.Config.Service.Port

	err := router.Run(addr)
	if err != nil {
		fmt.Println(err, "error when start server")
	}

	fmt.Println("Your service is up and running at " + addr)
}

func (h *Http) setupGracefulShutdown() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	go h.respondToSigterm(done)
}

func (h *Http) respondToSigterm(done chan os.Signal) {
	<-done
	defer os.Exit(0)

	shutdownConfig := h.Config.Server.Shutdown

	log.Info().Msg("Received SIGTERM.")
	log.Info().Int64("seconds", shutdownConfig.GracePeriodSeconds).Msg("Entering grace period.")
	h.State = ServerStateInGracePeriod
	time.Sleep(time.Duration(shutdownConfig.GracePeriodSeconds) * time.Second)

	log.Info().Int64("seconds", shutdownConfig.CleanupPeriodSeconds).Msg("Entering cleanup period.")
	h.State = ServerStateInCleanupPeriod
	time.Sleep(time.Duration(shutdownConfig.CleanupPeriodSeconds) * time.Second)

	log.Info().Msg("Cleaning up completed. Shutting down now.")
}
