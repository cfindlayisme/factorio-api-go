// REST API for factorio servers
//
// Author: Chuck Findlay <chuck@findlayis.me>
// License: LGPL v3.0
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorcon/rcon"
)

func getRconConnection() *rcon.Conn {
	conn, err := rcon.Dial(os.Getenv("RCONSERVER")+":"+os.Getenv("RCONPORT"), os.Getenv("RCONPASSWORD"))

	if err != nil {
		log.Fatal("Error while connecting to RCON server: ", err)
	}

	return conn
}

func getVersion(c *gin.Context) {

	conn := getRconConnection()

	response, err := conn.Execute("/version")
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, response)
}

func main() {

	router := gin.Default()
	router.GET("/version", getVersion)

	router.Run("localhost:8080")
}
