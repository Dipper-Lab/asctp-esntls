package helpers

import (
	"net/http"

	"github.com/dipper-lab/asctp-esntls/consts"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func ValidateRole(c *gin.Context, functionality string) {

	if !slices.Contains(consts.RBACMap[c.GetString("role")], functionality) {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user unauthorized"})
		return
	}

}
