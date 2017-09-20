# \ArtifactsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiArtifactsGet**](ArtifactsApi.md#ApiArtifactsGet) | **Get** /api/artifacts | 
[**ApiArtifactsIdContentGet**](ArtifactsApi.md#ApiArtifactsIdContentGet) | **Get** /api/artifacts/{id}/content | 
[**ApiArtifactsIdContentPut**](ArtifactsApi.md#ApiArtifactsIdContentPut) | **Put** /api/artifacts/{id}/content | 
[**ApiArtifactsIdDelete**](ArtifactsApi.md#ApiArtifactsIdDelete) | **Delete** /api/artifacts/{id} | Delete a ArtifactResource by ID
[**ApiArtifactsIdGet**](ArtifactsApi.md#ApiArtifactsIdGet) | **Get** /api/artifacts/{id} | Get a ArtifactResource by ID
[**ApiArtifactsIdPut**](ArtifactsApi.md#ApiArtifactsIdPut) | **Put** /api/artifacts/{id} | Modify a ArtifactResource by ID
[**ApiArtifactsPost**](ArtifactsApi.md#ApiArtifactsPost) | **Post** /api/artifacts | Create a ArtifactResource


# **ApiArtifactsGet**
> ApiArtifactsGet()



Lists all of the artifacts in the current Octopus installation, from all releases. The results will be sorted by date from most recently to least recently created.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiArtifactsIdContentGet**
> ApiArtifactsIdContentGet()



Gets the content associated with an artifact.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiArtifactsIdContentPut**
> ApiArtifactsIdContentPut()



Sets the content associated with an artifact.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiArtifactsIdDelete**
> TaskResource ApiArtifactsIdDelete($id)

Delete a ArtifactResource by ID

Deletes an existing artifact.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ArtifactResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiArtifactsIdGet**
> ArtifactResource ApiArtifactsIdGet($id)

Get a ArtifactResource by ID

Gets a single artifact by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ArtifactResource to load | 

### Return type

[**ArtifactResource**](ArtifactResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiArtifactsIdPut**
> ArtifactResource ApiArtifactsIdPut($id, $artifactResource)

Modify a ArtifactResource by ID

Modifies an existing artifact.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ArtifactResource to modify | 
 **artifactResource** | [**ArtifactResource**](ArtifactResource.md)| The ArtifactResource resource to create | [optional] 

### Return type

[**ArtifactResource**](ArtifactResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiArtifactsPost**
> ArtifactResource ApiArtifactsPost($artifactResource)

Create a ArtifactResource

Creates a new artifacts.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifactResource** | [**ArtifactResource**](ArtifactResource.md)| The ArtifactResource resource to create | [optional] 

### Return type

[**ArtifactResource**](ArtifactResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

