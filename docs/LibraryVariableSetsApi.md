# \LibraryVariableSetsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiLibraryvariablesetsAllGet**](LibraryVariableSetsApi.md#ApiLibraryvariablesetsAllGet) | **Get** /api/libraryvariablesets/all | Get a list of LibraryVariableSetResources
[**ApiLibraryvariablesetsGet**](LibraryVariableSetsApi.md#ApiLibraryvariablesetsGet) | **Get** /api/libraryvariablesets | Get a list of LibraryVariableSetResources
[**ApiLibraryvariablesetsIdDelete**](LibraryVariableSetsApi.md#ApiLibraryvariablesetsIdDelete) | **Delete** /api/libraryvariablesets/{id} | Delete a LibraryVariableSetResource by ID
[**ApiLibraryvariablesetsIdGet**](LibraryVariableSetsApi.md#ApiLibraryvariablesetsIdGet) | **Get** /api/libraryvariablesets/{id} | Get a LibraryVariableSetResource by ID
[**ApiLibraryvariablesetsIdPut**](LibraryVariableSetsApi.md#ApiLibraryvariablesetsIdPut) | **Put** /api/libraryvariablesets/{id} | Modify a LibraryVariableSetResource by ID
[**ApiLibraryvariablesetsPost**](LibraryVariableSetsApi.md#ApiLibraryvariablesetsPost) | **Post** /api/libraryvariablesets | Create a LibraryVariableSetResource


# **ApiLibraryvariablesetsAllGet**
> []LibraryVariableSetResource ApiLibraryvariablesetsAllGet()

Get a list of LibraryVariableSetResources

Lists all the library variable sets in the Octopus Deploy installation. The results will be sorted alphabetically by name.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]LibraryVariableSetResource**](LibraryVariableSetResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLibraryvariablesetsGet**
> ResourceCollectionLibraryVariableSetResource ApiLibraryvariablesetsGet($skip, $take)

Get a list of LibraryVariableSetResources

Lists all of the library variable sets in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionLibraryVariableSetResource**](ResourceCollection[LibraryVariableSetResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLibraryvariablesetsIdDelete**
> TaskResource ApiLibraryvariablesetsIdDelete($id)

Delete a LibraryVariableSetResource by ID

Deletes an existing library variable set.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the LibraryVariableSetResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLibraryvariablesetsIdGet**
> LibraryVariableSetResource ApiLibraryvariablesetsIdGet($id)

Get a LibraryVariableSetResource by ID

Gets a single library variable set by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the LibraryVariableSetResource to load | 

### Return type

[**LibraryVariableSetResource**](LibraryVariableSetResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLibraryvariablesetsIdPut**
> LibraryVariableSetResource ApiLibraryvariablesetsIdPut($id, $libraryVariableSetResource)

Modify a LibraryVariableSetResource by ID

Modifies an existing library variable set.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the LibraryVariableSetResource to modify | 
 **libraryVariableSetResource** | [**LibraryVariableSetResource**](LibraryVariableSetResource.md)| The LibraryVariableSetResource resource to create | [optional] 

### Return type

[**LibraryVariableSetResource**](LibraryVariableSetResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLibraryvariablesetsPost**
> LibraryVariableSetResource ApiLibraryvariablesetsPost($libraryVariableSetResource)

Create a LibraryVariableSetResource

Creates a new library variable set.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryVariableSetResource** | [**LibraryVariableSetResource**](LibraryVariableSetResource.md)| The LibraryVariableSetResource resource to create | [optional] 

### Return type

[**LibraryVariableSetResource**](LibraryVariableSetResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

