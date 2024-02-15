package main

import (
	"user_service/routes"
)

func main() {
	router := routes.SetupRoutes()
	router.Run(":8015") // run our server on port 8080
}
