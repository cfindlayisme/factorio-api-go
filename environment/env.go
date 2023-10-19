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
		fmt.Println("Failed to run due to invalid env variable for PORT! Is it an integer?")
	}

	return listenPort
}
