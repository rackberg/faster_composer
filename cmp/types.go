package cmp

type ComposerInfo struct {
	Repositories []RepositoryInfo `json:"repositories"`
	Requirements map[string]string `json:"require"`
	DevRequirements map[string]string `json:"require-dev"`
}

type RepositoryInfo struct {
	Type string `json:"type"`
	Url string `json:"url"`
}

type PackagistInfo struct {
	Notify string `json:"notify"`
	NotifyBatch string `json:"notify-batch"`
	ProvidersUrl string `json:"providers-url"`
	Search string `json:"search"`
	ProviderIncludes map[string]ProviderInclude `json:"provider-includes"`
}

type ProviderInclude struct {
	Sha256 string `json:"sha256"`
}