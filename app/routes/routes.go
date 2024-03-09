package routes

import (
	"os"
	"time"

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
		"X-Requested-With",
		"baggagebaggage",
		"sentry-trace",
	}

	r := gin.New()
	// r.Use(auth.CheckTokenMiddleware())
	r.Use(cors.New(config))

	r.Use(ratelimit.RateLimit("auth", 10, 60*time.Second))

	r.POST("/register", api.RegisterHandler)
	r.POST("/nonce", api.UserNonceHandler)
	r.POST("/login", api.LoginHandler)
	r.POST("/balance", api.GetBalanceHandler)
	r.GET("/pull", api.PullHandler)
	if os.Getenv("DEBUG") == "TRUE" {
		r.POST("/testbalance", api.TestHandler)
	}

	// r.POST("/verify", api.VerifyHandler)

	return r
}
