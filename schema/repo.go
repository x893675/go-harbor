package schema

import "time"

type Repo struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	ProjectID     int64     `json:"project_id"`
	Description   string    `json:"description,omitempty"`
	PullCount     int64     `json:"pull_count"`
	CreationTime  time.Time `json:"creation_time"`
	UpdateTime    time.Time `json:"update_time"`
	ArtifactCount int64     `json:"artifact_count"`
}

type ListReposOptions struct {
	ProjectName string
	//Query string to query resources. Supported query patterns are "exact match(k=v)", "fuzzy match(k=~v)", "range(k=[min~max])",
	//"list with union releationship(k={v1 v2 v3})" and "list with intersetion relationship(k=(v1 v2 v3))".
	//The value of range and list can be string(enclosed by " or '),
	//integer or time(in format "2020-04-09 02:36:00").
	//All of these query patterns should be put in the query string "q=xxx" and splitted by ",". e.g. q=k1=v1,k2=~v2,k3=[min~max]
	QueryString string
	Page        string
	PageSize    string
}
