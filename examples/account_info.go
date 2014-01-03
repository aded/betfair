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

	details, err := s.GetAccountDetails()
	checkErr(err)

	funds, err := s.GetAccountFunds()
	checkErr(err)

	fmt.Printf("%s %s (%s)\n", details.FirstName, details.LastName, details.LocaleCode)
	fmt.Printf("\tAvailable: %s %.2f\n", details.CurrencyCode, funds.AvailableToBetBalance)
	fmt.Printf("\tExposure : %s %.2f\n", details.CurrencyCode, funds.Exposure)
}
