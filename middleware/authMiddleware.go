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
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No authorization")})
			c.Abort()
			return
		}

		claims, err, _ := helpers.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("faciltity", claims.Facility)
		c.Set("id", claims.Id)
		c.Set("org_id", claims.OrgId)
		c.Next()
	}
}

// func AdminAuthenticate() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		clientToken := c.Request.Header.Get("token")

// 		if clientToken == "" {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No authorization")})
// 			c.Abort()
// 			return
// 		}

// 		claims, err, _ := helpers.ValidateOrgToken(clientToken)

// 		if err != "" {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("id", claims.Id)
// 		c.Set("name", claims.Name)
// 		c.Set("email", claims.Email)
// 		c.Set("org_id", claims.OrgId)
// 		c.Next()
// 	}
// }