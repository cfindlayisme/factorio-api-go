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
		fmt.Println("Failed to run due to invalid env variable for PORT! Setting to default of 8080")
		return 8080
	}

	return listenPort
}
