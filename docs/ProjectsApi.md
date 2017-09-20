# \ProjectsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectsAllGet**](ProjectsApi.md#ApiProjectsAllGet) | **Get** /api/projects/all | Get a list of ProjectResources
[**ApiProjectsGet**](ProjectsApi.md#ApiProjectsGet) | **Get** /api/projects | Get a list of ProjectResources
[**ApiProjectsIdDelete**](ProjectsApi.md#ApiProjectsIdDelete) | **Delete** /api/projects/{id} | Delete a ProjectResource by ID
[**ApiProjectsIdLogoGet**](ProjectsApi.md#ApiProjectsIdLogoGet) | **Get** /api/projects/{id}/logo | 
[**ApiProjectsIdLogoPost**](ProjectsApi.md#ApiProjectsIdLogoPost) | **Post** /api/projects/{id}/logo | 
[**ApiProjectsIdLogoPut**](ProjectsApi.md#ApiProjectsIdLogoPut) | **Put** /api/projects/{id}/logo | 
[**ApiProjectsIdOrSlugGet**](ProjectsApi.md#ApiProjectsIdOrSlugGet) | **Get** /api/projects/{idOrSlug} | Get a ProjectResource by ID or slug
[**ApiProjectsIdPut**](ProjectsApi.md#ApiProjectsIdPut) | **Put** /api/projects/{id} | Modify a ProjectResource by ID
[**ApiProjectsPost**](ProjectsApi.md#ApiProjectsPost) | **Post** /api/projects | Create a ProjectResource


# **ApiProjectsAllGet**
> []ProjectResource ApiProjectsAllGet()

Get a list of ProjectResources

Lists the name and ID of all of the projects in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProjectResource**](ProjectResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsGet**
> ResourceCollectionProjectResource ApiProjectsGet($skip, $take)

Get a list of ProjectResources

Lists all of the projects in the current Octopus installation, from all project groups. The results will be sorted alphabetically by name.


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

# **ApiProjectsIdDelete**
> TaskResource ApiProjectsIdDelete($id)

Delete a ProjectResource by ID

Deletes an existing project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ProjectResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsIdLogoGet**
> *os.File ApiProjectsIdLogoGet($id)



Gets the logo associated with the project.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsIdLogoPost**
> ApiProjectsIdLogoPost($id)



Updates the logo associated with the project.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

void (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsIdLogoPut**
> ApiProjectsIdLogoPut($id)



Updates the logo associated with the project.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

void (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsIdOrSlugGet**
> ProjectResource ApiProjectsIdOrSlugGet($idOrSlug)

Get a ProjectResource by ID or slug

Gets a single project by ID or Slug.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idOrSlug** | **string**| ID or slug of the ProjectResource to load | 

### Return type

[**ProjectResource**](ProjectResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsIdPut**
> ProjectResource ApiProjectsIdPut($id, $projectResource)

Modify a ProjectResource by ID

Modifies an existing project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ProjectResource to modify | 
 **projectResource** | [**ProjectResource**](ProjectResource.md)| The ProjectResource resource to create | [optional] 

### Return type

[**ProjectResource**](ProjectResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsPost**
> ProjectResource ApiProjectsPost($projectResource)

Create a ProjectResource

Creates a new project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectResource** | [**ProjectResource**](ProjectResource.md)| The ProjectResource resource to create | [optional] 

### Return type

[**ProjectResource**](ProjectResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

