package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zzayne/go-blog/config"
	"github.com/zzayne/go-blog/middleware"
	"github.com/zzayne/go-blog/router"
)

func main() {
	fmt.Println("gin.Version: ", gin.Version)

	app := gin.New()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	maxSize := int64(config.ServerConfig.MaxMultipartMemory)
	app.MaxMultipartMemory = maxSize << 20 // 3 MiB

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	app.Use(gin.Logger())
	app.Use(middleware.CORSMiddleware())

	router.Route(app)

	app.Run(":" + fmt.Sprintf("%d", config.ServerConfig.Port))
}
