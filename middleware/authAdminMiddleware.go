package middleware

import (
	"fmt"
	"net/http"

	"github.com/dipper-lab/asctp-esntls/helpers"
	"github.com/gin-gonic/gin"
)

func AdminAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No authorization")})
			c.Abort()
			return
		}

		claims, err, _ := helpers.ValidateOrgToken(clientToken)

		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("id", claims.Id)
		c.Set("name", claims.Name)
		c.Set("email", claims.Email)
		c.Set("org_id", claims.OrgId)
		c.Next()
	}
}
