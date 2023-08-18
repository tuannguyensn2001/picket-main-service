package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"picket-main-service/src/app"
	"picket-main-service/src/config"
	"picket-main-service/src/utils"
	"strconv"
)

func CheckAuth(config config.IConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//var list []int
		//for i := 0; i < 1000; i++ {
		//	list = append(list, i+1)
		//}
		//userId := lo.Sample(list)
		//ctx.Set("user_id", userId)
		//ctx.Next()
		//
		//return
		token, err := utils.GetBearerToken(ctx.GetHeader("authorization"))
		if err != nil {
			log.Error().Err(err).Send()
			panic(app.NewForbiddenError(err, "forbidden"))
		}
		t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetSecretKey()), nil
		})
		if err != nil {
			panic(err)
		}

		if claims, ok := t.Claims.(*jwt.RegisteredClaims); ok && t.Valid {
			userId, err := strconv.Atoi(claims.ID)
			if err != nil {
				panic(app.NewForbiddenError(err))
			}
			ctx.Set("user_id", userId)
			ctx.Next()
			return
		}

		panic(app.NewForbiddenError(errors.New("forbidden")))

	}
}
