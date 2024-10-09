package main

import (
	"NBAPI/internal/server"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting server...")
	server := server.NewServer()
	logrus.Infof("Server listening on %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
