package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"picket-main-service/src/config"
	auth_repository "picket-main-service/src/features/auth/repository"
	auth_transport "picket-main-service/src/features/auth/transport"
	auth_usecase "picket-main-service/src/features/auth/usecase"
	"picket-main-service/src/middlewares"
	google_service "picket-main-service/src/services/google"
)

func Routes(ctx context.Context, r *gin.Engine, config config.IConfig) {

	r.Use(middlewares.Cors)
	r.Use(middlewares.Recover())

	googleService := google_service.New(config)
	authRepository := auth_repository.New(config.GetDB())
	authUsecase := auth_usecase.New(authRepository, config, googleService)
	authTransport := auth_transport.New(ctx, authUsecase)

	g := r.Group("/api")
	{
		g.POST("/v1/auth/register", authTransport.Register)
		g.POST("/v1/auth/login", authTransport.Login)
		g.POST("/v1/auth/login/google", authTransport.LoginGoogle)
		g.GET("/v1/auth/profile", middlewares.CheckAuth(config), authTransport.GetProfile)
	}
}
