package middleware

import (
	"fmt"
	"net/http"

	"github.com/dipper-lab/asctp-esntls/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("No authorization")})
			c.Abort()
			return
		}

		claims, err, _ := helpers.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("facility", claims.Facility)
		c.Set("id", claims.Id)
		c.Set("org_id", claims.OrgId)
		c.Next()
	}
}
