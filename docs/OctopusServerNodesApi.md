# \OctopusServerNodesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiOctopusservernodesAllGet**](OctopusServerNodesApi.md#ApiOctopusservernodesAllGet) | **Get** /api/octopusservernodes/all | Get a list of OctopusServerNodeResources
[**ApiOctopusservernodesGet**](OctopusServerNodesApi.md#ApiOctopusservernodesGet) | **Get** /api/octopusservernodes | Get a list of OctopusServerNodeResources
[**ApiOctopusservernodesIdDelete**](OctopusServerNodesApi.md#ApiOctopusservernodesIdDelete) | **Delete** /api/octopusservernodes/{id} | Delete a OctopusServerNodeResource by ID
[**ApiOctopusservernodesIdGet**](OctopusServerNodesApi.md#ApiOctopusservernodesIdGet) | **Get** /api/octopusservernodes/{id} | Get a OctopusServerNodeResource by ID
[**ApiOctopusservernodesIdPut**](OctopusServerNodesApi.md#ApiOctopusservernodesIdPut) | **Put** /api/octopusservernodes/{id} | Modify a OctopusServerNodeResource by ID
[**ApiOctopusservernodesPingGet**](OctopusServerNodesApi.md#ApiOctopusservernodesPingGet) | **Get** /api/octopusservernodes/ping | 


# **ApiOctopusservernodesAllGet**
> []OctopusServerNodeResource ApiOctopusservernodesAllGet()

Get a list of OctopusServerNodeResources

Lists the name and ID of all Octopus Server nodes.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]OctopusServerNodeResource**](OctopusServerNodeResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOctopusservernodesGet**
> ResourceCollectionOctopusServerNodeResource ApiOctopusservernodesGet($skip, $take)

Get a list of OctopusServerNodeResources

List all of the Octopus Server nodes participating in the current Octopus Server cluster.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionOctopusServerNodeResource**](ResourceCollection[OctopusServerNodeResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOctopusservernodesIdDelete**
> TaskResource ApiOctopusservernodesIdDelete($id)

Delete a OctopusServerNodeResource by ID

Deletes an Octopus Server node.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the OctopusServerNodeResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOctopusservernodesIdGet**
> OctopusServerNodeResource ApiOctopusservernodesIdGet($id)

Get a OctopusServerNodeResource by ID

Gets an Octopus Server node by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the OctopusServerNodeResource to load | 

### Return type

[**OctopusServerNodeResource**](OctopusServerNodeResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOctopusservernodesIdPut**
> OctopusServerNodeResource ApiOctopusservernodesIdPut($id, $octopusServerNodeResource)

Modify a OctopusServerNodeResource by ID

Modifies an Octopus Server node.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the OctopusServerNodeResource to modify | 
 **octopusServerNodeResource** | [**OctopusServerNodeResource**](OctopusServerNodeResource.md)| The OctopusServerNodeResource resource to create | [optional] 

### Return type

[**OctopusServerNodeResource**](OctopusServerNodeResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOctopusservernodesPingGet**
> ApiOctopusservernodesPingGet()



Returns HTTP OK (200) when the Octopus Server node is not draining.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

