package ratelimit

import (
	"fmt"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/uussoop/idp-go/pkg/cache"
)

func calculateMaxRequestsPerHour(maxRequest int, duration time.Duration) string {
	// Calculate the number of requests per hour
	requestsPerHour := float64(maxRequest) / duration.Hours()

	// Convert the result to a string
	return fmt.Sprintf("%d ", int(requestsPerHour))
}

func RateLimit(_key string, maxRequest int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		agentSign := "agent-sign"
		fingerPrint := fmt.Sprintf("%s-%s", c.ClientIP(), agentSign)

		key := fmt.Sprintf("rl:%s:%s", fingerPrint, _key)

		getCache := cache.GetCache()
		data, exptime, ok := getCache.GetWithExpiration(key)
		reqNum := 0

		if ok {
			reqNum = data.(int)
			c.Writer.Header().Set("X-RateLimit-Limit", calculateMaxRequestsPerHour(maxRequest, duration))
			c.Writer.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d ", int(maxRequest-reqNum)))
			c.Writer.Header().Set("X-RateLimit-Reset", strconv.FormatInt(exptime.Unix(), 10))
		} else {
			c.Writer.Header().Set("X-RateLimit-Limit", calculateMaxRequestsPerHour(maxRequest, duration))
			c.Writer.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d ", int(maxRequest)))
			c.Writer.Header().Set("X-RateLimit-Reset", strconv.FormatInt(time.Now().Add(duration).Unix(), 10))
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
