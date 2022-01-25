package main

import (
	"context"
	"fmt"
	"github.com/go-labs/internal/configs"
	"github.com/go-labs/internal/logging"
	"github.com/go-labs/internal/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config, err := configs.InitConfig("configs")
	if err != nil {
		panic(err)
	}
	var httpSrv interface{}
	if config.Port != 0 {
		httpSrv = startHttpServer(config.Port)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logging.Info().Msg("Shutting down server...")

	if httpSrv != nil {
		stopHttpServer(httpSrv, 5*time.Second)
		httpSrv = nil
	}
}
func startHttpServer(port int) *http.Server {
	//@mark: initialize http web server and start
	router := routers.InitRouter()
	logging.Info().Msgf("Application started, listening and serving HTTP on: %d", port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Fatal().Msgf("listen: %s\n", err)
		}
	}()
	return srv
}

func stopHttpServer(server interface{}, ts time.Duration) {
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	srv := server.(*http.Server)
	ctx, cancel := context.WithTimeout(context.Background(), ts)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logging.Fatal().Msgf("Server forced to shutdown:%s", err.Error())
	}
}
