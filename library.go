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

type Library struct {

	LibraryVariableSetId string `json:"LibraryVariableSetId,omitempty"`

	LibraryVariableSetName string `json:"LibraryVariableSetName,omitempty"`

	Templates []ActionTemplateParameterResource `json:"Templates,omitempty"`

	Variables map[string]PropertyValueResource `json:"Variables,omitempty"`

	Links map[string]string `json:"Links,omitempty"`
}
