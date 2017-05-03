package lib

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	//"net/url"
	"strings"
	"time"

	"github.com/juju/ratelimit"
	"github.com/raben/conoha/lib/models"
)

const (
	Version = "0.1.0"
)

type Client struct {
	client *http.Client

	MaxAttempts int
	bucket      *ratelimit.Bucket

	AuthConfig models.AuthConfig
}

var retryableStatusCodes = map[int]struct{}{
	503: {}, // Rate limit hit
	500: {}, // Internal server error. Try again at a later time.
}

func NewClient() *Client {
	transport := &http.Transport{
		TLSNextProto: make(map[string]func(string, *tls.Conn) http.RoundTripper),
	}
	client := http.DefaultClient
	client.Transport = transport
	rate := 505 * time.Millisecond
	attempts := 1
	return &Client{
		client:      client,
		MaxAttempts: attempts,
		bucket:      ratelimit.NewBucket(rate, 1),
	}
}

func (c *Client) SetAuth(config models.AuthConfig) *Client {
	c.AuthConfig = config
	return c
}

func (c *Client) get(path string, data interface{}) error {
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return err
	}
	return c.do(req, data)
}

func (c *Client) post(path string, value []byte, data interface{}) error {
	req, err := c.newRequest("POST", path, strings.NewReader(string(value)))
	if err != nil {
		return err
	}
	return c.do(req, data)
}

func (c *Client) delete(path string, data interface{}) error {
	req, err := c.newRequest("DELETE", path, nil)
	if err != nil {
		return err
	}
	return c.do(req, data)
}

func (c *Client) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	if len(c.AuthConfig.AuthToken) > 0 {
		req.Header.Set("X-Auth-Token", c.AuthConfig.AuthToken)
	}

	if req.Method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, nil
}

func (c *Client) do(req *http.Request, data interface{}) error {
	c.bucket.Wait(1)

	var apiError error
	for tryCount := 1; tryCount <= c.MaxAttempts; tryCount++ {

		resp, err := c.client.Do(req)
		if err != nil {
			return err
		}

		body, err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()
		if err != nil {
			return err
		}
		if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusMultipleChoices || resp.StatusCode == http.StatusAccepted || resp.StatusCode == http.StatusNoContent {
			if data != nil {
				if string(body) == `[]` {
					data = nil
				} else {
					if err := json.Unmarshal(body, data); err != nil {
						return err
					}
				}
			}
			return nil
		}
		apiError = errors.New(string(body))
		if !isCodeRetryable(resp.StatusCode) {
			break
		}

		delay := backoffDuration(tryCount)
		time.Sleep(delay)
	}

	return apiError
}

func backoffDuration(retryCount int) time.Duration {
	// Upper limit of delay at ~1 minute
	if retryCount > 7 {
		retryCount = 7
	}

	rand.Seed(time.Now().UnixNano())
	delay := (1 << uint(retryCount)) * (rand.Intn(150) + 500)
	return time.Duration(delay) * time.Millisecond
}

func isCodeRetryable(statusCode int) bool {
	if _, ok := retryableStatusCodes[statusCode]; ok {
		return true
	}

	return false
}
