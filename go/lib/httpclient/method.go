package httpclient

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func (c *Client) Do(ctx context.Context, method, uri string, headers http.Header, body io.Reader) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, uri, body)
	if err != nil {
		return nil, err
	}
	req.Header = headers

	resp, err := c.HttpCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) GetDecode(ctx context.Context, uri string, headers http.Header, target interface{}) error {
	body, err := c.Do(ctx, http.MethodGet, uri, headers, nil)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}

func (c *Client) PostDecode(ctx context.Context, uri string, headers http.Header, body, target interface{}) error {
	respBody, err := c.Do(ctx, http.MethodPost, uri, headers, convertToReader(body))
	if err != nil {
		return err
	}

	return json.Unmarshal(respBody, target)
}
