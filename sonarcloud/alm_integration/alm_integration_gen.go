package alm_integration

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ProvisionProjectsRequest Provision project from ALM
type ProvisionProjectsRequest struct {
	InstallationKeys       string `form:"installationKeys,omitempty"`       // Comma separated list of repositories and their IDs..
	Organization           string `form:"organization,omitempty"`           // The key of the organization
	NewCodeDefinitionType  string `form:"newCodeDefinitionType,omitempty"`  // Configure type of new code definition. Can be 'days' or 'previous_version'.
	NewCodeDefinitionValue string `form:"newCodeDefinitionValue,omitempty"` // Configure value of new code definition. Can be number of days or 'previous_version'.
}

// ProvisionProjectsResponse is the response for ProvisionProjectsRequest
type ProvisionProjectsResponse struct {
	Projects []struct {
		ProjectKey string `json:"projectKey,omitempty"`
	} `json:"projects,omitempty"`
}

// ListRepositoriesRequest List repositories from ALM
type ListRepositoriesRequest struct {
	Organization string `form:"organization,omitempty"` // The key of the organization
}

// ListRepositoriesResponse is the response for ListRepositoriesRequest
type ListRepositoriesResponse struct {
	Repositories []struct {
		InstallationKey string `json:"installationKey,omitempty"`
		Label           string `json:"label,omitempty"`
		LinkedProjects  []struct {
			Key  string `json:"key,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"linkedProjects,omitempty"`
		Private bool   `json:"private,omitempty"`
		Slug    string `json:"slug,omitempty"`
	} `json:"repositories,omitempty"`
}

// ShowAppInfoRequest Gives information on the ALM application. E.g. GitHub or Azure Devops.
type ShowAppInfoRequest struct{}

// ShowAppInfoResponse is the response for ShowAppInfoRequest
type ShowAppInfoResponse struct {
	Application struct {
		BackgroundColor string `json:"backgroundColor,omitempty"`
		IconPath        string `json:"iconPath,omitempty"`
		InstallationUrl string `json:"installationUrl,omitempty"`
		Key             string `json:"key,omitempty"`
		Name            string `json:"name,omitempty"`
		WhiteIconPath   string `json:"whiteIconPath,omitempty"`
	} `json:"application,omitempty"`
}

// ShowBoundOrganizationRequest Show information on the ALM organization.
type ShowBoundOrganizationRequest struct {
	Organization string `form:"organization,omitempty"` // The key of the organization
}

// ShowBoundOrganizationResponse is the response for ShowBoundOrganizationRequest
type ShowBoundOrganizationResponse struct {
	AlmOrganization struct {
		AlmUrl      string `json:"almUrl,omitempty"`
		Avatar      string `json:"avatar,omitempty"`
		Description string `json:"description,omitempty"`
		Key         string `json:"key,omitempty"`
		Name        string `json:"name,omitempty"`
		Personal    bool   `json:"personal,omitempty"`
		Url         string `json:"url,omitempty"`
	} `json:"almOrganization,omitempty"`
}

// IsBoundToMonorepoRequest Returns whether project is bound to a monorepo.
type IsBoundToMonorepoRequest struct {
	Project string `form:"project,omitempty"` // The key of the project
}

// IsBoundToMonorepoResponse is the response for IsBoundToMonorepoRequest
type IsBoundToMonorepoResponse struct {
	IsBoundToMonorepo bool `json:"isBoundToMonorepo,omitempty"`
}
