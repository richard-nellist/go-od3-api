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

type TaskResource struct {

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Arguments map[string]interface{} `json:"Arguments,omitempty"`

	State int32 `json:"State,omitempty"`

	Completed string `json:"Completed,omitempty"`

	QueueTime time.Time `json:"QueueTime,omitempty"`

	QueueTimeExpiry time.Time `json:"QueueTimeExpiry,omitempty"`

	StartTime time.Time `json:"StartTime,omitempty"`

	LastUpdatedTime time.Time `json:"LastUpdatedTime,omitempty"`

	CompletedTime time.Time `json:"CompletedTime,omitempty"`

	ServerNode string `json:"ServerNode,omitempty"`

	Duration string `json:"Duration,omitempty"`

	ErrorMessage string `json:"ErrorMessage,omitempty"`

	HasBeenPickedUpByProcessor bool `json:"HasBeenPickedUpByProcessor,omitempty"`

	IsCompleted bool `json:"IsCompleted,omitempty"`

	FinishedSuccessfully bool `json:"FinishedSuccessfully,omitempty"`

	HasPendingInterruptions bool `json:"HasPendingInterruptions,omitempty"`

	CanRerun bool `json:"CanRerun,omitempty"`

	HasWarningsOrErrors bool `json:"HasWarningsOrErrors,omitempty"`

	LastModifiedOn time.Time `json:"LastModifiedOn,omitempty"`

	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	Links map[string]string `json:"Links,omitempty"`
}