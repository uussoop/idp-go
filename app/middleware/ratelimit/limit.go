package ratelimit

import (
	"fmt"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/uussoop/idp-go/pkg/cache"
)

func RateLimit(_key string, maxRequest int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		agentSign := "agent-sign"
		fingerPrint := fmt.Sprintf("%s-%s", c.ClientIP(), agentSign)

		key := fmt.Sprintf("rl:%s:%s", fingerPrint, _key)

		getCache := cache.GetCache()
		data, _, ok := getCache.GetWithExpiration(key)
		reqNum := 0

		if ok {
			reqNum = data.(int)
		}

		if ok {
			reqNum = data.(int)
			if data.(int) == maxRequest {
				getCache.Set(key, reqNum+1, duration)
				c.AbortWithStatusJSON(429, gin.H{"message": "rate limit exceeded"})
				return
			}
			if data.(int) > maxRequest {
				c.AbortWithStatusJSON(429, gin.H{"message": "rate limit exceeded"})
				return
			}
		}
		getCache.Set(key, reqNum+1, duration)
		c.Next()
	}
}
