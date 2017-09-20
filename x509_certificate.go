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

type X509Certificate struct {

	SubjectDistinguishedName string `json:"SubjectDistinguishedName,omitempty"`

	IssuerDistinguishedName string `json:"IssuerDistinguishedName,omitempty"`

	Thumbprint string `json:"Thumbprint,omitempty"`

	NotAfter time.Time `json:"NotAfter,omitempty"`

	NotBefore time.Time `json:"NotBefore,omitempty"`

	Version int32 `json:"Version,omitempty"`

	SerialNumber string `json:"SerialNumber,omitempty"`

	SignatureAlgorithmName string `json:"SignatureAlgorithmName,omitempty"`
}
