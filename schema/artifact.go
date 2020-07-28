package schema

import (
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"time"
)

type Artifact struct {
	ScanOverview      map[string]*NativeReportSummary `json:"scan_overview,omitempty"`
	Tags              []*Tag                          `json:"tags,omitempty"`
	Labels            []*Label                        `json:"labels,omitempty"`
	ID                int64                           `json:"id"`
	Type              string                          `json:"type"`                // image, chart, etc
	MediaType         string                          `json:"media_type"`          // the media type of artifact. Mostly, it's the value of `manifest.config.mediatype`
	ManifestMediaType string                          `json:"manifest_media_type"` // the media type of manifest/index
	ProjectID         int64                           `json:"project_id"`
	RepositoryID      int64                           `json:"repository_id"`
	Digest            string                          `json:"digest"`
	Size              int64                           `json:"size"`
	PushTime          time.Time                       `json:"push_time"`
	PullTime          time.Time                       `json:"pull_time"`
	ExtraAttrs        map[string]interface{}          `json:"extra_attrs"` // only contains the simple attributes specific for the different artifact type, most of them should come from the config layer
	Annotations       map[string]string               `json:"annotations"`
	References        []*Reference                    `json:"references"` // child artifacts referenced by the parent artifact if the artifact is an index
	AdditionLinks     map[string]*AdditionLink        `json:"addition_links"`
}

// AdditionLink is a link via that the addition can be fetched
type AdditionLink struct {
	HREF     string `json:"href"`
	Absolute bool   `json:"absolute"` // specify the href is an absolute URL or not
}

type Severity string

type NativeReportSummary struct {
	ReportID        string                `json:"report_id"`
	ScanStatus      string                `json:"scan_status"`
	Severity        Severity              `json:"severity"`
	Duration        int64                 `json:"duration"`
	Summary         *VulnerabilitySummary `json:"summary"`
	StartTime       time.Time             `json:"start_time"`
	EndTime         time.Time             `json:"end_time"`
	Scanner         Scanner               `json:"scanner,omitempty"`
	CompletePercent int                   `json:"complete_percent"`
}

type Scanner struct {
	// The name of the scanner.
	Name string `json:"name"`
	// The name of the scanner's provider.
	Vendor string `json:"vendor"`
	// The version of the scanner.
	Version string `json:"version"`
}

type SeveritySummary map[Severity]int

type VulnerabilitySummary struct {
	Total   int             `json:"total"`
	Fixable int             `json:"fixable"`
	Summary SeveritySummary `json:"summary"`
}

type Tag struct {
	RepositoryId int64     `json:"repository_id"`
	Name         string    `json:"name,omitempty"`
	PushTime     time.Time `json:"push_time,omitempty"`
	PullTime     time.Time `json:"pull_time	,omitempty"`
	Signed       bool      `json:"signed"`
	Id           int64     `json:"id"`
	Immutable    bool      `json:"immutable"`
	ArtifactId   int64     `json:"artifact_id"`
}

type Label struct {
	UpdateTime   time.Time `json:"update_time,omitempty"`
	Description  string    `json:"description,omitempty"`
	Color        string    `json:"color,omitempty"`
	CreationTime time.Time `json:"creation_time,omitempty"`
	Deleted      bool      `json:"deleted"`
	Scope        string    `json:"scope,omitempty"`
	ProjectId    int64     `json:"project_id"`
	Id           int64     `json:"id"`
	Name         string    `json:"name,omitempty"`
}

type Reference struct {
	ID          int64             `json:"id"`
	ParentID    int64             `json:"parent_id"`
	ChildID     int64             `json:"child_id"`
	ChildDigest string            `json:"child_digest"`
	Platform    *v1.Platform      `json:"platform"`
	URLs        []string          `json:"urls"`
	Annotations map[string]string `json:"annotations"`
}

type ArtifactsListOptions struct {
	ProjectName         string
	RepositoryName      string
	QueryString         string
	Page                string
	PageSize            string
	WithTag             *bool
	WithLabel           *bool
	WithScanOverview    *bool
	WithSignature       *bool
	WithImmutableStatus *bool
}
