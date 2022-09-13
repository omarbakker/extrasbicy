package main

import (
	"extrasbicy/server"
	"fmt"
)

func main() {
	server.RegisterHandlers()
	err := server.Serve()
	if err != nil {
		fmt.Printf("server error: %s\n", err)
	}
}
