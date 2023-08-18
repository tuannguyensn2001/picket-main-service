package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"picket-main-service/src/config"
	answersheet_transport "picket-main-service/src/features/answersheet/transport"
	answersheet_usecase "picket-main-service/src/features/answersheet/usecase"
	auth_repository "picket-main-service/src/features/auth/repository"
	auth_transport "picket-main-service/src/features/auth/transport"
	auth_usecase "picket-main-service/src/features/auth/usecase"
	job_repository "picket-main-service/src/features/job/repository"
	job_transport "picket-main-service/src/features/job/transport"
	job_usecase "picket-main-service/src/features/job/usecase"
	test_repository "picket-main-service/src/features/test/repository"
	test_transport "picket-main-service/src/features/test/transport"
	test_usecase "picket-main-service/src/features/test/usecase"
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

	jobRepository := job_repository.New(config.GetDB())
	jobUsecase := job_usecase.New(jobRepository)
	testRepository := test_repository.New(config.GetDB())
	testUsecase := test_usecase.New(testRepository, config.GetRedis())
	testTransport := test_transport.New(ctx, testUsecase)

	answersheetUsecase := answersheet_usecase.New(nil, testUsecase, jobUsecase, config, config.GetRedis())
	answersheetTransport := answersheet_transport.New(ctx, answersheetUsecase)
	job_transport.New(ctx, jobUsecase, config)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	g := r.Group("/api")
	{
		g.POST("/v1/auth/register", authTransport.Register)
		g.POST("/v1/auth/login", authTransport.Login)
		g.POST("/v1/auth/login/google", authTransport.LoginGoogle)
		g.GET("/v1/auth/profile", middlewares.CheckAuth(config), authTransport.GetProfile)

		g.POST("/v1/answersheets/start", middlewares.CheckAuth(config), answersheetTransport.Start)
		g.GET("/v1/answersheets/test/:id/check-doing", middlewares.CheckAuth(config), answersheetTransport.CheckUserDoingTest)
		g.GET("/v1/answersheets/test/:id/content", middlewares.CheckAuth(config), answersheetTransport.GetContent)
		g.GET("/v1/answersheets/test/:id/assignment", middlewares.CheckAuth(config), answersheetTransport.GetCurrentTest)
		g.POST("/v1/answersheets/answer", middlewares.CheckAuth(config), answersheetTransport.UserAnswer)
		g.GET("/v1/tests/preview/:id", testTransport.GetPreview)
		g.POST("/v1/answersheets/submit", middlewares.CheckAuth(config), answersheetTransport.SubmitTest)

		g.GET("/v2/answersheets/test/:id/content", middlewares.CheckAuth(config), answersheetTransport.GetContent)
	}
}
