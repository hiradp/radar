package radar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/hiradp/radar/pkg/util"
)

// Scan takes a hostname of type string and returns a
// pointer to a Host.
func Scan(host string) (*Host, error) {
	baseUrl := util.GetEnv("SERVER_URL", "https://api.ssllabs.com/api/v2/analyze")

	// Fetch the data from the cache to ensure we get status READY immediately.
	url := fmt.Sprintf("%s?&all=on&fromCache=on&host=%s", baseUrl, host)
	resp, err := fetch(url)
	if err != nil {
		return nil, err
	}

	today := time.Now()
	yesterday := today.Add(-24 * time.Hour)
	testTime := time.Unix(resp.TestTime/1000, 0)

	// WARN the user if the data is old and start a new analysis.
	if testTime.Before(yesterday) {
		log.Println("WARN analysis was conducted more than a day ago, kicking off a new analysis.")
		url = fmt.Sprintf("%s?&all=on&startNew=on&publish=publish&host=%s", baseUrl, host)
		_, err := fetch(url)
		if err != nil {
			log.Println("WARN failed to kick off a new analysis.")
		}
	}

	// Clean the data and return it
	return process(*resp), nil
}

// fetch takes a string url and returns the analyzeResponse or
// an error, if one occurs.
func fetch(url string) (*analyzeResponse, error) {
	log.Println("Making a GET request to", url)

	response, err := http.Get(url)
	if err != nil {
		log.Println("ERROR failed to make GET request.", err)
		return nil, err
	}

	if response == nil {
		log.Fatalln("Unable to fetch response from", url)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("ERROR failed to read response body.")
		return nil, err
	}

	var resp analyzeResponse
	if err = json.Unmarshal(responseBody, &resp); err != nil {
		log.Println("ERROR failed to parse response body.")
		return nil, err
	}

	return &resp, nil
}

// process takes in a dto of type analyzeResponse and cleans
// the data so that only necessary information is carried through
// the program.
// If we want to add more information, we need to modify this function
func process(dto analyzeResponse) *Host {
	host := Host{
		Name:      dto.Host,
		Endpoints: make([]Endpoint, len(dto.Endpoints)),
	}

	today := time.Now()

	for i, e := range dto.Endpoints {
		certDto := e.Details.Cert
		notBefore := time.Unix(certDto.NotBefore/1000, 0)
		notAfter := time.Unix(certDto.NotAfter/1000, 0)
		host.Endpoints[i] = Endpoint{
			IPAddress:  e.IPAddress,
			ServerName: e.ServerName,
			Grade:      e.Grade,
			Cert: Cert{
				Issuer:    certDto.IssuerLabel,
				ExpiresAt: notAfter.UTC().String(),
				// 0 - not checked
				// 1 - certificate revoked
				// 2 - certificate not revoked
				// 3 - revocation check error
				// 4 - no revocation information
				// 5 - internal error
				IsValid: certDto.RevocationStatus == 2 && today.Before(notAfter) && today.After(notBefore),
			},
		}
	}

	return &host
}
