package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
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
	redis                    *redis.Client
	KafkaUrl                 string
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
	d, _ := db.DB()

	d.SetMaxOpenConns(1000)

	client, err := grpc.Dial(structure.AnswerSheetService, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}

	rd := redis.NewClient(&redis.Options{
		Addr: structure.RedisUrl,
		//Password: "mypassword",
		DB: 1,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			log.Info().Interface("connect", cn).Msg("on connect redis")
			return nil
		},
	})
	status := rd.Ping(context.TODO())
	log.Info().Interface("status", status).Send()
	if status.Err() != nil {
		return nil, status.Err()
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
		redis:                    rd,
		KafkaUrl:                 structure.KafkaUrl,
	}

	return &result, nil

}
