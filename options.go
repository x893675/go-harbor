package goharbor

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	Version = "v2.0"
	DefaultHarborHost = "http://localhost"
)

type Opt func(*Client) error

func WithHost(host string) Opt {
	return func(c *Client) error {
		if strings.HasSuffix(host, "/") {
			host = host[:len(host)-1]
		}
		//c.host = host

		hostURL, err := ParseHostURL(host)
		if err != nil {
			return err
		}
		c.host = hostURL.Host
		c.scheme = hostURL.Scheme

		//c.basePath = host
		return nil
	}
}

// ParseHostURL parses a url string, validates the string is a host url, and
// returns the parsed URL
func ParseHostURL(host string) (*url.URL, error) {
	protoAddrParts := strings.SplitN(host, "://", 2)
	if len(protoAddrParts) == 1 {
		return nil, fmt.Errorf("unable to parse harbor host `%s`", host)
	}

	proto, addr := protoAddrParts[0], protoAddrParts[1]
	return &url.URL{
		Scheme: proto,
		Host:   addr,
	}, nil
}

// WithHTTPClient overrides the client http client with the specified one
func WithHTTPClient(client *http.Client) Opt {
	return func(c *Client) error {
		if client != nil {
			c.client = client
		}
		return nil
	}
}

// WithScheme overrides the client scheme with the specified one
func WithScheme(scheme string) Opt {
	return func(c *Client) error {
		c.scheme = scheme
		return nil
	}
}

func WithBasicAuth(username, password string) Opt {
	return func(c *Client) error {
		c.basicAuth = &BasicAuth{Username: username, Password: password}
		return nil
	}
}
