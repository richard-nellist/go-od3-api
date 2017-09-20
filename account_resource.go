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

type AccountResource struct {

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	EnvironmentIds []string `json:"EnvironmentIds,omitempty"`

	TenantedDeploymentParticipation int32 `json:"TenantedDeploymentParticipation,omitempty"`

	TenantIds []string `json:"TenantIds,omitempty"`

	TenantTags []string `json:"TenantTags,omitempty"`

	AccountType int32 `json:"AccountType,omitempty"`

	LastModifiedOn time.Time `json:"LastModifiedOn,omitempty"`

	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	Links map[string]string `json:"Links,omitempty"`
}