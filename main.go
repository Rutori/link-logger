package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"link-logger/config"
	"link-logger/db"
)

func main() {
	var pathConfig string
	flag.StringVar(&pathConfig, "config", "config.json", "The config")
	flag.Parse()

	confFile, err := ioutil.ReadFile(pathConfig)
	if err != nil {
		log.Fatalln("config file missing")
	}

	err = config.Load(confFile)
	if err != nil {
		log.Fatalf("config file damaged: %s\n", err.Error())
	}

	err = db.Init()
	if err != nil {
		log.Fatalln("cannot init db")
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Storage.Port), controllers())
	if err != nil {
		log.Fatalf("\nserve error: %s\n", err.Error())
	}
}
