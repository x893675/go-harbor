package goharbor

import (
	"context"
	"encoding/json"
	"github.com/x893675/go-harbor/schema"
)

func (cli *Client) GetSystemInfo(ctx context.Context) (schema.GeneralInfo, error) {
	var systemInfo schema.GeneralInfo

	serverResp, err := cli.get(ctx, "/systeminfo", nil, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return systemInfo, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&systemInfo)
	return systemInfo, err
}
