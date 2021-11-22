package httpdump

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// DumpRequest : rearrange request structure
func DumpRequest(req *http.Request) string {
	// Handling nil pointer
	if req == nil {
		return ""
	}

	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		return fmt.Sprintf("%+v", req)
	}
	return string(requestDump)
}

// DumpResponse : rearrange response structure
func DumpResponse(resp *http.Response) string {
	// Handling nil pointer
	if resp == nil {
		return ""
	}

	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return fmt.Sprintf("%+v", resp)
	}
	return string(responseDump)
}
