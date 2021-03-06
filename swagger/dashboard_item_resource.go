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

type DashboardItemResource struct {

	Id string `json:"Id,omitempty"`

	ProjectId string `json:"ProjectId,omitempty"`

	EnvironmentId string `json:"EnvironmentId,omitempty"`

	ReleaseId string `json:"ReleaseId,omitempty"`

	DeploymentId string `json:"DeploymentId,omitempty"`

	TaskId string `json:"TaskId,omitempty"`

	TenantId string `json:"TenantId,omitempty"`

	ChannelId string `json:"ChannelId,omitempty"`

	ReleaseVersion string `json:"ReleaseVersion,omitempty"`

	Created time.Time `json:"Created,omitempty"`

	QueueTime time.Time `json:"QueueTime,omitempty"`

	CompletedTime time.Time `json:"CompletedTime,omitempty"`

	State int32 `json:"State,omitempty"`

	HasPendingInterruptions bool `json:"HasPendingInterruptions,omitempty"`

	HasWarningsOrErrors bool `json:"HasWarningsOrErrors,omitempty"`

	ErrorMessage string `json:"ErrorMessage,omitempty"`

	Duration string `json:"Duration,omitempty"`

	IsCurrent bool `json:"IsCurrent,omitempty"`

	IsPrevious bool `json:"IsPrevious,omitempty"`

	IsCompleted bool `json:"IsCompleted,omitempty"`

	LastModifiedOn time.Time `json:"LastModifiedOn,omitempty"`

	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	Links map[string]string `json:"Links,omitempty"`
}
