package vcodeapi

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/brian1917/vcodeHMAC"
)

func BuildCreate(credsFile, appID string, version string) ([]byte, error) {

	// Create HTTP client and request
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://analysiscenter.veracode.com/api/5.0/createbuild.do?app_id="+appID+"&version="+version, nil)
	if err != nil {
		return nil, err
	}

	// Set authorization header
	authHeader, err := vcodeHMAC.GenerateAuthHeader(credsFile, req.Method, req.URL.String())
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", authHeader)

	// Make HTTP Request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Process response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, errors.New("buildCreate.do call error: " + resp.Status)
	}

	// Return data and nil error
	return data, nil
}
