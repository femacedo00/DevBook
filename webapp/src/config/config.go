package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL represents the URL tha connects to the API
	APIURL = ""
	// Port is the port where the web application is running
	Port = 0
	// HashKey authenticates the cookie
	HashKey []byte
	// BlockKey encrypt the cookie
	BlockKey []byte
)

// Load initializes the environment variables
func Load() {
	var error error
	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Port, error = strconv.Atoi(os.Getenv("APP_PORT"))
	if error != nil {
		log.Fatal()
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}
