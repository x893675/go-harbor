package schema

import "time"

// CVEAllowlist defines the data model for a CVE allowlist
type CVEAllowlist struct {
	ID           int64              `json:"id"`
	ProjectID    int64              `json:"project_id"`
	ExpiresAt    *int64             `json:"expires_at,omitempty"`
	Items        []CVEAllowlistItem `json:"items"`
	ItemsText    string             `json:"-"`
	CreationTime time.Time          `json:"creation_time"`
	UpdateTime   time.Time          `json:"update_time"`
}

// CVEAllowlistItem defines one item in the CVE allowlist
type CVEAllowlistItem struct {
	CVEID string `json:"cve_id"`
}
