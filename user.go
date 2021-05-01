package goharbor

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/x893675/go-harbor/schema"
)

func (cli *Client) SearchUsers(ctx context.Context, options schema.QueryUserOptions) ([]schema.UserEntity, error) {
	var users []schema.UserEntity

	query := url.Values{}

	if v := options.Name; v != "" {
		query.Set("username", v)
	}

	if v := options.Page; v != "" {
		query.Set("page", v)
	}

	if v := options.PageSize; v != "" {
		query.Set("page_size", v)
	}

	serverResp, err := cli.get(ctx, "/users/search", query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return users, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&users)
	return users, err
}

func (cli *Client) CheckUserExist(ctx context.Context, username string) (bool, error) {
	users, err := cli.SearchUsers(ctx, schema.QueryUserOptions{
		Name:     username,
		PageSize: "20",
	})
	if err != nil || len(users) == 0 {
		return false, err
	}
	return true, nil
}
