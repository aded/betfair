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
	"errors"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
)

var ukEndpoints = map[string]string{
	"certLogin" : "https://identitysso-api.betfair.com/api/certlogin",
	"auth"		: "https://identitysso.betfair.com/api/",
	"betting"	: "https://api.betfair.com/exchange/betting/rest/v1.0/",
	"account"	: "https://api.betfair.com/exchange/account/rest/v1.0/",
}

var auEndpoints = map[string]string{
	"certLogin" : "https://identitysso-api.betfair.com/api/certlogin",
	"auth"		: "https://identitysso.betfair.com/api/",
	"betting"	: "https://api-au.betfair.com/exchange/betting/rest/v1.0/",
	"account"	: "https://api-au.betfair.com/exchange/account/rest/v1.0/",
}

var endpointMap = map[string]map[string]string{
	"UK": ukEndpoints,
	"AU": auEndpoints,
}

type Config struct {
	Exchange	string
	CertFile 	string
	KeyFile 	string
	Username 	string
	Password 	string
	AppKey 		string
	Locale		string
}

type Session struct {
	config		*Config
	token 		string
	httpClient	*http.Client
	appKey 		string
}

// Create a new session. Please note that you have to login to retrieve a
// valid session token.
func NewSession(c *Config) (*Session, error) {

	s := new(Session)

	if _, exists := endpointMap[c.Exchange]; exists == false {
		return s, errors.New("Invalid Config.Exchange: must be UK or AU.")
	}
	if _, err := os.Stat(c.CertFile); os.IsNotExist(err) {
		return s, errors.New("Config.CertFile does not exist.")
	}
	if _, err := os.Stat(c.KeyFile); os.IsNotExist(err) {
		return s, errors.New("Config.KeyFile does not exist.")
	}
	if c.Username == "" {
		return s, errors.New("Config.Username is empty.")		
	}
	if c.Password == "" {
		return s, errors.New("Config.Password is empty.")		
	}
	if c.AppKey == "" {
		return s, errors.New("Config.AppKey is empty.")		
	}
	if c.Locale == "" {
		c.Locale = "en"
	}

	s.config = c

	return s, nil
}

// Builds URLs for API methods.
func (s *Session) getUrl(key, method string) (string, error) {
	if _, exists := endpointMap[s.config.Exchange][key]; exists == false {
		return "", errors.New("Invalid endpoint key: " + key)
	}
	return endpointMap[s.config.Exchange][key] + method, nil
}

// Makes requests to Betfair API via http client.
func doRequest(s *Session, key, method string, body *strings.Reader) ([]byte, error) {

	endpoint, err := s.getUrl(key, method)
	if err != nil {
		return nil, err
	}	

	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Application", s.config.AppKey)
	req.Header.Add("X-Authentication", s.token)

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
