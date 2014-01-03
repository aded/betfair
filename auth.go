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
	"encoding/json"
	"strings"
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

	body := strings.NewReader("username=" + s.config.Username + "&password=" + s.config.Password)

	data, err := doRequest(s, "certLogin", "", body)
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

	// Get application keys. It seems we currently have one dev app only.
	apps, err := s.GetDeveloperAppKeys()
	if err != nil {
		return err
	}
	if len(apps) < 1 {
		return errors.New("Cannot get app keys")
	}
	if len(apps[0].AppVersions) != 2 {
		return errors.New("Invalid amount of app versions")		
	}
	if apps[0].AppVersions[0].DelayData {
		s.appKeys[DELAY_DATA] = apps[0].AppVersions[0].ApplicationKey
		s.appKeys[LIVE_DATA] = apps[0].AppVersions[1].ApplicationKey
	} else {
		s.appKeys[DELAY_DATA] = apps[0].AppVersions[1].ApplicationKey
		s.appKeys[LIVE_DATA] = apps[0].AppVersions[0].ApplicationKey		
	}

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
