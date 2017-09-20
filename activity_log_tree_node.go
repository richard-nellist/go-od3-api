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

type ActivityLogTreeNode struct {

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Started time.Time `json:"Started,omitempty"`

	Ended time.Time `json:"Ended,omitempty"`

	Status int32 `json:"Status,omitempty"`

	ShowAtSummaryLevel bool `json:"ShowAtSummaryLevel,omitempty"`

	ProgressPercentage int32 `json:"ProgressPercentage,omitempty"`

	ProgressMessage string `json:"ProgressMessage,omitempty"`

	Children []ActivityLogTreeNode `json:"Children,omitempty"`

	LogElements []ActivityLogEntry `json:"LogElements,omitempty"`
}
