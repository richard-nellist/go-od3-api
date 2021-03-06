/* 
 * Octopus Server API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 3.17.1+Branch.master.Sha.434caf20746e16780a8fab99f2fd4f4894a7283e
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import (
	"time"
)

type ProjectResource struct {

	Id string `json:"Id,omitempty"`

	VariableSetId string `json:"VariableSetId,omitempty"`

	DeploymentProcessId string `json:"DeploymentProcessId,omitempty"`

	DiscreteChannelRelease bool `json:"DiscreteChannelRelease,omitempty"`

	IncludedLibraryVariableSetIds []string `json:"IncludedLibraryVariableSetIds,omitempty"`

	DefaultToSkipIfAlreadyInstalled bool `json:"DefaultToSkipIfAlreadyInstalled,omitempty"`

	TenantedDeploymentMode int32 `json:"TenantedDeploymentMode,omitempty"`

	VersioningStrategy VersioningStrategyResource `json:"VersioningStrategy,omitempty"`

	ReleaseCreationStrategy ReleaseCreationStrategyResource `json:"ReleaseCreationStrategy,omitempty"`

	Templates []ActionTemplateParameterResource `json:"Templates,omitempty"`

	AutoDeployReleaseOverrides []AutoDeployReleaseOverrideResource `json:"AutoDeployReleaseOverrides,omitempty"`

	Name string `json:"Name,omitempty"`

	Slug string `json:"Slug,omitempty"`

	Description string `json:"Description,omitempty"`

	IsDisabled bool `json:"IsDisabled,omitempty"`

	ProjectGroupId string `json:"ProjectGroupId,omitempty"`

	LifecycleId string `json:"LifecycleId,omitempty"`

	AutoCreateRelease bool `json:"AutoCreateRelease,omitempty"`

	ProjectConnectivityPolicy ProjectConnectivityPolicy `json:"ProjectConnectivityPolicy,omitempty"`

	LastModifiedOn time.Time `json:"LastModifiedOn,omitempty"`

	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	Links map[string]string `json:"Links,omitempty"`
}
