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

type VariableScopeValues struct {

	Environments []ReferenceDataItem `json:"Environments,omitempty"`

	Machines []ReferenceDataItem `json:"Machines,omitempty"`

	Actions []ReferenceDataItem `json:"Actions,omitempty"`

	Roles []ReferenceDataItem `json:"Roles,omitempty"`

	Channels []ReferenceDataItem `json:"Channels,omitempty"`

	TenantTags []ReferenceDataItem `json:"TenantTags,omitempty"`
}
