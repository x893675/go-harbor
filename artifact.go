package goharbor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/x893675/goharbor/errdefs"
	"github.com/x893675/goharbor/schema"
	"net/url"
	"strings"
)

func (cli *Client) ListArtifacts(ctx context.Context, options schema.ArtifactsListOptions) ([]schema.Artifact, error) {
	if options.ProjectName == "" || options.RepositoryName == "" {
		return nil, errdefs.InvalidParameter(fmt.Errorf("project name or repo name must valid"))
	}

	options.RepositoryName = strings.ReplaceAll(options.RepositoryName, "/", "%2F")

	var artifacts []schema.Artifact

	query := url.Values{}

	if options.WithTag != nil {
		if *options.WithTag {
			query.Set("with_tag", "true")
		} else {
			query.Set("with_tag", "false")
		}
	}

	if options.WithLabel != nil {
		if *options.WithLabel {
			query.Set("with_laabel", "true")
		} else {
			query.Set("with_label", "false")
		}
	}

	if options.WithSignature != nil {
		if *options.WithTag {
			query.Set("with_signature", "true")
		} else {
			query.Set("with_signature", "false")
		}
	}

	if options.WithImmutableStatus != nil {
		if *options.WithTag {
			query.Set("with_immutable_status", "true")
		} else {
			query.Set("with_immutable_status", "false")
		}
	}

	if options.WithScanOverview != nil {
		if *options.WithScanOverview {
			query.Set("with_scan_overview", "true")
		} else {
			query.Set("with_scan_overview", "false")
		}
	}

	if v := options.QueryString; v != "" {
		query.Set("q", v)
	}

	if v := options.Page; v != "" {
		query.Set("page", v)
	}

	if v := options.PageSize; v != "" {
		query.Set("page_size", v)
	}

	serverResp, err := cli.get(ctx, fmt.Sprintf("/projects/%s/repositories/%s/artifacts", options.ProjectName, options.RepositoryName), query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return artifacts, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&artifacts)
	return artifacts, err
}
