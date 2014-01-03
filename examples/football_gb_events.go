package main

import (
	"github.com/aded/betfair"
	"flag"
	"log"
	"os"
	"encoding/json"
	"fmt"
)

var confFile = flag.String("conf", "", "A json configuration file")

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	if *confFile == "" {
		log.Fatal("You must specify a json configuration file (-conf=CONF_FILE)")
	}
	file, err := os.Open(*confFile)
	checkErr(err)
	defer file.Close()

	dec := json.NewDecoder(file)
	config := new(betfair.Config)
	dec.Decode(&config)

	s, err := betfair.NewSession(config)
	checkErr(err)

	loginErr := s.LoginNonInteractive()
	checkErr(loginErr)
	defer s.Logout()

	s.Live = true

	filter := new(betfair.MarketFilter)
	filter.EventTypeIds = []string{"1"}
	filter.MarketCountries = []string{"GB"}

	events, err := s.ListEvents(filter)
	checkErr(err)

	for _, event := range events {
		fmt.Printf("%s: [%s] %s (%d markets)\n", event.Event.OpenDate, event.Event.CountryCode, event.Event.Name, event.MarketCount)
	}
}
