// Copyright 2013 Alessandro De Donno

// "Betfair API-NG Golang Library" is dual-licensed: for free software projects
// please refer to GPLv3 (see declaration above), for commercial software
// please contact the author.
// If you are a contributor and need any clarification, please contact the
// author.

// For free software projects:

// This file is part of "Betfair API-NG Golang Library".
// "Betfair API-NG Golang Library" is free software: you can redistribute it
// and/or modify it under the terms of the GNU General Public License as
// published by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.
// "Betfair API-NG Golang Library" is distributed in the hope that it will be
// useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
// You should have received a copy of the GNU General Public License
// along with "Betfair API-NG Golang Library".  If not, see
// <http://www.gnu.org/licenses/>.

// CREDITS

// 	Thanks to Iacob and his message for posterity :)
//  https://groups.google.com/d/msg/golang-nuts/dEfqPOSccIc/hoq8jdPTBIcJ

package betfair

import (
	"testing"
	"encoding/json"
	"os"
)

var (
	s 			*Session
	marketId 	string
)

func TestNewSession(t *testing.T) {
	// Get a local configuration for testing
	file, err := os.Open("betfair_test.conf.json")
	if err != nil {
		t.Error(err.Error())
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	config := &Config{}
	dec.Decode(&config)
	// Test
	testS, err := NewSession(config)
	if err != nil {
		t.Error(err.Error())
	}
	// Assign local session to global var for next tests
	s = testS
}

func TestLoginNonInteractive(t *testing.T) {
	if err := s.LoginNonInteractive(); err != nil {
		t.Error(err.Error())
	}
}

func TestKeepAlive(t *testing.T) {
	if err := s.KeepAlive(); err != nil {
		t.Error(err.Error())
	}	
}

func TestListCountries(t *testing.T) {
	filter := new(MarketFilter)
	res, err := s.ListCountries(filter)
	if err != nil {
		t.Error(err.Error())
	}
	if len(res) < 1 {
		t.Error("Result is empty")
	}
}

func TestListCompetitions(t *testing.T) {
	filter := new(MarketFilter)
	res, err := s.ListCompetitions(filter)
	if err != nil {
		t.Error(err.Error())
	}
	if len(res) < 1 {
		t.Error("Result is empty")
	}
}

func TestListEvents(t *testing.T) {
	filter := new(MarketFilter)
	res, err := s.ListEvents(filter)
	if err != nil {
		t.Error(err.Error())
	}
	if len(res) < 1 {
		t.Error("Result is empty")
	}
}

func TestListEventTypes(t *testing.T) {
	filter := new(MarketFilter)
	res, err := s.ListEventTypes(filter)
	if err != nil {
		t.Error(err.Error())
	}
	if len(res) < 1 {
		t.Error("Result is empty")
	}
}

func TestListMarketCatalogue(t *testing.T) {
	filter := new(MarketFilter)
	res, err := s.ListMarketCatalogue(filter, 10)
	if err != nil {
		t.Error(err.Error())
	}
	if len(res) < 1 {
		t.Error("Result is empty")
	}
	// Get a marketId for further tests
	marketId = res[0].MarketId
}

func TestListMarketBook(t *testing.T) {
	marketIds := []string{marketId}
	res, err := s.ListMarketBook(marketIds)
	if err != nil {
		t.Error(err.Error())
	}
	if len(res) < 1 {
		t.Error("Result is empty")
	}
}

func TestListMarketTypes(t *testing.T) {
	filter := new(MarketFilter)
	filter.MarketIds = []string{marketId}
	res, err := s.ListMarketTypes(filter)
	if err != nil {
		t.Error(err.Error())
	}
	if len(res) < 1 {
		t.Error("Result is empty")
	}
}

func TestGetAccountDetails(t *testing.T) {
	_, err := s.GetAccountDetails()
	if err != nil {
		t.Error(err.Error())
	}	
}

func TestGetAccountFunds(t *testing.T) {
	_, err := s.GetAccountFunds()
	if err != nil {
		t.Error(err.Error())
	}	
}

func TestGetDeveloperAppKeys(t *testing.T) {
	_, err := s.GetDeveloperAppKeys()
	if err != nil {
		t.Error(err.Error())
	}	
}

func TestLogout(t *testing.T) {
	if err := s.Logout(); err != nil {
		t.Error(err.Error())
	}	
}
