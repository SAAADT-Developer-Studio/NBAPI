package main

import (
	"NBAPI/internal/server"
	"fmt"
)

func main() {

	fmt.Println("Starting server...")
	server := server.NewServer()
	fmt.Printf("Server listening on %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
