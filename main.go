package main

import (
	"fmt"
	"log"
	"os"

	"github.com/theherk/mai/pkg/macaddress"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no macs given")
	}
	if os.Getenv("MAI_APIKEY") == "" {
		log.Fatal("no api key; set MAI_APIKEY")
	}
	api := macaddress.API{Key: os.Getenv("MAI_APIKEY")}
	for _, q := range os.Args[1:] {
		info, err := api.Get(q)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(info)
	}
}
