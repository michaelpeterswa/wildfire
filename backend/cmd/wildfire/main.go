package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/michaelpeterswa/wildfire/backend/internal/cache"
	"github.com/michaelpeterswa/wildfire/backend/internal/db"
	"github.com/michaelpeterswa/wildfire/backend/internal/logging"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	logger, err := logging.InitZapLogger()
	if err != nil {
		log.Fatal("unable to acquire zap logger")
	}

	_, err = cache.InitRedis(ctx, "", 1, "")
	if err != nil {
		logger.Error("unable to acquire redis client", zap.Error(err))
	}

	_, err = db.InitMongo(ctx, "")
	if err != nil {
		logger.Error("unable to acquire mongo client", zap.Error(err))
	}

	r := gin.New()
	r.Use(cors.Default())

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"healthy": "ok",
		})
	})

	err = r.Run()
	if err != nil {
		logger.Fatal("unable to start gin server", zap.Error(err))
	}
}
