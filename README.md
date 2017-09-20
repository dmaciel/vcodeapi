# Veracode API Package

## Package Documentation
See here: https://godoc.org/github.com/brian1917/vcodeapi

## Description
Go package that provides easy access to the Veracode APIs. Each API has two files: one for making the http request and one for parsing the response.
For example, `detailedreport.go` calls the Veracode API and returns a `[byte]` and `detailedreportparser.go` parses the
XML response and returns usable objects such as flaws.

## Included APIs
1. getapplist.do (/api/5.0/getapplist.do)
2. getbuildlist.do (/api/5.0/getbuildlist.do)
3. getdetailedreport.do (/api/5.0/detailedreport.do)
4. getsandboxlist.do (/api/5.0/getsandboxlist.do)