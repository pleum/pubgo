package pubgo

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func createRequest(c Client, route string, query url.Values, auth bool) (*bytes.Buffer, error) {
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		return nil, err
	}

	if auth {
		req.Header.Set("Accept", "application/vnd.api+json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}

	if c.gzip {
		req.Header.Set("Accept-Encoding", "gzip")
	}

	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case http.StatusUnauthorized:
			return nil, errors.New("API key invalid or missing")
		case http.StatusNotFound:
			return nil, errors.New("The specified resource was not found")
		case http.StatusUnsupportedMediaType:
			return nil, errors.New("Content type incorrect or not specified")
		case http.StatusTooManyRequests:
			return nil, errors.New("Too many requests")
		default:
			return nil, errors.New(res.Status)
		}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(body), nil
}
