package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tonystrawberry/playground.go.bank/token"
)

const (
	authorizationHeaderKey  = "authorization"         // authorization header key
	authorizationTypeBearer = "bearer"                // authorization type bearer
	authorizationPayloadKey = "authorization_payload" // authorization payload key (set to the context)
)

// authMiddleware is a middleware that checks the authorization header
// and verifies the token.
// If the token is valid, the payload will be set to the context.
// If the token is invalid, the request will be aborted with a 401 status code.
// If the authorization header is missing, the request will be aborted with a 401 status code.
// If the authorization header is invalid, the request will be aborted with a 401 status code.
// If the authorization type is invalid, the request will be aborted with a 401 status code.
// If the authorization type is missing, the request will be aborted with a 401 status code.
// If the authorization type is not bearer, the request will be aborted with a 401 status code.
func (server *Server) authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader(authorizationHeaderKey)
		if len(authorization) == 0 {
			err := errors.New("authorization header is required")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		fields := strings.Fields(authorization)
		if len(fields) != 2 {
			err := errors.New("authorization header is invalid")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := errors.New("authorization type is invalid")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set(authorizationPayloadKey, payload)

		ctx.Next()
	}
}
