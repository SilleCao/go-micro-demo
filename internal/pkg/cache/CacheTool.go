package cache

import (
	"encoding/json"
	"time"

	"github.com/SilleCao/golang/go-micro-demo/internal/config"
	"github.com/SilleCao/golang/go-micro-demo/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Set(ctx *gin.Context, key string, value interface{}, expiration string) {
	d, err := time.ParseDuration(expiration)
	if err != nil {
		logger.Err("invalid expiration", ctx, err)
		return
	}
	p, err := json.Marshal(value)
	if err != nil {
		logger.Err("marshal data to json fail", ctx, err)
		return
	}
	err = config.Redis().Set(key, p, d).Err()
	if err != nil {
		logger.Err("add to cache fail", ctx, err)
	}
}

func Get(ctx *gin.Context, key string, dest interface{}) error {
	p, err := config.Redis().Get(key).Result()
	if err != nil {
		logger.Err("add to cache fail", ctx, err)
		return err
	}
	return json.Unmarshal([]byte(p), dest)
}
