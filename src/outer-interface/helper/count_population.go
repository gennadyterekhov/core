package helper

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	cfg := config.GetConfig()
	versionFlag := flag.Bool("v", false, "Show current version")
	flag.Parse()

	if versionFlag != nil && *versionFlag {
		fmt.Printf("Version: %s \n", cfg.Version)
		os.Exit(0)
	}

	gin.SetMode(cfg.AppMode)

	router := initRouter()
	pg.InitDB()
	cfg.HTTPConfig.Host = ""
	err := router.Run(net.JoinHostPort(cfg.HTTPConfig.Host, cfg.HTTPConfig.Port))

	if err != nil {
		log.Fatalf("can not start HTTP server error: %v", err)
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/search", api_handlers.Search)
	router.GET("/health-check", api_handlers.HealthCheck)

	return router
}
