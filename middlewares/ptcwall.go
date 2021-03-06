package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// PtcwallAuthRequired rejects request if client ip is not in the list
func PtcwallAuthRequired(postbackPassword string, whitelistIPs string) gin.HandlerFunc {
	ips := make(map[string]struct{})
	for _, v := range strings.Split(whitelistIPs, ",") {
		ips[v] = struct{}{}
	}

	return func(c *gin.Context) {
		password := c.Query("pwd")

		if _, ok := ips[c.ClientIP()]; !ok || password != postbackPassword || postbackPassword == "" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
