package environment

import (
	"fmt"
	"os"
	"strconv"
)

func GetListenPort() int {
	listenPortString, stringExists := os.LookupEnv("PORT")

	if !stringExists {
		return 8080
	}

	listenPort, err := strconv.Atoi(listenPortString)

	if err != nil {
		fmt.Println("Failed to set port due to invalid env variable! Setting to default of 8080")
		return 8080
	}

	return listenPort
}

func GetRconConnectUrl() string {
	return os.Getenv("RCONSERVER") + ":" + os.Getenv("RCONPORT")
}

func GetRconPassword() string {
	password, passwordExists := os.LookupEnv("RCONPASSWORD")

	if !passwordExists {
		fmt.Println("RCOCPASSWORD enviorment variable blank!")
	}

	return password
}
