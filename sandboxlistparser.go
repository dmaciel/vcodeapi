package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"errors"
	"log"
)

type Sandbox struct {
	SandboxID   string `xml:"sandbox_id,attr"`
	SandboxName string `xml:"sandbox_name,attr"`
	Owner       string `xml:"owner,attr"`
}

func ParseSandboxList(username, password, app_id string) ([]Sandbox, error) {
	var sandboxes []Sandbox
	var errMsg error = nil

	sandboxListAPI, err := sandboxList(username, password, app_id)
	if err != nil {
		log.Fatal(err)
	}
	decoder := xml.NewDecoder(bytes.NewReader(sandboxListAPI))
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()

		if t == nil {
			break
		}
		// Inspect the type of the token just read
		switch se := t.(type) {
		case xml.StartElement:
			// Read StartElement and check for flaw
			if se.Name.Local == "sandbox" {
				var sandbox Sandbox
				decoder.DecodeElement(&sandbox, &se)
				sandboxes = append(sandboxes, sandbox)
			}
			if se.Name.Local == "error" {
				errMsg = errors.New("api for GetSandboxList returned with an error element")
			}
		}
	}
	return sandboxes, errMsg
}
