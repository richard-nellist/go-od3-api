# \MachinesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiMachinesAllGet**](MachinesApi.md#ApiMachinesAllGet) | **Get** /api/machines/all | Get a list of MachineResources
[**ApiMachinesDiscoverGet**](MachinesApi.md#ApiMachinesDiscoverGet) | **Get** /api/machines/discover | 
[**ApiMachinesGet**](MachinesApi.md#ApiMachinesGet) | **Get** /api/machines | Get a list of MachineResources
[**ApiMachinesIdConnectionGet**](MachinesApi.md#ApiMachinesIdConnectionGet) | **Get** /api/machines/{id}/connection | 
[**ApiMachinesIdDelete**](MachinesApi.md#ApiMachinesIdDelete) | **Delete** /api/machines/{id} | Delete a MachineResource by ID
[**ApiMachinesIdGet**](MachinesApi.md#ApiMachinesIdGet) | **Get** /api/machines/{id} | Get a MachineResource by ID
[**ApiMachinesIdPut**](MachinesApi.md#ApiMachinesIdPut) | **Put** /api/machines/{id} | Modify a MachineResource by ID
[**ApiMachinesIdTasksGet**](MachinesApi.md#ApiMachinesIdTasksGet) | **Get** /api/machines/{id}/tasks | 
[**ApiMachinesPost**](MachinesApi.md#ApiMachinesPost) | **Post** /api/machines | Create a MachineResource


# **ApiMachinesAllGet**
> []MachineResource ApiMachinesAllGet()

Get a list of MachineResources

Lists all of the machines in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]MachineResource**](MachineResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMachinesDiscoverGet**
> MachineResource ApiMachinesDiscoverGet()



Interrogate a machine for communication details so that it may be added to the installation.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

[**MachineResource**](MachineResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMachinesGet**
> ResourceCollectionMachineResource ApiMachinesGet($skip, $take)

Get a list of MachineResources

Lists all of the registered machines in the current Octopus installation, from all environments. The results will be sorted alphabetically by name.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionMachineResource**](ResourceCollection[MachineResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMachinesIdConnectionGet**
> MachineConnectionStatus ApiMachinesIdConnectionGet($id)



Get the status of the network connection between the Octopus server and a machine.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**MachineConnectionStatus**](MachineConnectionStatus.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMachinesIdDelete**
> TaskResource ApiMachinesIdDelete($id)

Delete a MachineResource by ID

Deletes an existing machine.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the MachineResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMachinesIdGet**
> MachineResource ApiMachinesIdGet($id)

Get a MachineResource by ID

Gets a single machine by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the MachineResource to load | 

### Return type

[**MachineResource**](MachineResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMachinesIdPut**
> MachineResource ApiMachinesIdPut($id, $machineResource)

Modify a MachineResource by ID

Modifies an existing machine.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the MachineResource to modify | 
 **machineResource** | [**MachineResource**](MachineResource.md)| The MachineResource resource to create | [optional] 

### Return type

[**MachineResource**](MachineResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMachinesIdTasksGet**
> ApiMachinesIdTasksGet()



Get the history of related tasks for a machine.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiMachinesPost**
> MachineResource ApiMachinesPost($machineResource)

Create a MachineResource

Creates a new machine.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **machineResource** | [**MachineResource**](MachineResource.md)| The MachineResource resource to create | [optional] 

### Return type

[**MachineResource**](MachineResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

