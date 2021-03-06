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
	"encoding/json"
	"fmt"
)

type ProjectGroupsApi struct {
	Configuration *Configuration
}

func NewProjectGroupsApi() *ProjectGroupsApi {
	configuration := NewConfiguration()
	return &ProjectGroupsApi{
		Configuration: configuration,
	}
}

func NewProjectGroupsApiWithBasePath(basePath string) *ProjectGroupsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &ProjectGroupsApi{
		Configuration: configuration,
	}
}

/**
 * Get a list of ProjectGroupResources
 * Lists the name and ID of all of the project groups in the current Octopus installation. The results will be sorted alphabetically by name.
 *
 * @return []ProjectGroupResource
 */
func (a ProjectGroupsApi) ApiProjectgroupsAllGet() ([]ProjectGroupResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/projectgroups/all"

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
	var successPayload = new([]ProjectGroupResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiProjectgroupsAllGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

/**
 * Get a list of ProjectGroupResources
 * Lists all of the project groups in the current Octopus installation. The results will be sorted alphabetically by name.
 *
 * @param skip Number of items to skip
 * @param take Number of items to take
 * @return *ResourceCollectionProjectGroupResource
 */
func (a ProjectGroupsApi) ApiProjectgroupsGet(skip int32, take int32) (*ResourceCollectionProjectGroupResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/projectgroups"

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
	var successPayload = new(ResourceCollectionProjectGroupResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiProjectgroupsGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Delete a ProjectGroupResource by ID
 * Deletes an existing project group.
 *
 * @param id ID of the ProjectGroupResource to delete
 * @return *TaskResource
 */
func (a ProjectGroupsApi) ApiProjectgroupsIdDelete(id string) (*TaskResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/projectgroups/{id}"
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
	var successPayload = new(TaskResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiProjectgroupsIdDelete", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Get a ProjectGroupResource by ID
 * Gets a single project group by ID.
 *
 * @param id ID of the ProjectGroupResource to load
 * @return *ProjectGroupResource
 */
func (a ProjectGroupsApi) ApiProjectgroupsIdGet(id string) (*ProjectGroupResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/projectgroups/{id}"
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
	var successPayload = new(ProjectGroupResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiProjectgroupsIdGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Get a list of ProjectResources
 * Lists all of the projects that belong to the given project group.
 *
 * @param skip Number of items to skip
 * @param take Number of items to take
 * @return *ResourceCollectionProjectGroupResource
 */
func (a ProjectGroupsApi) ApiProjectgroupsIdProjectsGet(skip int32, take int32) (*ResourceCollectionProjectGroupResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/projectgroups/{id}/projects"

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
	var successPayload = new(ResourceCollectionProjectGroupResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiProjectgroupsIdProjectsGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Modify a ProjectGroupResource by ID
 * Modifies an existing project group.
 *
 * @param id ID of the ProjectGroupResource to modify
 * @param projectGroupResource The ProjectGroupResource resource to create
 * @return *ProjectGroupResource
 */
func (a ProjectGroupsApi) ApiProjectgroupsIdPut(id string, projectGroupResource ProjectGroupResource) (*ProjectGroupResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/projectgroups/{id}"
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
	// body params
	localVarPostBody = &projectGroupResource
	var successPayload = new(ProjectGroupResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiProjectgroupsIdPut", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Create a ProjectGroupResource
 * Creates a new project group.
 *
 * @param projectGroupResource The ProjectGroupResource resource to create
 * @return *ProjectGroupResource
 */
func (a ProjectGroupsApi) ApiProjectgroupsPost(projectGroupResource ProjectGroupResource) (*ProjectGroupResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/projectgroups"

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
	// body params
	localVarPostBody = &projectGroupResource
	var successPayload = new(ProjectGroupResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiProjectgroupsPost", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

