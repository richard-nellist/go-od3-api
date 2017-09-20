# \DeploymentProcessesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeploymentprocessesGet**](DeploymentProcessesApi.md#ApiDeploymentprocessesGet) | **Get** /api/deploymentprocesses | Get a list of DeploymentProcessResources
[**ApiDeploymentprocessesIdGet**](DeploymentProcessesApi.md#ApiDeploymentprocessesIdGet) | **Get** /api/deploymentprocesses/{id} | Get a DeploymentProcessResource by ID
[**ApiDeploymentprocessesIdPut**](DeploymentProcessesApi.md#ApiDeploymentprocessesIdPut) | **Put** /api/deploymentprocesses/{id} | 
[**ApiDeploymentprocessesIdTemplateGet**](DeploymentProcessesApi.md#ApiDeploymentprocessesIdTemplateGet) | **Get** /api/deploymentprocesses/{id}/template | 


# **ApiDeploymentprocessesGet**
> ResourceCollectionDeploymentProcessResource ApiDeploymentprocessesGet($skip, $take)

Get a list of DeploymentProcessResources

Lists all the deployment processeses in the current Octopus installation, sorted by Id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionDeploymentProcessResource**](ResourceCollection[DeploymentProcessResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiDeploymentprocessesIdGet**
> DeploymentProcessResource ApiDeploymentprocessesIdGet($id)

Get a DeploymentProcessResource by ID

Gets a deployment process by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the DeploymentProcessResource to load | 

### Return type

[**DeploymentProcessResource**](DeploymentProcessResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiDeploymentprocessesIdPut**
> DeploymentProcessResource ApiDeploymentprocessesIdPut($id)



Modifies a deployment process. Only allowed for deployment processes owned by a project (cannot be used to change the deployment process owned by a release).  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**DeploymentProcessResource**](DeploymentProcessResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiDeploymentprocessesIdTemplateGet**
> ReleaseTemplateResource ApiDeploymentprocessesIdTemplateGet($id)



Gets all of the information necessary for creating or editing a release using this deployment process.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**ReleaseTemplateResource**](ReleaseTemplateResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

