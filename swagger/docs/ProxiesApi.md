# \ProxiesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProxiesAllGet**](ProxiesApi.md#ApiProxiesAllGet) | **Get** /api/proxies/all | Get a list of ProxyResources
[**ApiProxiesGet**](ProxiesApi.md#ApiProxiesGet) | **Get** /api/proxies | Get a list of ProxyResources
[**ApiProxiesIdDelete**](ProxiesApi.md#ApiProxiesIdDelete) | **Delete** /api/proxies/{id} | Delete a ProxyResource by ID
[**ApiProxiesIdGet**](ProxiesApi.md#ApiProxiesIdGet) | **Get** /api/proxies/{id} | Get a ProxyResource by ID
[**ApiProxiesIdPut**](ProxiesApi.md#ApiProxiesIdPut) | **Put** /api/proxies/{id} | Modify a ProxyResource by ID
[**ApiProxiesPost**](ProxiesApi.md#ApiProxiesPost) | **Post** /api/proxies | Create a ProxyResource


# **ApiProxiesAllGet**
> []ProxyResource ApiProxiesAllGet()

Get a list of ProxyResources

Lists the name and ID of all of the proxies in the current Octopus installation. The results will be sorted by name.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProxyResource**](ProxyResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProxiesGet**
> ResourceCollectionProxyResource ApiProxiesGet($skip, $take)

Get a list of ProxyResources

Lists all of the proxies in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionProxyResource**](ResourceCollection[ProxyResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProxiesIdDelete**
> TaskResource ApiProxiesIdDelete($id)

Delete a ProxyResource by ID

Deletes an existing proxy.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ProxyResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProxiesIdGet**
> ProxyResource ApiProxiesIdGet($id)

Get a ProxyResource by ID

Gets a proxy by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ProxyResource to load | 

### Return type

[**ProxyResource**](ProxyResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProxiesIdPut**
> ProxyResource ApiProxiesIdPut($id, $proxyResource)

Modify a ProxyResource by ID

Modifies an existing proxy.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ProxyResource to modify | 
 **proxyResource** | [**ProxyResource**](ProxyResource.md)| The ProxyResource resource to create | [optional] 

### Return type

[**ProxyResource**](ProxyResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProxiesPost**
> ProxyResource ApiProxiesPost($proxyResource)

Create a ProxyResource

Creates a proxy.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proxyResource** | [**ProxyResource**](ProxyResource.md)| The ProxyResource resource to create | [optional] 

### Return type

[**ProxyResource**](ProxyResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

