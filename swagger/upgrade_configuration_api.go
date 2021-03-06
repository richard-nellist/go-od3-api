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
	"net/url"
	"strings"
)

type UpgradeConfigurationApi struct {
	Configuration *Configuration
}

func NewUpgradeConfigurationApi() *UpgradeConfigurationApi {
	configuration := NewConfiguration()
	return &UpgradeConfigurationApi{
		Configuration: configuration,
	}
}

func NewUpgradeConfigurationApiWithBasePath(basePath string) *UpgradeConfigurationApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &UpgradeConfigurationApi{
		Configuration: configuration,
	}
}

/**
 * 
 * Gets information about the upgrade configuration in use by the Octopus Server.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
 *
 * @return void
 */
func (a UpgradeConfigurationApi) ApiUpgradeconfigurationGet() (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/upgradeconfiguration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(APIKeyHeader)' required
	// set key with prefix in header
	localVarHeaderParams["X-Octopus-ApiKey"] = a.Configuration.GetAPIKeyWithPrefix("X-Octopus-ApiKey")
	// authentication '(APIKeyQuery)' required
	// set key with prefix in query string
	localVarQueryParams["ApiKey"] =  a.Configuration.GetAPIKeyWithPrefix("ApiKey")
	// authentication '(NugetApiKeyHeader)' required
	// set key with prefix in header
	localVarHeaderParams["X-NuGet-ApiKey"] = a.Configuration.GetAPIKeyWithPrefix("X-NuGet-ApiKey")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiUpgradeconfigurationGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

/**
 * 
 * Updates the upgrade configuration used by the Octopus Server.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
 *
 * @return void
 */
func (a UpgradeConfigurationApi) ApiUpgradeconfigurationPut() (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/upgradeconfiguration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(APIKeyHeader)' required
	// set key with prefix in header
	localVarHeaderParams["X-Octopus-ApiKey"] = a.Configuration.GetAPIKeyWithPrefix("X-Octopus-ApiKey")
	// authentication '(APIKeyQuery)' required
	// set key with prefix in query string
	localVarQueryParams["ApiKey"] =  a.Configuration.GetAPIKeyWithPrefix("ApiKey")
	// authentication '(NugetApiKeyHeader)' required
	// set key with prefix in header
	localVarHeaderParams["X-NuGet-ApiKey"] = a.Configuration.GetAPIKeyWithPrefix("X-NuGet-ApiKey")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiUpgradeconfigurationPut", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

