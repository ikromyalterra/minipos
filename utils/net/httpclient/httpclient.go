package httpclient

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

//HttpDoer HttpDoer
type HttpDoer interface {
	HttpDo() (*http.Response, error)
}

//HttpParam HttpParam
type HttpParam struct {
	// http url
	Url string `json:"url"`
	// http method, available method: get, post, put, patch, delete
	Method string `json:"method"`
	// http header
	Header map[string]string `json:"header"`
	// http body
	Body string `json:"body"`
	// http timeout, in miliseconds
	Timeout int `json:"timeout"`
}

//HttpDo HttpDo
func (httpParam *HttpParam) HttpDo() (*http.Response, error) {
	headers := makeHeader(httpParam.Header)
	timeout := time.Duration(httpParam.Timeout) * time.Second
	switch httpParam.Method {
	case "get":
		return get(httpParam.Url, headers, timeout)
	case "post":
		return post(httpParam.Url, strings.NewReader(httpParam.Body), headers, timeout)
	case "put":
		return put(httpParam.Url, strings.NewReader(httpParam.Body), headers, timeout)
	case "patch":
		return patch(httpParam.Url, strings.NewReader(httpParam.Body), headers, timeout)
	case "delete":
		return delete(httpParam.Url, headers, timeout)
	default:
		return post(httpParam.Url, strings.NewReader(httpParam.Body), headers, timeout)
	}
}

// Get makes a HTTP GET request to provided URL
func get(url string, headers http.Header, timeout time.Duration) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return response, errors.Wrap(err, "GET - request creation failed")
	}

	request.Header = headers

	return do(request, timeout)
}

// Post makes a HTTP POST request to provided URL and requestBody
func post(url string, body io.Reader, headers http.Header, timeout time.Duration) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return response, errors.Wrap(err, "POST - request creation failed")
	}

	request.Header = headers

	return do(request, timeout)
}

// Put makes a HTTP PUT request to provided URL and requestBody
func put(url string, body io.Reader, headers http.Header, timeout time.Duration) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return response, errors.Wrap(err, "PUT - request creation failed")
	}

	request.Header = headers

	return do(request, timeout)
}

// Patch makes a HTTP PATCH request to provided URL and requestBody
func patch(url string, body io.Reader, headers http.Header, timeout time.Duration) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodPatch, url, body)
	if err != nil {
		return response, errors.Wrap(err, "PATCH - request creation failed")
	}

	request.Header = headers

	return do(request, timeout)
}

// Delete makes a HTTP DELETE request with provided URL
func delete(url string, headers http.Header, timeout time.Duration) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return response, errors.Wrap(err, "DELETE - request creation failed")
	}

	request.Header = headers

	return do(request, timeout)
}

// Do makes an HTTP request with the native `http.Do` interface
func do(request *http.Request, timeout time.Duration) (*http.Response, error) {
	var response *http.Response
	var err error

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	var client = &http.Client{Timeout: timeout, Transport: tr}

	response, err = client.Do(request)

	return response, err
}

func makeHeader(headers map[string]string) http.Header {
	result := http.Header{}
	for key, value := range headers {
		result.Add(key, value)
	}
	return result
}

//GetHeaderContentType GetHeaderContentType
func GetHeaderContentType(headers map[string]string) string {
	for k, v := range headers {
		k = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		if v != "" {
			if strings.ToLower(k) == "content-type" {
				return v
			}
		}
	}
	return ""
}
