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

type IdentityResource struct {

	IdentityProviderName string `json:"IdentityProviderName,omitempty"`

	Claims map[string]IdentityClaimResource `json:"Claims,omitempty"`
}
