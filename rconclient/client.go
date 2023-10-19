package rconclient

import (
	"log"
	"net/http"
	"os"
	"strings"

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

func formatVersion(version string) string {
	return trimResponse(version)
}

func formatAge(age string) string {
	return trimResponse(age)
}

func trimResponse(response string) string {
	return strings.TrimLeft(strings.TrimRight(response, "\"\n"), "\"")
}