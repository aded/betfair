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

package betfair

import (
	"errors"
	"time"
	"net"
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"strings"
	"io/ioutil"
	"net/http"
)

type certLoginResult struct {
	LoginStatus		string	`json:"loginStatus"`
	SessionToken	string 	`json:"sessionToken"`
}

type keepAliveResult struct {
	Token	string `json:"token"`
	Product	string `json:"product"`
	Status	string `json:"status"`
	Error	string `json:"error"`
}

// The non-interactive login method for API-NG requires that you create and
// upload a self signed certificate which will be used, alongside your
// username and password, to authenticate your credentials and generate a
// session token.
func (s *Session) LoginNonInteractive() error {
	url, err := s.getUrl("certLogin", "")
	if err != nil {
		return err
	}

	cert, err := tls.LoadX509KeyPair(s.config.CertFile, s.config.KeyFile)
	if err != nil {
		return err
	}

	ssl := &tls.Config {
		Certificates: []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	ssl.Rand = rand.Reader

	s.httpClient = &http.Client {
		Transport: &http.Transport {
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, time.Duration(time.Second*3))
			},
			TLSClientConfig: ssl,
		},
	}

	reqBody := strings.NewReader("username=" + s.config.Username + "&password=" + s.config.Password)
	req, _ := http.NewRequest("POST", url, reqBody)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-Application", s.config.AppKey)

	res, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var result certLoginResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return err
	}
	if result.LoginStatus != "SUCCESS" {
		return errors.New(result.LoginStatus)
	}

	s.token = result.SessionToken
		
	return nil
}

// You can use Keep Alive to reset the session timeout.
// The session time is currently 20 minutes.  Therefore, you should request Keep Alive
// within this time to prevent session expiry.
func (s *Session) KeepAlive() error {

	var result keepAliveResult

	data, err := doRequest(s, "auth", "keepAlive", strings.NewReader(""))
	
	if err != nil {
		return err
	}	
	if err = json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.Status != "SUCCESS" {
		return errors.New(result.Error)
	}

	return nil
}

// Logout from Betfair.
func (s *Session) Logout() error {

	var result keepAliveResult	

	data, err := doRequest(s, "auth", "logout", strings.NewReader(""))
	
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.Status != "SUCCESS" {
		return errors.New(result.Error)
	}

	return nil
}
