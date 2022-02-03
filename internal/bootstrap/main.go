package bootstrap

import (
	"os"
	"strconv"

	"github.com/Pauloo27/fazendocafe/internal/server"
	"github.com/joho/godotenv"
)

var (
	port int
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	rawPort := os.Getenv("CAFE_PORT")

	var err error
	port, err = strconv.Atoi(rawPort)
	if err != nil {
		panic(err)
	}
}

func Start() {
	server.Start(port)
}
