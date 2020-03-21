package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping function to ping test the server
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
