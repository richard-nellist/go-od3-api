# \DeploymentsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeploymentsGet**](DeploymentsApi.md#ApiDeploymentsGet) | **Get** /api/deployments | 
[**ApiDeploymentsIdDelete**](DeploymentsApi.md#ApiDeploymentsIdDelete) | **Delete** /api/deployments/{id} | Delete a DeploymentResource by ID
[**ApiDeploymentsIdGet**](DeploymentsApi.md#ApiDeploymentsIdGet) | **Get** /api/deployments/{id} | Get a DeploymentResource by ID
[**ApiDeploymentsPost**](DeploymentsApi.md#ApiDeploymentsPost) | **Post** /api/deployments | 


# **ApiDeploymentsGet**
> ApiDeploymentsGet()



Lists all of the deployments in the current Octopus installation, from projects, releases and environments accessible by the current user. The results will be sorted from most recent to least recent deployment.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiDeploymentsIdDelete**
> TaskResource ApiDeploymentsIdDelete($id)

Delete a DeploymentResource by ID

Deletes a deployment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the DeploymentResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiDeploymentsIdGet**
> DeploymentResource ApiDeploymentsIdGet($id)

Get a DeploymentResource by ID

Gets a deployment by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the DeploymentResource to load | 

### Return type

[**DeploymentResource**](DeploymentResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiDeploymentsPost**
> DeploymentResource ApiDeploymentsPost()



Creates a new deployment.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

[**DeploymentResource**](DeploymentResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

