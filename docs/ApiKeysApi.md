# \ApiKeysApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiUsersUserIdApikeysGet**](ApiKeysApi.md#ApiUsersUserIdApikeysGet) | **Get** /api/users/{userId}/apikeys | 
[**ApiUsersUserIdApikeysIdDelete**](ApiKeysApi.md#ApiUsersUserIdApikeysIdDelete) | **Delete** /api/users/{userId}/apikeys/{id} | Delete a ApiKeyResource by ID
[**ApiUsersUserIdApikeysIdGet**](ApiKeysApi.md#ApiUsersUserIdApikeysIdGet) | **Get** /api/users/{userId}/apikeys/{id} | Get a ApiKeyResource by ID
[**ApiUsersUserIdApikeysPost**](ApiKeysApi.md#ApiUsersUserIdApikeysPost) | **Post** /api/users/{userId}/apikeys | 


# **ApiUsersUserIdApikeysGet**
> ApiUsersUserIdApikeysGet()



Lists all API keys for a user, returning the most recent results first.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersUserIdApikeysIdDelete**
> TaskResource ApiUsersUserIdApikeysIdDelete($id)

Delete a ApiKeyResource by ID

Revokes an existing API key.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ApiKeyResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersUserIdApikeysIdGet**
> ApiKeyResource ApiUsersUserIdApikeysIdGet($id)

Get a ApiKeyResource by ID

Gets a API key by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ApiKeyResource to load | 

### Return type

[**ApiKeyResource**](ApiKeyResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersUserIdApikeysPost**
> ApiKeyResource ApiUsersUserIdApikeysPost($userId)



Generates a new API key for the specified user. The API key returned in the result must be saved by the caller, as it cannot be retrieved subsequently from the Octopus server.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| ID of the user | 

### Return type

[**ApiKeyResource**](ApiKeyResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

