package main

import (
	"aceranking/dao"
	"aceranking/router"
	"flag"
	"log"
)

func main() {
	mode := flag.String("mode", "develop", "Running mode: develop or release.")
	//dataFile := flag.String("data", "data/static/data.bin", "Static data file.")
	flag.Parse()

	dao.InitDatabase(*mode)

	//data.LoadDataFromFiles(*dataFile)

	addr := "0.0.0.0:8081"
	log.Println("Listening and serving HTTP on " + addr)
	r := router.Router()
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
