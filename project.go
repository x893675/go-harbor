package goharbor

import (
	"context"
	"encoding/json"
	"github.com/x893675/go-harbor/schema"
	"net/url"
)

func (cli *Client) ListProjects(ctx context.Context, options schema.ProjectListOptions) ([]schema.Project, error) {
	var projects []schema.Project

	query := url.Values{}

	if options.Public != nil {
		if *options.Public {
			query.Set("public", "1")
		} else {
			query.Set("public", "0")
		}
	}

	if v := options.Name; v != "" {
		query.Set("name", v)
	}

	if v := options.Owner; v != "" {
		query.Set("owner", v)
	}

	if v := options.Page; v != "" {
		query.Set("page", v)
	}

	if v := options.PageSize; v != "" {
		query.Set("page_size", v)
	}

	serverResp, err := cli.get(ctx, "/projects", query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return projects, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&projects)
	return projects, err
}

func (cli *Client) CreateProject(ctx context.Context, body schema.CreateProjectOptions) error {
	serverResp, err := cli.post(ctx, "/projects", nil, body, nil)
	defer ensureReaderClosed(serverResp)
	return err
}

func (cli *Client) ProjectExist(ctx context.Context, name string) (bool, error) {
	query := url.Values{}
	query.Set("project_name", name)
	serverResp, err := cli.head(ctx, "/projects", query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return false, wrapResponseError(err, serverResp, "projects", name)
	}
	return true, nil
}
