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

type SubscriptionApi struct {
	Configuration *Configuration
}

func NewSubscriptionApi() *SubscriptionApi {
	configuration := NewConfiguration()
	return &SubscriptionApi{
		Configuration: configuration,
	}
}

func NewSubscriptionApiWithBasePath(basePath string) *SubscriptionApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &SubscriptionApi{
		Configuration: configuration,
	}
}

/**
 * Get a list of SubscriptionResources
 * Lists all the subscriptions in the Octopus Deploy installation.
 *
 * @return []SubscriptionResource
 */
func (a SubscriptionApi) ApiSubscriptionsAllGet() ([]SubscriptionResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/subscriptions/all"

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
	var successPayload = new([]SubscriptionResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiSubscriptionsAllGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Get a list of SubscriptionResources
 * Lists all of the subscriptions in the current Octopus installation. The results will be sorted alphabetically by name.
 *
 * @param skip Number of items to skip
 * @param take Number of items to take
 * @return *ResourceCollectionSubscriptionResource
 */
func (a SubscriptionApi) ApiSubscriptionsGet(skip int32, take int32) (*ResourceCollectionSubscriptionResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/subscriptions"

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
	var successPayload = new(ResourceCollectionSubscriptionResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiSubscriptionsGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Delete a SubscriptionResource by ID
 * Deletes an existing subscription.
 *
 * @param id ID of the SubscriptionResource to delete
 * @return *TaskResource
 */
func (a SubscriptionApi) ApiSubscriptionsIdDelete(id string) (*TaskResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/subscriptions/{id}"
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
	var localVarAPIResponse = &APIResponse{Operation: "ApiSubscriptionsIdDelete", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Get a SubscriptionResource by ID
 * Get a subscription
 *
 * @param id ID of the SubscriptionResource to load
 * @return *SubscriptionResource
 */
func (a SubscriptionApi) ApiSubscriptionsIdGet(id string) (*SubscriptionResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/subscriptions/{id}"
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
	var successPayload = new(SubscriptionResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiSubscriptionsIdGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Modify a SubscriptionResource by ID
 * Updates an existing subscription
 *
 * @param id ID of the SubscriptionResource to modify
 * @param subscriptionResource The SubscriptionResource resource to create
 * @return *SubscriptionResource
 */
func (a SubscriptionApi) ApiSubscriptionsIdPut(id string, subscriptionResource SubscriptionResource) (*SubscriptionResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/subscriptions/{id}"
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
	localVarPostBody = &subscriptionResource
	var successPayload = new(SubscriptionResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiSubscriptionsIdPut", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Create a SubscriptionResource
 * Creates a new subscription
 *
 * @param subscriptionResource The SubscriptionResource resource to create
 * @return *SubscriptionResource
 */
func (a SubscriptionApi) ApiSubscriptionsPost(subscriptionResource SubscriptionResource) (*SubscriptionResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/api/subscriptions"

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
	localVarPostBody = &subscriptionResource
	var successPayload = new(SubscriptionResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ApiSubscriptionsPost", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
