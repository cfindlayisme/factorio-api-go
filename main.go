// REST API for factorio servers
//
// Author: Chuck Findlay <chuck@findlayis.me>
// License: LGPL v3.0
package main

import (
	"github.com/cfindlayisme/factorio-api-go/rconclient"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/version", rconclient.GetVersion)
	router.GET("/age", rconclient.GetAge)

	router.Run("localhost:8080")
}
