package rconclient

import (
	"log"
	"net/http"
	"strings"

	"github.com/cfindlayisme/factorio-api-go/environment"
	"github.com/gin-gonic/gin"
	"github.com/gorcon/rcon"
)

func getRconConnection() *rcon.Conn {
	conn, err := rcon.Dial(environment.GetRconConnectUrl(), environment.GetRconPassword())

	if err != nil {
		log.Fatal("Error while connecting to RCON server: ", err)
	}

	return conn
}

func GetVersion(c *gin.Context) {

	conn := getRconConnection()

	response, err := conn.Execute("/version")
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, formatVersion(response))
}

func GetAge(c *gin.Context) {

	conn := getRconConnection()

	response, err := conn.Execute("/time")
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, formatAge(response))
}

func GetSeed(c *gin.Context) {

	conn := getRconConnection()

	response, err := conn.Execute("/seed")
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, formatSeed(response))
}

func formatVersion(version string) string {
	return trimResponse(version)
}

func formatAge(age string) string {
	return trimResponse(age)
}

func formatSeed(seed string) string {
	return trimResponse(seed)
}

func trimResponse(response string) string {
	return strings.TrimLeft(strings.TrimRight(response, "\"\n"), "\"")
}
