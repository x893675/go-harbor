package schema

import "time"

// Project holds the details of a project.
type Project struct {
	ProjectID    int64             `json:"project_id"`
	OwnerID      int               `json:"owner_id"`
	Name         string            `json:"name"`
	CreationTime time.Time         `json:"creation_time"`
	UpdateTime   time.Time         `json:"update_time"`
	Deleted      bool              `json:"deleted"`
	OwnerName    string            `json:"owner_name"`
	Role         int               `json:"current_user_role_id"`
	RoleList     []int             `json:"current_user_role_ids"`
	RepoCount    int64             `json:"repo_count"`
	ChartCount   uint64            `json:"chart_count"`
	Metadata     map[string]string `json:"metadata"`
	CVEAllowlist CVEAllowlist      `json:"cve_allowlist"`
	RegistryID   int64             `json:"registry_id"`
}

type ProjectListOptions struct {
	Name     string
	Public   *bool
	Owner    string
	Page     string
	PageSize string
}

type CreateProjectOptions struct {
	Name         string           `json:"project_name"`
	Metadata     *ProjectMetadata `json:"metadata,omitempty"`
	CVEAllowlist *CVEAllowlist    `json:"cve_allowlist,omitempty"`
	StorageLimit *int64           `json:"storage_limit,omitempty"`
	CountLimit   *int64           `json:"count_limit,omitempty"`
}

type ProjectMetadata struct {
	// Whether content trust is enabled or not. If it is enabled, user can't pull unsigned images from this project. The valid values are "true", "false".
	EnableContentTrust string `json:"enable_content_trust,omitempty"`
	// Whether scan images automatically when pushing. The valid values are "true", "false".
	AutoScan string `json:"auto_scan,omitempty"`
	// If the vulnerability is high than severity defined here, the images can't be pulled. The valid values are "none", "low", "medium", "high", "critical".
	Severity string `json:"severity,omitempty"`
	// Whether this project reuse the system level CVE whitelist as the whitelist of its own. The valid values are "true", "false". If it is set to "true" the actual whitelist associate with this project, if any, will be ignored.
	ReuseSysCveWhitelist string `json:"reuse_sys_cve_whitelist,omitempty"`
	// The public status of the project. The valid values are "true", "false".
	Public string `json:"public,omitempty"`
	// Whether prevent the vulnerable images from running. The valid values are "true", "false".
	PreventVul string `json:"prevent_vul,omitempty"`
}

type WebHookJobsListOptions struct {
	// Relevant project ID, required
	ProjectID string
	// the policy ID
	PolicyID string
}

type WebHookJob struct {
	Status     string `json:"status,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
	EventType  string `json:"event_type,omitempty"`
	CreateTime string `json:"creation_time,omitempty"`
	JobDetail  string `json:"job_detail,omitempty"`
	ID         int64  `json:"id,omitempty"`
	NotifyType string `json:"notify_type,omitempty"`
	PolicyID   int64  `json:"policy_id,omitempty"`
}

type ProjectMemberListOptions struct {
	ProjectID  int64
	EntityName string
}

type ProjectMemberEntity struct {
	ID         int64  `json:"id,omitempty"`
	ProjectID  int64  `json:"project_id,omitempty"`
	EntityName string `json:"entity_name,omitempty"`
	RoleName   string `json:"role_name,omitempty"`
	RoleID     int64  `json:"role_id,omitempty"`
	EntityID   int64  `json:"entity_id,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
}
