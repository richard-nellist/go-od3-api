# \ProjectGroupsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectgroupsAllGet**](ProjectGroupsApi.md#ApiProjectgroupsAllGet) | **Get** /api/projectgroups/all | Get a list of ProjectGroupResources
[**ApiProjectgroupsGet**](ProjectGroupsApi.md#ApiProjectgroupsGet) | **Get** /api/projectgroups | Get a list of ProjectGroupResources
[**ApiProjectgroupsIdDelete**](ProjectGroupsApi.md#ApiProjectgroupsIdDelete) | **Delete** /api/projectgroups/{id} | Delete a ProjectGroupResource by ID
[**ApiProjectgroupsIdGet**](ProjectGroupsApi.md#ApiProjectgroupsIdGet) | **Get** /api/projectgroups/{id} | Get a ProjectGroupResource by ID
[**ApiProjectgroupsIdProjectsGet**](ProjectGroupsApi.md#ApiProjectgroupsIdProjectsGet) | **Get** /api/projectgroups/{id}/projects | Get a list of ProjectResources
[**ApiProjectgroupsIdPut**](ProjectGroupsApi.md#ApiProjectgroupsIdPut) | **Put** /api/projectgroups/{id} | Modify a ProjectGroupResource by ID
[**ApiProjectgroupsPost**](ProjectGroupsApi.md#ApiProjectgroupsPost) | **Post** /api/projectgroups | Create a ProjectGroupResource


# **ApiProjectgroupsAllGet**
> []ProjectGroupResource ApiProjectgroupsAllGet()

Get a list of ProjectGroupResources

Lists the name and ID of all of the project groups in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProjectGroupResource**](ProjectGroupResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectgroupsGet**
> ResourceCollectionProjectGroupResource ApiProjectgroupsGet($skip, $take)

Get a list of ProjectGroupResources

Lists all of the project groups in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionProjectGroupResource**](ResourceCollection[ProjectGroupResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectgroupsIdDelete**
> TaskResource ApiProjectgroupsIdDelete($id)

Delete a ProjectGroupResource by ID

Deletes an existing project group.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ProjectGroupResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectgroupsIdGet**
> ProjectGroupResource ApiProjectgroupsIdGet($id)

Get a ProjectGroupResource by ID

Gets a single project group by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ProjectGroupResource to load | 

### Return type

[**ProjectGroupResource**](ProjectGroupResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectgroupsIdProjectsGet**
> ResourceCollectionProjectGroupResource ApiProjectgroupsIdProjectsGet($skip, $take)

Get a list of ProjectResources

Lists all of the projects that belong to the given project group.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionProjectGroupResource**](ResourceCollection[ProjectGroupResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectgroupsIdPut**
> ProjectGroupResource ApiProjectgroupsIdPut($id, $projectGroupResource)

Modify a ProjectGroupResource by ID

Modifies an existing project group.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ProjectGroupResource to modify | 
 **projectGroupResource** | [**ProjectGroupResource**](ProjectGroupResource.md)| The ProjectGroupResource resource to create | [optional] 

### Return type

[**ProjectGroupResource**](ProjectGroupResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectgroupsPost**
> ProjectGroupResource ApiProjectgroupsPost($projectGroupResource)

Create a ProjectGroupResource

Creates a new project group.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectGroupResource** | [**ProjectGroupResource**](ProjectGroupResource.md)| The ProjectGroupResource resource to create | [optional] 

### Return type

[**ProjectGroupResource**](ProjectGroupResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

