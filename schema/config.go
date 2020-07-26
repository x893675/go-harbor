package schema

// HTTPAuthProxy wraps the settings for HTTP auth proxy
type HTTPAuthProxy struct {
	Endpoint            string `json:"endpoint"`
	TokenReviewEndpoint string `json:"tokenreivew_endpoint"`
	VerifyCert          bool   `json:"verify_cert"`
	SkipSearch          bool   `json:"skip_search"`
	ServerCertificate   string `json:"server_certificate"`
}
