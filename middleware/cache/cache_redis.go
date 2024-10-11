package cache

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-gin-template/common/response"
	"go-gin-template/config/redis"
	"log"
	"net/http"
	"time"
)

const cacheKey = "cache::"     // 缓存前缀
const expire = 1 * time.Minute // 缓存时间

var ctx = context.Background()

// Cache 缓存中间件
func Cache() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := cacheKey + c.Request.URL.Path // 生成缓存键

		// 从缓存中读取结果
		cachedResult, err := redis.Rdb.Get(ctx, key).Result()
		if err == nil && cachedResult != "" {
			var result response.R
			if err := json.Unmarshal([]byte(cachedResult), &result); err == nil {
				c.JSON(http.StatusOK, result)
				c.Abort() // 终止后续处理, 直接返回缓存数据
				return
			}
		}

		// 缓存未命中，继续执行后续处理
		c.Next()

		// 如果有响应，将结果存入缓存
		err = redis.Rdb.Set(ctx, key, c.Writer.(response.ResponseWriter).B.String(), expire).Err()
		if err != nil {
			log.Printf("缓存请求结果异常: %v", err)
		}
	}
}
