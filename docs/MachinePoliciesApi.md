# \MachinePoliciesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiMachinepoliciesAllGet**](MachinePoliciesApi.md#ApiMachinepoliciesAllGet) | **Get** /api/machinepolicies/all | Get a list of MachinePolicyResources
[**ApiMachinepoliciesGet**](MachinePoliciesApi.md#ApiMachinepoliciesGet) | **Get** /api/machinepolicies | Get a list of MachinePolicyResources
[**ApiMachinepoliciesIdDelete**](MachinePoliciesApi.md#ApiMachinepoliciesIdDelete) | **Delete** /api/machinepolicies/{id} | 
[**ApiMachinepoliciesIdGet**](MachinePoliciesApi.md#ApiMachinepoliciesIdGet) | **Get** /api/machinepolicies/{id} | Get a MachinePolicyResource by ID
[**ApiMachinepoliciesIdMachinesGet**](MachinePoliciesApi.md#ApiMachinepoliciesIdMachinesGet) | **Get** /api/machinepolicies/{id}/machines | Get a list of MachineResources
[**ApiMachinepoliciesIdPut**](MachinePoliciesApi.md#ApiMachinepoliciesIdPut) | **Put** /api/machinepolicies/{id} | Modify a MachinePolicyResource by ID
[**ApiMachinepoliciesPost**](MachinePoliciesApi.md#ApiMachinepoliciesPost) | **Post** /api/machinepolicies | Create a MachinePolicyResource


# **ApiMachinepoliciesAllGet**
> []MachinePolicyResource ApiMachinepoliciesAllGet()

Get a list of MachinePolicyResources

Lists all the machine policies in the Octopus Deploy installation.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]MachinePolicyResource**](MachinePolicyResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMachinepoliciesGet**
> ResourceCollectionMachinePolicyResource ApiMachinepoliciesGet($skip, $take)

Get a list of MachinePolicyResources

Lists all of the machine policies in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionMachinePolicyResource**](ResourceCollection[MachinePolicyResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMachinepoliciesIdDelete**
> ApiMachinepoliciesIdDelete($id)



Deletes an existing machine policy.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiMachinepoliciesIdGet**
> MachinePolicyResource ApiMachinepoliciesIdGet($id)

Get a MachinePolicyResource by ID

Gets a single machine policy by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the MachinePolicyResource to load | 

### Return type

[**MachinePolicyResource**](MachinePolicyResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMachinepoliciesIdMachinesGet**
> ResourceCollectionMachinePolicyResource ApiMachinepoliciesIdMachinesGet($skip, $take)

Get a list of MachineResources

Lists all of the machines that belong to the given machine policy.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionMachinePolicyResource**](ResourceCollection[MachinePolicyResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMachinepoliciesIdPut**
> MachinePolicyResource ApiMachinepoliciesIdPut($id, $machinePolicyResource)

Modify a MachinePolicyResource by ID

Modifies an existing machine policy.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the MachinePolicyResource to modify | 
 **machinePolicyResource** | [**MachinePolicyResource**](MachinePolicyResource.md)| The MachinePolicyResource resource to create | [optional] 

### Return type

[**MachinePolicyResource**](MachinePolicyResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMachinepoliciesPost**
> MachinePolicyResource ApiMachinepoliciesPost($machinePolicyResource)

Create a MachinePolicyResource

Creates a new machine policy.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **machinePolicyResource** | [**MachinePolicyResource**](MachinePolicyResource.md)| The MachinePolicyResource resource to create | [optional] 

### Return type

[**MachinePolicyResource**](MachinePolicyResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

