package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Headers() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Protects from MimeType Sniffing
		c.Header("X-Content-Type-Options", "nosniff")
		// Prevents browser from pre-fetching DNS
		c.Header("X-DNS-Prefetch-Control", "off")
		// Denies website content to be served in an iframe
		c.Header("X-Frame-Options", "DENY")
		c.Header("Strict-Transport-Security", "max-age=5184000; includeSubDomains")
		// Prevents Internet Explorer from executing downloads in site's context
		c.Header("X-Download-Options", "noopen")
		// Minimal XSS protection
		c.Header("X-XSS-Protection", "1; mode=block")
	}
}

// CORS adds Cross-Origin Resource Sharing support
func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		MaxAge:           86400,
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "PATCH", "HEAD"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
}
