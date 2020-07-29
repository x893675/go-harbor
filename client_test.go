package goharbor

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type FakeRoundTripper struct {
	message  string
	status   int
	header   map[string]string
	requests []*http.Request
}

func (rt *FakeRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	body := strings.NewReader(rt.message)
	rt.requests = append(rt.requests, r)
	res := &http.Response{
		StatusCode: rt.status,
		Body:       ioutil.NopCloser(body),
		Header:     make(http.Header),
	}
	for k, v := range rt.header {
		res.Header.Set(k, v)
	}
	return res, nil
}

func (rt *FakeRoundTripper) Reset() {
	rt.requests = nil
}

func newTestClient(rt http.RoundTripper) *Client {
	endpoint := "http://localhost:4243"
	c := http.Client{
		Transport: rt,
	}
	client, _ := NewClientWithOpts(WithHTTPClient(&c),
		WithHost(endpoint),
		WithScheme("http"),
		WithBasicAuth("admin", "Harbor12345"))
	return client
}
