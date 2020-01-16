package vcodeapi

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/brian1917/vcodeHMAC"
)

func buildCreate(credsFile, appID string, version string) ([]byte, error) {

	// Create HTTP form
	form := url.Values{}
	form.Add("app_id", appID)
	form.Add("version", version)

	// Create HTTP client and request
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://analysiscenter.veracode.com/api/5.0/createBuild.do",
		strings.NewReader(form.Encode()))
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
		return nil, errors.New("getbuildlist.do call error: " + resp.Status)
	}

	// Return data and nil error
	return data, nil
}
