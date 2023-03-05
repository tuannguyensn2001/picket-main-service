package config

import (
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type IConfig interface {
	GetDB() *gorm.DB
	CheckIsProduction() bool
	GetPort() string
	GetSecretKey() string
	GetGoogleClientId() string
	GetGoogleClientSecret() string
	GetClientUrl() string
	GetConnectToAnswersheetService() *grpc.ClientConn
}

func (c config) GetDB() *gorm.DB {
	return c.db
}

func (c config) CheckIsProduction() bool {
	return c.appEnv == "production"
}

func (c config) GetPort() string {
	return c.appHttpPort
}

func (c config) GetSecretKey() string {
	return c.appSecretKey
}

func (c config) GetGoogleClientId() string {
	return c.oauth2GoogleClientId
}

func (c config) GetGoogleClientSecret() string {
	return c.oauth2GoogleClientSecret
}

func (c config) GetClientUrl() string {
	return c.clientUrl
}

func (c config) GetConnectToAnswersheetService() *grpc.ClientConn {
	return c.connToAnswersheetService
}
