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

type DeploymentProcessResource struct {

	Id string `json:"Id,omitempty"`

	ProjectId string `json:"ProjectId,omitempty"`

	Steps []DeploymentStepResource `json:"Steps,omitempty"`

	Version int32 `json:"Version"`

	LastSnapshotId string `json:"LastSnapshotId,omitempty"`

	LastModifiedOn time.Time `json:"LastModifiedOn,omitempty"`

	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	Links map[string]string `json:"Links,omitempty"`
}