package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picket-main-service/src/app"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				val, ok := err.(*app.Error)
				if ok {
					ctx.AbortWithStatusJSON(val.StatusCode, val)
					return
				}
				err = app.NewInternalError(err.(error))
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			}
		}()

		ctx.Next()

	}
}
