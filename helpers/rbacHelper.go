package helpers

import (
	"github.com/dipper-lab/asctp-esntls/consts"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func ValidateRole(c *gin.Context, functionality string) (bool, string) {

	if slices.Contains(consts.RBACMap[c.GetString("role")], functionality) {
		return true, ""
	}

	return false, "user unauthorized"
}
