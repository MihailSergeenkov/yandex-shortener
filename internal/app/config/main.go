package config

import (
	"flag"
	"os"
)

type Settings struct {
	RunAddr string
	BaseUrl string
}

var Params Settings

func ParseFlags() {
	flag.StringVar(&Params.RunAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&Params.BaseUrl, "b", "http://localhost:8080", "address and port to urls")

	flag.Parse()

	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		Params.RunAddr = envRunAddr
	}

	if envBaseUrl := os.Getenv("BASE_URL"); envBaseUrl != "" {
		Params.BaseUrl = envBaseUrl
	}
}