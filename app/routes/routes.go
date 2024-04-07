package routes

import (
	"os"
	"strconv"
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

	r.POST("/register", api.RegisterHandler)
	r.POST("/nonce", api.UserNonceHandler)
	r.POST("/login", api.LoginHandler)
	ratelimittimes := os.Getenv("RATE_LIMIT_TIME_PER_MINUTE")

	ratelimittimesint, err := strconv.ParseInt(ratelimittimes, 10, 64)
	if err != nil {
		ratelimittimesint = 60
	}

	r.Use(ratelimit.RateLimit("auth", int(ratelimittimesint), 60*time.Second))
	r.POST("/balance", api.GetBalanceHandler)
	r.GET("/pull", api.PullHandler)
	if os.Getenv("DEBUG") == "TRUE" {
		r.POST("/testbalance", api.TestHandler)
	}

	// r.POST("/verify", api.VerifyHandler)

	return r
}
