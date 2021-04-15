/*
=====================================
;	Author : Geraldie Tanu Saputra
;	Email  : geraldie.saputra@soluix.ai
;	Date   : 15-04-2021
=====================================
*/

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
