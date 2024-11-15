package autoscan

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// EligibilityRequest Check if automated scan can be enabled
type EligibilityRequest struct {
	AutoEnable  string `form:"autoEnable,omitempty"`  // Current autoenable settings
	IgnoreCache string `form:"ignoreCache,omitempty"` // ignore cache
	ProjectKey  string `form:"projectKey,omitempty"`  // SonarCloud project key
}

// EligibilityResponse is the response for EligibilityRequest
type EligibilityResponse struct {
	Eligible  bool `json:"eligible,omitempty"`
	Languages []struct {
		Language               string  `json:"language,omitempty"`
		PercentageOfRepository float64 `json:"percentageOfRepository,omitempty"`
		Supported              bool    `json:"supported,omitempty"`
	} `json:"languages,omitempty"`
}

// ActivationRequest Activate automated scan
type ActivationRequest struct {
	AutoEnable string `form:"autoEnable,omitempty"` // Filter the projects for which last analysis is older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided.
	ProjectKey string `form:"projectKey,omitempty"` // SonarCloud project key
}

// EligibilityStatusRequest Check eligibility status for projects
type EligibilityStatusRequest struct {
	ProjectKeys string `form:"projectKeys,omitempty"` // Comma separated list of project keys.
}
