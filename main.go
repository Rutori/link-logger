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
	closer := make(chan struct{})

	go func() {
		er := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Storage.Host, config.Storage.Port), controllers())
		if er != nil {
			log.Fatalf("\nserve error: %s\n", er.Error())
		}
	}()

	if config.Storage.CertPath != "" && config.Storage.KeyPath != "" {
		go func() {
			er := http.ListenAndServeTLS(fmt.Sprintf("%s:%d", config.Storage.Host, config.Storage.TLSPort), config.Storage.CertPath, config.Storage.KeyPath, controllers())
			if er != nil {
				log.Fatalf("\nserve error TLS: %s\n", er.Error())
			}
		}()
	}

	<-closer
}
