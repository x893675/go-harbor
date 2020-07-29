# go-harbor

[![Go Report Card](https://goreportcard.com/badge/github.com/x893675/go-harbor)](https://goreportcard.com/report/github.com/x893675/go-harbor)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/2a36af593ca14aefae4f630fb4e5b750)](https://www.codacy.com/manual/x893675/go-harbor?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=x893675/go-harbor&amp;utm_campaign=Badge_Grade)
[![codecov](https://codecov.io/gh/x893675/go-harbor/branch/master/graph/badge.svg)](https://codecov.io/gh/x893675/go-harbor)
![Test](https://github.com/x893675/go-harbor/workflows/Test/badge.svg)

This package presents a client for the Harbor Restful API. 

Now Support Harbor api version:

* Harbor v2.0

For more details, check the [Harbor API](https://goharbor.io/docs/1.10/build-customize-contribute/configure-swagger/#viewing-harbor-rest-api).

# Example

```go
package main

import (
	"context"
	"crypto/tls"
	"github.com/x893675/go-harbor"
	"github.com/x893675/go-harbor/schema"
	"log"
	"net/http"
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
	pr, err := harborClient.ListProjects(context.TODO(), schema.ProjectListOptions{})
	if err != nil {
		panic(err)
	}
	for _, item := range pr {
		log.Printf("%+v", item)

	}
}
```

# Developing

All development commands can be seen in the Makefile.

Commited code must pass:

* golangci-lint
* go test

Running make test will run all checks, as well as install any required dependencies.

# TODO LIST

- [ ] Projects API
- [ ] Repos API
- [ ] Artifacts API
- [ ] Unit Tests