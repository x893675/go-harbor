package goharbor

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type BasicAuth struct {
	Username string
	Password string
}

type Client struct {
	client    *http.Client
	host      string
	basicAuth *BasicAuth
	version   string
	scheme    string
	basePath  string
}

func NewClientWithOpts(opts ...Opt) (*Client, error) {

	c := &Client{
		client:    http.DefaultClient,
		host:      DefaultHarborHost,
		basicAuth: nil,
		version:   Version,
		scheme:    "http",
		basePath:  "/api",
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// HTTPClient returns a copy of the HTTP client bound to the server
func (cli *Client) HTTPClient() *http.Client {
	return &*cli.client
}

func (cli *Client) Host() string {
	return cli.host
}

func (cli *Client) Scheme() string {
	return cli.scheme
}

func (cli *Client) Version() string {
	return cli.version
}

// getAPIPath returns the versioned request path to call the api.
// It appends the query parameters to the path if they are not empty.
func (cli *Client) getAPIPath(ctx context.Context, p string, query url.Values) string {
	var apiPath string
	if cli.version != "" {
		v := strings.TrimPrefix(cli.version, "v")
		apiPath = path.Join(cli.basePath, "/v"+v, p)
	} else {
		apiPath = path.Join(cli.basePath, p)
	}
	return (&url.URL{Path: apiPath, RawQuery: query.Encode()}).String()
}
