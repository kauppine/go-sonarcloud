package dop_translation

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// THIS HAS BEEN MODIFIED SINCE GENERATION TO SUPPORT BOTH GET AND POST REQUESTS

// ProjectBindingsRequest Bind ALM project to repo
type ProjectBindingsRequest struct {
	ProjectId    string `form:"projectId,omitempty"`    // SonarCloud project ID, not project key.
	RepositoryId string `form:"repositoryId,omitempty"` // ALM repository ID, in GitHub it is an integer.
}

// ProjectBindingsbindingIdRequest Update bindings for a repo
type ProjectBindingsbindingIdRequest struct {
	ProjectId string `form:"projectId,omitempty"` // SonarCloud project ID, not project key.
}
