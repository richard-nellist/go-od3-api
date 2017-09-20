# \ChannelsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiChannelsGet**](ChannelsApi.md#ApiChannelsGet) | **Get** /api/channels | Get a list of ChannelResources
[**ApiChannelsIdDelete**](ChannelsApi.md#ApiChannelsIdDelete) | **Delete** /api/channels/{id} | Delete a ChannelResource by ID
[**ApiChannelsIdGet**](ChannelsApi.md#ApiChannelsIdGet) | **Get** /api/channels/{id} | Get a ChannelResource by ID
[**ApiChannelsIdPut**](ChannelsApi.md#ApiChannelsIdPut) | **Put** /api/channels/{id} | Modify a ChannelResource by ID
[**ApiChannelsIdReleasesGet**](ChannelsApi.md#ApiChannelsIdReleasesGet) | **Get** /api/channels/{id}/releases | Get a list of ReleaseResources
[**ApiChannelsPost**](ChannelsApi.md#ApiChannelsPost) | **Post** /api/channels | Create a ChannelResource
[**ApiChannelsRuleTestGet**](ChannelsApi.md#ApiChannelsRuleTestGet) | **Get** /api/channels/rule-test | 
[**ApiChannelsRuleTestPost**](ChannelsApi.md#ApiChannelsRuleTestPost) | **Post** /api/channels/rule-test | 
[**ApiProjectsIdChannelsGet**](ChannelsApi.md#ApiProjectsIdChannelsGet) | **Get** /api/projects/{id}/channels | Get a list of ChannelResources


# **ApiChannelsGet**
> ResourceCollectionChannelResource ApiChannelsGet($skip, $take)

Get a list of ChannelResources

Lists all of the channels in the current Octopus installation, from all projects, sorted by name.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionChannelResource**](ResourceCollection[ChannelResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiChannelsIdDelete**
> TaskResource ApiChannelsIdDelete($id)

Delete a ChannelResource by ID

Deletes an existing channel.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ChannelResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiChannelsIdGet**
> ChannelResource ApiChannelsIdGet($id)

Get a ChannelResource by ID

Get a channel


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ChannelResource to load | 

### Return type

[**ChannelResource**](ChannelResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiChannelsIdPut**
> ChannelResource ApiChannelsIdPut($id, $channelResource)

Modify a ChannelResource by ID

Updates an existing channel


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ChannelResource to modify | 
 **channelResource** | [**ChannelResource**](ChannelResource.md)| The ChannelResource resource to create | [optional] 

### Return type

[**ChannelResource**](ChannelResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiChannelsIdReleasesGet**
> ResourceCollectionChannelResource ApiChannelsIdReleasesGet($skip, $take)

Get a list of ReleaseResources

Lists all of the releases that belong to the given channel. Releases will be ordered from most recent to least recent.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionChannelResource**](ResourceCollection[ChannelResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiChannelsPost**
> ChannelResource ApiChannelsPost($channelResource)

Create a ChannelResource

Creates a new channel


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelResource** | [**ChannelResource**](ChannelResource.md)| The ChannelResource resource to create | [optional] 

### Return type

[**ChannelResource**](ChannelResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiChannelsRuleTestGet**
> ApiChannelsRuleTestGet()



Perform channel version rule test against provided package version  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiChannelsRuleTestPost**
> ApiChannelsRuleTestPost()



Perform channel version rule test against provided package version  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiProjectsIdChannelsGet**
> ResourceCollectionProjectResource ApiProjectsIdChannelsGet($skip, $take)

Get a list of ChannelResources

Lists all the channels for the given project


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

