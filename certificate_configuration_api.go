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
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type CertificateConfigurationApi struct {
	Configuration *Configuration
}

func NewCertificateConfigurationApi() *CertificateConfigurationApi {
	configuration := NewConfiguration()
	return &CertificateConfigurationApi{
		Configuration: configuration,
	}
}

func NewCertificateConfigurationApiWithBasePath(basePath string) *CertificateConfigurationApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &CertificateConfigurationApi{
		Configuration: configuration,
	}
}

/**
 * Get a CertificateConfigurationResource by ID
 * Gets a certificate by ID.
 *
 * @param id ID of the CertificateConfigurationResource to load
 * @return *CertificateConfigurationResource
 */
func (a CertificateConfigurationApi) ApiCertificatesIdICertificateGlobalcertificateAzureGet(id string) (*CertificateConfigurationResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/certificates/(?<id>(?i)^(certificate-global|certificate-azure))"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

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
	var successPayload = new(CertificateConfigurationResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiCertificatesIdICertificateGlobalcertificateAzureGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Get a list of CertificateConfigurationResources
 * Lists all of the X509 certificates in the current Octopus installation.
 *
 * @param skip Number of items to skip
 * @param take Number of items to take
 * @return *ResourceCollectionCertificateConfigurationResource
 */
func (a CertificateConfigurationApi) ApiConfigurationCertificatesGet(skip int32, take int32) (*ResourceCollectionCertificateConfigurationResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/configuration/certificates"

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
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))
	localVarQueryParams.Add("take", a.Configuration.APIClient.ParameterToString(take, ""))

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
	var successPayload = new(ResourceCollectionCertificateConfigurationResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiConfigurationCertificatesGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Get a CertificateConfigurationResource by ID
 * Gets a certificate by ID.
 *
 * @param id ID of the CertificateConfigurationResource to load
 * @return *CertificateConfigurationResource
 */
func (a CertificateConfigurationApi) ApiConfigurationCertificatesIdGet(id string) (*CertificateConfigurationResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/configuration/certificates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

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
	var successPayload = new(CertificateConfigurationResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiConfigurationCertificatesIdGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * Downloads the public portion of the certificate in .cer format  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
 *
 * @param id ID of the resource
 * @return **os.File
 */
func (a CertificateConfigurationApi) ApiConfigurationCertificatesIdPublicCerGet(id string) (**os.File, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/configuration/certificates/{id}/public-cer"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

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
		"application/x-x509-ca-cert",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(*os.File)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiConfigurationCertificatesIdPublicCerGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

