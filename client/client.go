package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Client struct {
	Config  Config
	HClient *http.Client
}

type Config struct {
	Host       string
	APIKey     string
	MaxRetries int
	SleepMS    int64
}

func New() (*Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("Error loading .env file")
	}

	var apiKey = os.Getenv("API_KEY")

	return &Client{
		Config: Config{
			Host:       "https://api.pubg.com/shards",
			APIKey:     apiKey,
			MaxRetries: 4,
			SleepMS:    500,
		},
		HClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

func (c *Client) Request(uri string, qs url.Values) (io.ReadCloser, error) {
	req, _ := http.NewRequest(http.MethodGet, c.Config.Host+uri, nil)
	req.URL.RawQuery = qs.Encode()

	req.Header.Set("Authorization", "Bearer "+c.Config.APIKey)
	req.Header.Set("Accept", "application/vnd.api+json")

	return c.request(req)
}

func (c *Client) request(req *http.Request) (io.ReadCloser, error) {
	attempts := c.Config.MaxRetries
	for attempts > 0 {
		res, err := c.HClient.Do(req)
		if err != nil {
			return nil, errors.Wrap(err, "Do")
		}

		if res.StatusCode == http.StatusOK {
			return res.Body, nil
		}

		if res.StatusCode == http.StatusNotFound {
			time.Sleep(time.Duration(c.Config.SleepMS) * time.Millisecond)
			attempts--
			res.Body.Close()
			continue
		}

		res.Body.Close()
		return nil, fmt.Errorf("unexpected status %d for %s", res.StatusCode, req.URL)
	}

	return nil, fmt.Errorf("could not find data after %d attempts", c.Config.MaxRetries)
}
