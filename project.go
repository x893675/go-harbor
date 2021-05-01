package goharbor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/x893675/go-harbor/errdefs"
	"github.com/x893675/go-harbor/schema"
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
	serverResp, err := cli.post(ctx, "/projects", nil, body, JSONContentTypeHeader)
	defer ensureReaderClosed(serverResp)
	return err
}

func (cli *Client) ProjectExist(ctx context.Context, name string) (bool, error) {
	query := url.Values{}
	query.Set("project_name", name)
	serverResp, err := cli.head(ctx, "/projects", query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		if errdefs.IsNotFound(err) {
			return false, nil
		} else {
			return false, wrapResponseError(err, serverResp, "projects", name)
		}
	}
	return true, nil
}

func (cli *Client) ListProjectWebhookJobs(ctx context.Context, options schema.WebHookJobsListOptions) ([]schema.WebHookJob, error) {
	if options.ProjectID == "" || options.PolicyID == "" {
		return nil, errdefs.InvalidParameter(fmt.Errorf("project id and policy id must valid"))
	}

	query := url.Values{}
	query.Set("policy_id", options.PolicyID)

	var jobs []schema.WebHookJob
	serverResp, err := cli.get(ctx, fmt.Sprintf("/projects/%s/webhook/jobs", options.ProjectID), query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return jobs, err
	}
	err = json.NewDecoder(serverResp.body).Decode(&jobs)
	return jobs, err
}

func (cli *Client) GetProject(ctx context.Context, projectNameOrID string) (*schema.Project, error) {
	project := &schema.Project{}
	query := url.Values{}
	serverResp, err := cli.get(ctx, fmt.Sprintf("/projects/%s", projectNameOrID), query, JSONContentTypeHeader)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return project, err
	}
	err = json.NewDecoder(serverResp.body).Decode(project)
	return project, err
}

func (cli *Client) DeleteProject(ctx context.Context, projectNameOrID string) error {
	serverResp, err := cli.delete(ctx, fmt.Sprintf("/projects/%s", projectNameOrID), nil, nil)
	defer ensureReaderClosed(serverResp)
	return err
}

func (cli *Client) ListProjectMembers(ctx context.Context, options *schema.ProjectMemberListOptions) ([]schema.ProjectMemberEntity, error) {
	var projectMembers []schema.ProjectMemberEntity
	if options == nil || options.ProjectID == 0 {
		return projectMembers, errdefs.Unavailable(fmt.Errorf("project id required"))
	}
	query := url.Values{}
	if v := options.EntityName; v != "" {
		query.Set("entityname", v)
	}
	serverResp, err := cli.get(ctx, fmt.Sprintf("/projects/%d/members", options.ProjectID), query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return projectMembers, err
	}
	err = json.NewDecoder(serverResp.body).Decode(&projectMembers)
	return projectMembers, err
}

func (cli *Client) CheckProjectMemberExist(ctx context.Context, projectID int64, username string) (bool, error) {

	result, err := cli.ListProjectMembers(ctx, &schema.ProjectMemberListOptions{
		ProjectID:  projectID,
		EntityName: username,
	})
	if err != nil || len(result) == 0 {
		return false, err
	}
	return true, nil
}

func (cli *Client) AddProjectMember(ctx context.Context, projectID int64, member schema.ProjectMember) error {
	serverResp, err := cli.post(ctx, fmt.Sprintf("/projects/%d/members", projectID), nil, member, JSONContentTypeHeader)
	defer ensureReaderClosed(serverResp)
	return err
}

func (cli *Client) RemoveProjectMember(ctx context.Context, projectID int64, memberID int64) error {
	serverResp, err := cli.delete(ctx, fmt.Sprintf("/projects/%d/members/%d", projectID, memberID), nil, nil)
	defer ensureReaderClosed(serverResp)
	return err
}
