package main

import (
	"acemap/data"
	_ "acemap/database"
	"acemap/router"
	"log"
)

func main() {
	data.LoadDataFromFiles()

	addr := "0.0.0.0:8081"
	log.Println("Listening and serving HTTP on " + addr)
	r := router.GetRouter()
	err := r.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
}
