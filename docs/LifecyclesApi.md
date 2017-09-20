# \LifecyclesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiLifecyclesAllGet**](LifecyclesApi.md#ApiLifecyclesAllGet) | **Get** /api/lifecycles/all | Get a list of LifecycleResources
[**ApiLifecyclesGet**](LifecyclesApi.md#ApiLifecyclesGet) | **Get** /api/lifecycles | Get a list of LifecycleResources
[**ApiLifecyclesIdDelete**](LifecyclesApi.md#ApiLifecyclesIdDelete) | **Delete** /api/lifecycles/{id} | Delete a LifecycleResource by ID
[**ApiLifecyclesIdGet**](LifecyclesApi.md#ApiLifecyclesIdGet) | **Get** /api/lifecycles/{id} | Get a LifecycleResource by ID
[**ApiLifecyclesIdPreviewGet**](LifecyclesApi.md#ApiLifecyclesIdPreviewGet) | **Get** /api/lifecycles/{id}/preview | 
[**ApiLifecyclesIdProjectsGet**](LifecyclesApi.md#ApiLifecyclesIdProjectsGet) | **Get** /api/lifecycles/{id}/projects | 
[**ApiLifecyclesIdPut**](LifecyclesApi.md#ApiLifecyclesIdPut) | **Put** /api/lifecycles/{id} | Modify a LifecycleResource by ID
[**ApiLifecyclesPost**](LifecyclesApi.md#ApiLifecyclesPost) | **Post** /api/lifecycles | Create a LifecycleResource


# **ApiLifecyclesAllGet**
> []LifecycleResource ApiLifecyclesAllGet()

Get a list of LifecycleResources

Lists all the lifecycles in the Octopus Deploy installation.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]LifecycleResource**](LifecycleResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLifecyclesGet**
> ResourceCollectionLifecycleResource ApiLifecyclesGet($skip, $take)

Get a list of LifecycleResources

Lists all of the lifecycles in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionLifecycleResource**](ResourceCollection[LifecycleResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLifecyclesIdDelete**
> TaskResource ApiLifecyclesIdDelete($id)

Delete a LifecycleResource by ID

Deletes an existing lifecycle.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the LifecycleResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLifecyclesIdGet**
> LifecycleResource ApiLifecyclesIdGet($id)

Get a LifecycleResource by ID

Gets a single lifecycle by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the LifecycleResource to load | 

### Return type

[**LifecycleResource**](LifecycleResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLifecyclesIdPreviewGet**
> LifecycleResource ApiLifecyclesIdPreviewGet($id)



Gets a single lifecycle by ID, as a preview.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**LifecycleResource**](LifecycleResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLifecyclesIdProjectsGet**
> []ProjectResource ApiLifecyclesIdProjectsGet($id)



Gets a all projects that use this lifecycle.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**[]ProjectResource**](ProjectResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLifecyclesIdPut**
> LifecycleResource ApiLifecyclesIdPut($id, $lifecycleResource)

Modify a LifecycleResource by ID

Modifies an existing lifecycle.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the LifecycleResource to modify | 
 **lifecycleResource** | [**LifecycleResource**](LifecycleResource.md)| The LifecycleResource resource to create | [optional] 

### Return type

[**LifecycleResource**](LifecycleResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLifecyclesPost**
> LifecycleResource ApiLifecyclesPost($lifecycleResource)

Create a LifecycleResource

Creates a new lifecycle.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lifecycleResource** | [**LifecycleResource**](LifecycleResource.md)| The LifecycleResource resource to create | [optional] 

### Return type

[**LifecycleResource**](LifecycleResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

