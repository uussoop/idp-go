package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/uussoop/idp-go/middleware/ratelimit"
	"github.com/uussoop/idp-go/routes/api"
)

func InitRouter() *gin.Engine {

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
	}

	r := gin.New()
	// r.Use(auth.CheckTokenMiddleware())
	r.Use(ratelimit.RateLimit("auth", 10, 60))

	r.POST("/register", api.RegisterHandler)
	r.POST("/nonce/", api.UserNonceHandler)
	r.POST("/login", api.LoginHandler)
	// r.POST("/verify", api.VerifyHandler)

	return r
}
