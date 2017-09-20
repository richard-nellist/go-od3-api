# \ReleasesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectsIdReleasesGet**](ReleasesApi.md#ApiProjectsIdReleasesGet) | **Get** /api/projects/{id}/releases | Get a list of ReleaseResources
[**ApiProjectsIdReleasesVersionGet**](ReleasesApi.md#ApiProjectsIdReleasesVersionGet) | **Get** /api/projects/{id}/releases/{version} | 
[**ApiReleasesGet**](ReleasesApi.md#ApiReleasesGet) | **Get** /api/releases | Get a list of ReleaseResources
[**ApiReleasesIdDefectsGet**](ReleasesApi.md#ApiReleasesIdDefectsGet) | **Get** /api/releases/{id}/defects | 
[**ApiReleasesIdDefectsPost**](ReleasesApi.md#ApiReleasesIdDefectsPost) | **Post** /api/releases/{id}/defects | 
[**ApiReleasesIdDefectsResolvePost**](ReleasesApi.md#ApiReleasesIdDefectsResolvePost) | **Post** /api/releases/{id}/defects/resolve | 
[**ApiReleasesIdDelete**](ReleasesApi.md#ApiReleasesIdDelete) | **Delete** /api/releases/{id} | Delete a ReleaseResource by ID
[**ApiReleasesIdDeploymentsGet**](ReleasesApi.md#ApiReleasesIdDeploymentsGet) | **Get** /api/releases/{id}/deployments | Get a list of DeploymentResources
[**ApiReleasesIdDeploymentsPreviewEnvironmentGet**](ReleasesApi.md#ApiReleasesIdDeploymentsPreviewEnvironmentGet) | **Get** /api/releases/{id}/deployments/preview/{environment} | 
[**ApiReleasesIdDeploymentsPreviewEnvironmentTenantGet**](ReleasesApi.md#ApiReleasesIdDeploymentsPreviewEnvironmentTenantGet) | **Get** /api/releases/{id}/deployments/preview/{environment}/{tenant} | 
[**ApiReleasesIdDeploymentsTemplateGet**](ReleasesApi.md#ApiReleasesIdDeploymentsTemplateGet) | **Get** /api/releases/{id}/deployments/template | 
[**ApiReleasesIdGet**](ReleasesApi.md#ApiReleasesIdGet) | **Get** /api/releases/{id} | Get a ReleaseResource by ID
[**ApiReleasesIdProgressionGet**](ReleasesApi.md#ApiReleasesIdProgressionGet) | **Get** /api/releases/{id}/progression | 
[**ApiReleasesIdPut**](ReleasesApi.md#ApiReleasesIdPut) | **Put** /api/releases/{id} | Modify a ReleaseResource by ID
[**ApiReleasesIdSnapshotVariablesPost**](ReleasesApi.md#ApiReleasesIdSnapshotVariablesPost) | **Post** /api/releases/{id}/snapshot-variables | 
[**ApiReleasesPost**](ReleasesApi.md#ApiReleasesPost) | **Post** /api/releases | Create a ReleaseResource


# **ApiProjectsIdReleasesGet**
> ResourceCollectionProjectResource ApiProjectsIdReleasesGet($skip, $take)

Get a list of ReleaseResources

Lists all of the releases that belong to the given project. Releases will be ordered from most recent to least recent.


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

# **ApiProjectsIdReleasesVersionGet**
> ReleaseResource ApiProjectsIdReleasesVersionGet($version, $id)



Gets a single release by project ID and version number.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string**| ERROR - NOT DEFINED | 
 **id** | **string**| ID of the resource | 

### Return type

[**ReleaseResource**](ReleaseResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReleasesGet**
> ResourceCollectionReleaseResource ApiReleasesGet($skip, $take)

Get a list of ReleaseResources

Lists all of the releases in the current Octopus installation, from all projects. The results will be sorted from most recent to least recent release.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionReleaseResource**](ResourceCollection[ReleaseResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReleasesIdDefectsGet**
> ApiReleasesIdDefectsGet()



Gets all defects for a release.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiReleasesIdDefectsPost**
> DefectResource ApiReleasesIdDefectsPost($id)



Record defect in a release.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**DefectResource**](DefectResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReleasesIdDefectsResolvePost**
> DefectResource ApiReleasesIdDefectsResolvePost($id)



Update or resolve defect in a release.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**DefectResource**](DefectResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReleasesIdDelete**
> TaskResource ApiReleasesIdDelete($id)

Delete a ReleaseResource by ID

Deletes an existing release, along with all of the deployments, tasks and other associated resources belonging to the release.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ReleaseResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReleasesIdDeploymentsGet**
> ResourceCollectionReleaseResource ApiReleasesIdDeploymentsGet($skip, $take)

Get a list of DeploymentResources

Lists all of the deployments that belong to the given release. Deployments will be ordered from most recent to least recent.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionReleaseResource**](ResourceCollection[ReleaseResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReleasesIdDeploymentsPreviewEnvironmentGet**
> DeploymentPreviewResource ApiReleasesIdDeploymentsPreviewEnvironmentGet($environment, $tenant, $id)



Gets a document that describes what steps will/won't be run during a deployment to a given environment.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string**| ID of the environment | 
 **tenant** | **string**| ID of the tenant | 
 **id** | **string**| ID of the resource | 

### Return type

[**DeploymentPreviewResource**](DeploymentPreviewResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReleasesIdDeploymentsPreviewEnvironmentTenantGet**
> DeploymentPreviewResource ApiReleasesIdDeploymentsPreviewEnvironmentTenantGet($environment, $tenant, $id)



Gets a document that describes what steps will/won't be run during a deployment to a given environment for the given tenant.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string**| ID of the environment | 
 **tenant** | **string**| ID of the tenant | 
 **id** | **string**| ID of the resource | 

### Return type

[**DeploymentPreviewResource**](DeploymentPreviewResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReleasesIdDeploymentsTemplateGet**
> DeploymentTemplateResource ApiReleasesIdDeploymentsTemplateGet($id)



Gets all of the information necessary for creating or editing a deployment for this release.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**DeploymentTemplateResource**](DeploymentTemplateResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReleasesIdGet**
> ReleaseResource ApiReleasesIdGet($id)

Get a ReleaseResource by ID

Gets a release by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ReleaseResource to load | 

### Return type

[**ReleaseResource**](ReleaseResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReleasesIdProgressionGet**
> LifecycleProgressionResource ApiReleasesIdProgressionGet($id)



Gets all of the information necessary for creating or editing a deployment for this release.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**LifecycleProgressionResource**](LifecycleProgressionResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReleasesIdPut**
> ReleaseResource ApiReleasesIdPut($id, $releaseResource)

Modify a ReleaseResource by ID

Updates an existing release.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ReleaseResource to modify | 
 **releaseResource** | [**ReleaseResource**](ReleaseResource.md)| The ReleaseResource resource to create | [optional] 

### Return type

[**ReleaseResource**](ReleaseResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReleasesIdSnapshotVariablesPost**
> ReleaseResource ApiReleasesIdSnapshotVariablesPost($id)



Refresh the variable snapshots associated with the release. The project's deployment process must not have changed since the release was created.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**ReleaseResource**](ReleaseResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReleasesPost**
> ReleaseResource ApiReleasesPost($releaseResource)

Create a ReleaseResource

Creates a new release.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseResource** | [**ReleaseResource**](ReleaseResource.md)| The ReleaseResource resource to create | [optional] 

### Return type

[**ReleaseResource**](ReleaseResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

