package verihubs

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Send call request to verihubs and build http request data by defined params
func (c *Client) Send(ctx context.Context, req Request, response interface{}) (int, error) {
	var dest = c.host + req.Path
	if len(req.Queries) > 0 {
		newDest, err := c.setQuery(dest, req.Queries)
		if err != nil {
			return http.StatusInternalServerError, errors.Wrap(err, "fail to add query params")
		}

		dest = newDest
	}

	var reqBody io.Reader
	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return http.StatusInternalServerError, errors.Wrap(err, "fail to marshal request body")
		}

		reqBody = bytes.NewReader(jsonBody)
	}

	newReq, err := http.NewRequestWithContext(ctx, req.Method, dest, reqBody)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "fail to create new request with context")
	}

	c.buildHeaders(newReq, req.Headers)

	res, err := c.http.Do(newReq)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "fail to send request")
	}

	if res.StatusCode >= http.StatusBadRequest {
		var errBasic Basic

		defer func() {
			_ = res.Body.Close()
		}()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return http.StatusInternalServerError, errors.Wrap(err, "fail to read body response")
		}

		err = json.Unmarshal(data, &errBasic)
		if err != nil {
			return http.StatusInternalServerError, errors.Wrap(err, "fail to unmarshal response")
		}

		return res.StatusCode, &errBasic
	}

	if response != nil {
		defer func() {
			_ = res.Body.Close()
		}()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return http.StatusInternalServerError, errors.Wrap(err, "fail to read body response")
		}

		err = json.Unmarshal(data, response)
		if err != nil {
			return http.StatusInternalServerError, errors.Wrap(err, "fail to unmarshal response")
		}
	}

	return res.StatusCode, nil
}

// setQuery set query params on host url
func (c *Client) setQuery(host string, queries map[string]string) (string, error) {
	urlHost, err := url.Parse(host)
	if err != nil {
		return "", errors.Wrap(err, "fail to parse url "+host)
	}

	q, _ := url.ParseQuery(urlHost.RawQuery)
	for key, val := range queries {
		q.Add(key, val)
	}

	urlHost.RawQuery = q.Encode()

	return urlHost.String(), nil
}

// buildHeaders set http headers app id, api key and param headers
func (c *Client) buildHeaders(req *http.Request, headers map[string]string) {
	req.Header.Add(HeaderAccept, HeaderJsonValue)
	req.Header.Add(HeaderAppID, c.appID)
	req.Header.Add(HeaderApiKey, c.apiKey)
	req.Header.Add(HeaderContentType, HeaderJsonValue)
	for key, val := range headers {
		req.Header.Set(key, val)
	}
}
