package infra_connector

type CertificateCreateRequest struct {
	Name           string   `json:"name"`
	DomainName     string   `json:"domainName"`
	Provider       string   `json:"provider"`
	Certificate    string   `json:"certificate"`
	CertificateCA  string   `json:"certificateCA"`
	CertificateKey string   `json:"certificateKey"`
	Shared         bool     `json:"shared"`
	Tags           []string `json:"tags"`
}

type CertificateUpdateRequest struct {
	Certificate    string `json:"certificate"`
	CertificateCA  string `json:"certificateCA"`
	CertificateKey string `json:"certificateKey"`
}
