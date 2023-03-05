package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"net/http"
	"os"
	"os/signal"
	"picket-main-service/src/config"
	"picket-main-service/src/routes"
	"syscall"
)

func server(config config.IConfig) *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			if config.CheckIsProduction() {
				gin.SetMode(gin.ReleaseMode)
			}
			r := gin.Default()

			r.Use(otelgin.Middleware("picket-main-service"))

			routes.Routes(ctx, r, config)

			srv := &http.Server{
				Addr:    fmt.Sprintf(":%s", config.GetPort()),
				Handler: r,
			}

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

			go func() {
				log.Info().Str("port", config.GetPort()).Msg("server is running")
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error().Err(err).Msg("server has error")
				}
			}()

			<-quit
			err := srv.Shutdown(ctx)
			if err != nil {
				log.Info().Err(err).Send()
			}

			log.Info().Msg("shutdown server")
		},
	}
}
