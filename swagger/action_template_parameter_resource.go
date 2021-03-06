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

type ActionTemplateParameterResource struct {

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Label string `json:"Label,omitempty"`

	HelpText string `json:"HelpText,omitempty"`

	DefaultValue PropertyValueResource `json:"DefaultValue,omitempty"`

	DisplaySettings map[string]string `json:"DisplaySettings,omitempty"`

	LastModifiedOn time.Time `json:"LastModifiedOn,omitempty"`

	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	Links map[string]string `json:"Links,omitempty"`
}
