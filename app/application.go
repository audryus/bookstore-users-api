package app

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/logger"
)

var (
	router = gin.Default()
)

//StartApplication ...
func StartApplication() {
	logger.Info("Mapping routes")
	mapUrls()
	logger.Info("about to start application")
	router.Run(":8080")
}
