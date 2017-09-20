# \FeedsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiFeedsAllGet**](FeedsApi.md#ApiFeedsAllGet) | **Get** /api/feeds/all | Get a list of FeedResources
[**ApiFeedsGet**](FeedsApi.md#ApiFeedsGet) | **Get** /api/feeds | Get a list of FeedResources
[**ApiFeedsIdDelete**](FeedsApi.md#ApiFeedsIdDelete) | **Delete** /api/feeds/{id} | Delete a FeedResource by ID
[**ApiFeedsIdGet**](FeedsApi.md#ApiFeedsIdGet) | **Get** /api/feeds/{id} | Get a FeedResource by ID
[**ApiFeedsIdPut**](FeedsApi.md#ApiFeedsIdPut) | **Put** /api/feeds/{id} | Modify a FeedResource by ID
[**ApiFeedsPost**](FeedsApi.md#ApiFeedsPost) | **Post** /api/feeds | Create a FeedResource


# **ApiFeedsAllGet**
> []FeedResource ApiFeedsAllGet()

Get a list of FeedResources

Lists all the feeds in the Octopus Deploy installation.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]FeedResource**](FeedResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiFeedsGet**
> ResourceCollectionFeedResource ApiFeedsGet($skip, $take)

Get a list of FeedResources

Lists all of the external NuGet feeds used by the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionFeedResource**](ResourceCollection[FeedResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiFeedsIdDelete**
> TaskResource ApiFeedsIdDelete($id)

Delete a FeedResource by ID

Deletes an existing feed.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the FeedResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiFeedsIdGet**
> FeedResource ApiFeedsIdGet($id)

Get a FeedResource by ID

Gets a single feed by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the FeedResource to load | 

### Return type

[**FeedResource**](FeedResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiFeedsIdPut**
> FeedResource ApiFeedsIdPut($id, $feedResource)

Modify a FeedResource by ID

Modifies an existing feed.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the FeedResource to modify | 
 **feedResource** | [**FeedResource**](FeedResource.md)| The FeedResource resource to create | [optional] 

### Return type

[**FeedResource**](FeedResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiFeedsPost**
> FeedResource ApiFeedsPost($feedResource)

Create a FeedResource

Creates a feed.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedResource** | [**FeedResource**](FeedResource.md)| The FeedResource resource to create | [optional] 

### Return type

[**FeedResource**](FeedResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

