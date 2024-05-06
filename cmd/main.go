package main

import "app/internal/api/router"

func main() {
	server := router.Setup()
	server.Run(":8080")
}
