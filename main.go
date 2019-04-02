package main

import (
	"articles-service/routes"
)

func main() {
	router := routes.SetupRouter()

	router.Run()
}
