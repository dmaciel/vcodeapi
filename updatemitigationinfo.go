package vcodeapi

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/brian1917/vcodeHMAC"
)

func updateMitigationInfo(credsFile, buildID, action, comment, flawList string) ([]byte, error) {
	var errorMsg error

	client := http.Client{}

	form := url.Values{}
	form.Add("build_id", buildID)
	form.Add("action", action)
	form.Add("comment", comment)
	form.Add("flaw_id_list", flawList)

	req, err := http.NewRequest("POST", "https://analysiscenter.veracode.com/api/updatemitigationinfo.do",
		strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", vcodeHMAC.GenerateAuthHeader(credsFile, req.Method, req.URL.String()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status != "200 OK" {
		errorMsg = errors.New("updatemitigationinfo.do call error: " + resp.Status)
	}

	return data, errorMsg
}
