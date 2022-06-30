package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yerlanov/xmexercise/internal/company/ipapi"
	"net/http"
)

var notAvailable = errors.New("service not available in your country")

var availableCountries = map[string]struct{}{
	"CY": {},
}

func LocationAccessMiddleware(handler func(c *gin.Context), service ipapi.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		info, err := service.GetIPInformation(c, c.RemoteIP())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "service unavailable"})
			return
		}

		if _, ok := availableCountries[info.CountryCode]; !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": notAvailable.Error()})
			return
		}
		handler(c)
	}
}
