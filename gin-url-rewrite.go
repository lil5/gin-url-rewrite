package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func UrlRewrite(r *gin.Engine, from, to string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, from) {
			url := strings.TrimPrefix(c.Request.URL.Path, from)
			url = fmt.Sprintf("%s%s", to, url)
			c.Request.URL.Path = url
			r.HandleContext(c)
		}

		c.Next()
	}
}
