package api

import (
	"expenses/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationPayload = "authorization_payload"

)

// Function verifies whether user has the right to call API functions.
func authMiddlewareCookie(maker token.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("access_token_bearer")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return 
		}

		if len(tokenString) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return 
		}
		
		payload, err := maker.VerifyToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(authorizationPayload, payload)
		ctx.Next()
	}
}

// Function verifies whether access_token, stored in cookie, is valid or not. 
func (server *Server) CheckAuth(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("access_token_bearer")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_, err = server.maker.VerifyToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
