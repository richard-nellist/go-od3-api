# \EnvironmentsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEnvironmentsAllGet**](EnvironmentsApi.md#ApiEnvironmentsAllGet) | **Get** /api/environments/all | Get a list of EnvironmentResources
[**ApiEnvironmentsGet**](EnvironmentsApi.md#ApiEnvironmentsGet) | **Get** /api/environments | Get a list of EnvironmentResources
[**ApiEnvironmentsIdDelete**](EnvironmentsApi.md#ApiEnvironmentsIdDelete) | **Delete** /api/environments/{id} | Delete a EnvironmentResource by ID
[**ApiEnvironmentsIdGet**](EnvironmentsApi.md#ApiEnvironmentsIdGet) | **Get** /api/environments/{id} | Get a EnvironmentResource by ID
[**ApiEnvironmentsIdMachinesGet**](EnvironmentsApi.md#ApiEnvironmentsIdMachinesGet) | **Get** /api/environments/{id}/machines | Get a list of MachineResources
[**ApiEnvironmentsIdPut**](EnvironmentsApi.md#ApiEnvironmentsIdPut) | **Put** /api/environments/{id} | Modify a EnvironmentResource by ID
[**ApiEnvironmentsIdSinglyScopedVariableDetailsGet**](EnvironmentsApi.md#ApiEnvironmentsIdSinglyScopedVariableDetailsGet) | **Get** /api/environments/{id}/singlyScopedVariableDetails | 
[**ApiEnvironmentsPost**](EnvironmentsApi.md#ApiEnvironmentsPost) | **Post** /api/environments | Create a EnvironmentResource
[**ApiEnvironmentsSortorderPut**](EnvironmentsApi.md#ApiEnvironmentsSortorderPut) | **Put** /api/environments/sortorder | 


# **ApiEnvironmentsAllGet**
> []EnvironmentResource ApiEnvironmentsAllGet()

Get a list of EnvironmentResources

Lists the name and ID of all of the environments in the current Octopus installation. The results will be sorted by the `SortOrder` field on each environment.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]EnvironmentResource**](EnvironmentResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEnvironmentsGet**
> ResourceCollectionEnvironmentResource ApiEnvironmentsGet($skip, $take)

Get a list of EnvironmentResources

Lists all of the environments in the current Octopus installation. The results will be sorted by the `SortOrder` field on each environment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionEnvironmentResource**](ResourceCollection[EnvironmentResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEnvironmentsIdDelete**
> TaskResource ApiEnvironmentsIdDelete($id)

Delete a EnvironmentResource by ID

Deletes an existing environment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the EnvironmentResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEnvironmentsIdGet**
> EnvironmentResource ApiEnvironmentsIdGet($id)

Get a EnvironmentResource by ID

Gets a single environment by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the EnvironmentResource to load | 

### Return type

[**EnvironmentResource**](EnvironmentResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEnvironmentsIdMachinesGet**
> ResourceCollectionEnvironmentResource ApiEnvironmentsIdMachinesGet($skip, $take)

Get a list of MachineResources

Lists all of the machines that belong to the given environment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionEnvironmentResource**](ResourceCollection[EnvironmentResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEnvironmentsIdPut**
> EnvironmentResource ApiEnvironmentsIdPut($id, $environmentResource)

Modify a EnvironmentResource by ID

Modifies an existing environment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the EnvironmentResource to modify | 
 **environmentResource** | [**EnvironmentResource**](EnvironmentResource.md)| The EnvironmentResource resource to create | [optional] 

### Return type

[**EnvironmentResource**](EnvironmentResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEnvironmentsIdSinglyScopedVariableDetailsGet**
> VariablesScopedToEnvironmentResponse ApiEnvironmentsIdSinglyScopedVariableDetailsGet($id)



Lists all the variable set names (projects and library variable sets) that have variables that are scoped to only the given environment  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**VariablesScopedToEnvironmentResponse**](VariablesScopedToEnvironmentResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEnvironmentsPost**
> EnvironmentResource ApiEnvironmentsPost($environmentResource)

Create a EnvironmentResource

Creates a new environment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentResource** | [**EnvironmentResource**](EnvironmentResource.md)| The EnvironmentResource resource to create | [optional] 

### Return type

[**EnvironmentResource**](EnvironmentResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEnvironmentsSortorderPut**
> ApiEnvironmentsSortorderPut()



Takes an array of environment IDs as the request body, uses the order of items in the array to sort the environments on the server. The ID of every environment must be specified.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

