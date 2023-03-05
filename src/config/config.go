package config

import (
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type config struct {
	appHttpPort              string
	appGrpcPort              string
	appHttpAddress           string
	appEnv                   string
	appSecretKey             string
	oauth2GoogleClientId     string
	oauth2GoogleClientSecret string
	clientUrl                string
	db                       *gorm.DB
	connToAnswersheetService *grpc.ClientConn
}

func GetConfig() (*config, error) {
	structure := bootstrap()

	db, err := gorm.Open(postgres.Open(structure.DatabaseUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	db.Use(otelgorm.NewPlugin())

	client, err := grpc.Dial("localhost:30000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}

	result := config{
		appHttpPort:              structure.AppHttpPort,
		appGrpcPort:              structure.AppGrpcPort,
		appHttpAddress:           structure.AppHttpAddress,
		appEnv:                   structure.AppEnv,
		appSecretKey:             structure.AppSecretKey,
		oauth2GoogleClientId:     structure.Oauth2GoogleClientId,
		oauth2GoogleClientSecret: structure.Oauth2GoogleClientSecret,
		clientUrl:                structure.ClientUrl,
		db:                       db,
		connToAnswersheetService: client,
	}

	return &result, nil

}
