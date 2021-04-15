package main

import (
	"log"

	Routes "golang-file-reader/app/routes"
)

func main() {
	log.Println("App starting...")
	Routes.InitMain()
	log.Println("App started")
}
