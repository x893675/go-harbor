package main

import (
	"context"
	"crypto/tls"
	"github.com/x893675/goharbor"
	"github.com/x893675/goharbor/schema"
	"log"
	"net/http"
	"strings"
)

const HarborAddress = "https://myharbor.com"

func main() {
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	harborClient, err := goharbor.NewClientWithOpts(goharbor.WithHost(HarborAddress),
		goharbor.WithHTTPClient(c),
		goharbor.WithBasicAuth("admin", "Harbor12345"))
	if err != nil {
		panic(err)
	}
	withScan := true
	Imagename := strings.ReplaceAll("hanamichi/hal9k", "/", "%2F")
	pr, err := harborClient.ListArtifacts(context.TODO(), schema.ArtifactsListOptions{
		ProjectName:      "test",
		RepositoryName:   Imagename,
		WithScanOverview: &withScan,
	})
	if err != nil {
		panic(err)
	}
	for _, item := range pr {
		log.Printf("%+v", item)
	}
}
