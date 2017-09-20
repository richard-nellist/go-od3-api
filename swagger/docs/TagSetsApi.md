# \TagSetsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTagsetsAllGet**](TagSetsApi.md#ApiTagsetsAllGet) | **Get** /api/tagsets/all | Get a list of TagSetResources
[**ApiTagsetsGet**](TagSetsApi.md#ApiTagsetsGet) | **Get** /api/tagsets | Get a list of TagSetResources
[**ApiTagsetsIdDelete**](TagSetsApi.md#ApiTagsetsIdDelete) | **Delete** /api/tagsets/{id} | Delete a TagSetResource by ID
[**ApiTagsetsIdGet**](TagSetsApi.md#ApiTagsetsIdGet) | **Get** /api/tagsets/{id} | Get a TagSetResource by ID
[**ApiTagsetsIdPut**](TagSetsApi.md#ApiTagsetsIdPut) | **Put** /api/tagsets/{id} | Modify a TagSetResource by ID
[**ApiTagsetsPost**](TagSetsApi.md#ApiTagsetsPost) | **Post** /api/tagsets | Create a TagSetResource
[**ApiTagsetsSortorderPut**](TagSetsApi.md#ApiTagsetsSortorderPut) | **Put** /api/tagsets/sortorder | 


# **ApiTagsetsAllGet**
> []TagSetResource ApiTagsetsAllGet()

Get a list of TagSetResources

Lists the details of all of the tag sets in the current Octopus installation. The results will be sorted by the `SortOrder` field on each tag set.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]TagSetResource**](TagSetResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTagsetsGet**
> ResourceCollectionTagSetResource ApiTagsetsGet($skip, $take)

Get a list of TagSetResources

Lists all of the tag sets in the current Octopus installation. The results will be sorted alphabetically by the `SortOrder` field on each tag set.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionTagSetResource**](ResourceCollection[TagSetResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTagsetsIdDelete**
> TaskResource ApiTagsetsIdDelete($id)

Delete a TagSetResource by ID

Deletes an existing tag set.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the TagSetResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTagsetsIdGet**
> TagSetResource ApiTagsetsIdGet($id)

Get a TagSetResource by ID

Gets a tag set by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the TagSetResource to load | 

### Return type

[**TagSetResource**](TagSetResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTagsetsIdPut**
> TagSetResource ApiTagsetsIdPut($id, $tagSetResource)

Modify a TagSetResource by ID

Modifies an existing tag set.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the TagSetResource to modify | 
 **tagSetResource** | [**TagSetResource**](TagSetResource.md)| The TagSetResource resource to create | [optional] 

### Return type

[**TagSetResource**](TagSetResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTagsetsPost**
> TagSetResource ApiTagsetsPost($tagSetResource)

Create a TagSetResource

Creates a new tag set.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagSetResource** | [**TagSetResource**](TagSetResource.md)| The TagSetResource resource to create | [optional] 

### Return type

[**TagSetResource**](TagSetResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTagsetsSortorderPut**
> ApiTagsetsSortorderPut()



Takes an array of tag set IDs as the request body, uses the order of items in the array to sort the tag sets on the server. The ID of every tag set must be specified.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

