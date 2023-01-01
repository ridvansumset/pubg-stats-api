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

const (
	ShardSteam = "steam"
	ShardXbox  = "xbox"
)

const (
	EndpointPlayers = "/players"
	EndpointMatches = "/matches/%s"
)

type Client struct {
	Config  Config
	HClient *http.Client
}

type Config struct {
	Host   string
	APIKey string
}

func New() (*Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("Error loading .env file")
	}

	var apiKey = os.Getenv("API_KEY")

	return &Client{
		Config: Config{
			Host:   "https://api.pubg.com/shards",
			APIKey: apiKey,
		},
		HClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

func (c *Client) Request(shard, endpoint string, qs *url.Values) (io.ReadCloser, error) {
	u := c.Config.Host + "/" + shard + endpoint

	req, _ := http.NewRequest(http.MethodGet, u, nil)

	if qs != nil {
		req.URL.RawQuery = qs.Encode()
	}

	req.Header.Set("Authorization", "Bearer "+c.Config.APIKey)
	req.Header.Set("Accept", "application/vnd.api+json")

	res, err := c.HClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Do")
	}

	if res.StatusCode == http.StatusOK {
		return res.Body, nil
	}

	return nil, fmt.Errorf("unexpected status %d for %s", res.StatusCode, req.URL)
}
