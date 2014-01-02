// Copyright 2013 Alessandro De Donno <mail@aded.it>

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
	"encoding/json"
	"strings"
)

// Response for Account details.
type AccountDetailsResponse struct {
	CurrencyCode	string
	FirstName		string
	LastName		string
	LocaleCode		string
	Region			string
	Timezone		string
	DiscountRate	float32
	PointsBalance	int
}

// Response for retrieving available to bet.
type AccountFundsResponse struct {
	AvailableToBetBalance	float32
	Exposure				float32
	RetainedCommission		float32
	ExposureLimit			float32
}

// Describes developer/vendor specific application.
type DeveloperApp struct {
	AppName		string
	AppId		uint64
	AppVersions	[]DeveloperAppVersion
}

// Describes a version of an external application.
type DeveloperAppVersion struct {
	Owner					string
	VersionId				uint64
	Version					string
	ApplicationKey			string
	DelayData				bool
	SubscriptionRequired	bool
	OwnerManaged			bool
	Active					bool
}

// Get Account details.
func (s *Session) GetAccountDetails() (AccountDetailsResponse, error) {
	var response AccountDetailsResponse
	err := doAccountRequest(s, "getAccountDetails", &response)
	return response, err
}

// Get available to bet amount.
func (s *Session) GetAccountFunds() (AccountFundsResponse, error) {
	var response AccountFundsResponse
	err := doAccountRequest(s, "getAccountFunds", &response)
	return response, err
}

// Get all application keys owned by the given developer/vendor.
func (s *Session) GetDeveloperAppKeys() ([]DeveloperApp, error) {
	var response []DeveloperApp
	err := doAccountRequest(s, "getDeveloperAppKeys", &response)
	return response, err
}

func doAccountRequest(s *Session, method string, v interface{}) error {
	data, err := doRequest(s, "account", method, strings.NewReader(""))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
