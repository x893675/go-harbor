package schema

type GeneralInfo struct {
	WithNotary                  bool          `json:"with_notary"`
	AuthMode                    string        `json:"auth_mode"`
	AuthProxySettings           HTTPAuthProxy `json:"authproxy_settings,omitempty"`
	RegistryURL                 string        `json:"registry_url"`
	ExtURL                      string        `json:"external_url"`
	ProjectCreationRestrict     string        `json:"project_creation_restriction"`
	SelfRegistration            bool          `json:"self_registration"`
	HasCARoot                   bool          `json:"has_ca_root"`
	HarborVersion               string        `json:"harbor_version"`
	RegistryStorageProviderName string        `json:"registry_storage_provider_name"`
	ReadOnly                    bool          `json:"read_only"`
	WithChartMuseum             bool          `json:"with_chartmuseum"`
	NotificationEnable          bool          `json:"notification_enable"`
}
