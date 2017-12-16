package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// httpRequest data model
type httpRequest struct {
	httpClient *http.Client
}

// NewRequest function for intialize httpRequest object
// Paramter, timeout in time.Duration
func NewRequest(timeout time.Duration) *httpRequest {
	return &httpRequest{
		httpClient: &http.Client{Timeout: time.Second * timeout},
	}
}

// newReq function for initalize http request,
// paramters, http method, uri path, body, and headers
func (c *httpRequest) newReq(method string, fullPath string, body io.Reader, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

// exec function, for execute http request
func (c *httpRequest) exec(req *http.Request, v interface{}) error {

	res, err := c.httpClient.Do(req)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	if v != nil {
		return json.NewDecoder(res.Body).Decode(v)
	}

	return nil
}

// Req public function for call http request
func (c *httpRequest) Req(method, path string, body io.Reader, v interface{}, headers map[string]string) error {
	req, err := c.newReq(method, path, body, headers)

	if err != nil {
		return err
	}

	if err := c.exec(req, v); err != nil {
		return err
	}

	return nil
}
