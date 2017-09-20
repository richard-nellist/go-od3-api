# \ProjectTriggersApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectsIdTriggersGet**](ProjectTriggersApi.md#ApiProjectsIdTriggersGet) | **Get** /api/projects/{id}/triggers | Get a list of ProjectTriggerResources
[**ApiProjecttriggersGet**](ProjectTriggersApi.md#ApiProjecttriggersGet) | **Get** /api/projecttriggers | Get a list of ProjectTriggerResources
[**ApiProjecttriggersIdDelete**](ProjectTriggersApi.md#ApiProjecttriggersIdDelete) | **Delete** /api/projecttriggers/{id} | Delete a ProjectTriggerResource by ID
[**ApiProjecttriggersIdGet**](ProjectTriggersApi.md#ApiProjecttriggersIdGet) | **Get** /api/projecttriggers/{id} | Get a ProjectTriggerResource by ID
[**ApiProjecttriggersIdPut**](ProjectTriggersApi.md#ApiProjecttriggersIdPut) | **Put** /api/projecttriggers/{id} | Modify a ProjectTriggerResource by ID
[**ApiProjecttriggersPost**](ProjectTriggersApi.md#ApiProjecttriggersPost) | **Post** /api/projecttriggers | Create a ProjectTriggerResource


# **ApiProjectsIdTriggersGet**
> ResourceCollectionProjectResource ApiProjectsIdTriggersGet($skip, $take)

Get a list of ProjectTriggerResources

Lists all the project triggers for the given project


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionProjectResource**](ResourceCollection[ProjectResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjecttriggersGet**
> ResourceCollectionProjectTriggerResource ApiProjecttriggersGet($skip, $take)

Get a list of ProjectTriggerResources

Gets all the project triggers in the current Octopus installation, sorted by Id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionProjectTriggerResource**](ResourceCollection[ProjectTriggerResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjecttriggersIdDelete**
> TaskResource ApiProjecttriggersIdDelete($id)

Delete a ProjectTriggerResource by ID

Deletes an existing project trigger.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ProjectTriggerResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjecttriggersIdGet**
> ProjectTriggerResource ApiProjecttriggersIdGet($id)

Get a ProjectTriggerResource by ID

Get a project trigger


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ProjectTriggerResource to load | 

### Return type

[**ProjectTriggerResource**](ProjectTriggerResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjecttriggersIdPut**
> ProjectTriggerResource ApiProjecttriggersIdPut($id, $projectTriggerResource)

Modify a ProjectTriggerResource by ID

Updates an existing project trigger


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ProjectTriggerResource to modify | 
 **projectTriggerResource** | [**ProjectTriggerResource**](ProjectTriggerResource.md)| The ProjectTriggerResource resource to create | [optional] 

### Return type

[**ProjectTriggerResource**](ProjectTriggerResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjecttriggersPost**
> ProjectTriggerResource ApiProjecttriggersPost($projectTriggerResource)

Create a ProjectTriggerResource

Creates a new project trigger


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectTriggerResource** | [**ProjectTriggerResource**](ProjectTriggerResource.md)| The ProjectTriggerResource resource to create | [optional] 

### Return type

[**ProjectTriggerResource**](ProjectTriggerResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

