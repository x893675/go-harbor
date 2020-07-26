package goharbor

import (
	"context"
	"encoding/json"
	"github.com/x893675/goharbor/schema"
	"net/url"
)

const (
	listAll = iota
	listPublic
	listPrivate
)

func (cli *Client) ListProjects(ctx context.Context, options schema.ProjectListOptions) ([]schema.Project, error) {
	var projects []schema.Project

	query := url.Values{}

	switch options.Public {
	case listPublic:
		query.Set("public", "1")
	case listPrivate:
		query.Set("public", "0")
	case listAll:
		fallthrough
	default:
		break
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
