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

type CommunityActionTemplateResource struct {

	Id string `json:"Id,omitempty"`

	Type_ string `json:"Type,omitempty"`

	Name string `json:"Name,omitempty"`

	Author string `json:"Author,omitempty"`

	Description string `json:"Description,omitempty"`

	Version int32 `json:"Version,omitempty"`

	Website string `json:"Website,omitempty"`

	HistoryUrl string `json:"HistoryUrl,omitempty"`

	Properties map[string]PropertyValueResource `json:"Properties,omitempty"`

	Parameters []ActionTemplateParameterResource `json:"Parameters,omitempty"`

	Links map[string]string `json:"Links,omitempty"`
}
