/* We define a config variable where we can store the server configuration which will
* contain fields like host and port to set up the server.
 */

package core

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigType struct {
	Host string
	Port string
}

// read the .env file from the root directory and set the config variables accordingly
func readConfig() ConfigType {
	// read the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		config := ConfigType {
			Host: "localhost",
			Port: "8080",
		}
		return config
	}

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	if port == "" {
		log.Println("No port provided. Using default port 8080")
		port = "8080"
	}

	if host == "" {
		log.Println("No host provided. Using default host localhost")
		host = "localhost"
	}

	// set the config variables
	config := ConfigType{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}

	return config
}

// create the config variable
var Config = readConfig()
