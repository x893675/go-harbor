package goharbor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/x893675/goharbor/errdefs"
	"github.com/x893675/goharbor/schema"
	"net/http"
	"net/url"
)

func (cli *Client) ListRepos(ctx context.Context, options schema.ListReposOptions) ([]schema.Repo, error) {
	if options.ProjectName == "" {
		return nil, errdefs.FromStatusCode(errors.New("bad request, project name is required"), http.StatusBadRequest)
	}
	var repos []schema.Repo
	query := url.Values{}

	if v := options.QueryString; v != "" {
		query.Set("q", v)
	}

	if v := options.Page; v != "" {
		query.Set("page", v)
	}

	if v := options.PageSize; v != "" {
		query.Set("page_size", v)
	}

	serverResp, err := cli.get(ctx, fmt.Sprintf("/projects/%s/repositories", options.ProjectName), query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return repos, err
	}
	err = json.NewDecoder(serverResp.body).Decode(&repos)
	return repos, err
}
