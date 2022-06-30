package middleware

import (
	"errors"
	"github.com/cristalhq/jwt/v3"
	"github.com/gin-gonic/gin"
	"github.com/yerlanov/xmexercise/internal/config"
	"net/http"
	"strings"
)

var (
	ErrEmptyToken = errors.New("authorization header is not provided")
	InvalidFormat = errors.New("invalid authorization header format")
)

func AuthMiddleware(handler func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		if len(authorizationHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrEmptyToken.Error()})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": InvalidFormat.Error()})
			return
		}

		jwtToken := fields[1]

		key := []byte(config.GetConfig("config.yml").JWTSecret)

		verifier, err := jwt.NewVerifierHS(jwt.HS256, key)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		_, err = jwt.ParseAndVerifyString(jwtToken, verifier)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		handler(c)
	}
}
