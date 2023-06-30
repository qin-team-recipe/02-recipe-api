package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qin-team-recipe/02-recipe-api/pkg/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func errorResponse(err error) gin.H {
	return gin.H{
		"message": err.Error(),
	}
}

func JwtAuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHandler := ctx.Request.Header.Get("authorization")

		if len(authHandler) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authHandler)
		// if len(fields) < 2 {
		// 	err := errors.New("invalid authorization header format")
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		// 	return
		// }

		accessToken := fields[0]
		payload, err := tokenMaker.VerifyJwtToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		/*
			ここにRedisに保存したPayloadの各情報を比較して整合していく予定
		*/

		// 有効期限の確認
		if payload.ExpiredAt < time.Now().Unix() {
			err := errors.New("authorization header is expired at")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Next()
	}
}
