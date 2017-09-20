# \TasksApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTasksGet**](TasksApi.md#ApiTasksGet) | **Get** /api/tasks | 
[**ApiTasksIdCancelPost**](TasksApi.md#ApiTasksIdCancelPost) | **Post** /api/tasks/{id}/cancel | 
[**ApiTasksIdDetailsGet**](TasksApi.md#ApiTasksIdDetailsGet) | **Get** /api/tasks/{id}/details | 
[**ApiTasksIdGet**](TasksApi.md#ApiTasksIdGet) | **Get** /api/tasks/{id} | Get a TaskResource by ID
[**ApiTasksIdQueuedBehindGet**](TasksApi.md#ApiTasksIdQueuedBehindGet) | **Get** /api/tasks/{id}/queued-behind | 
[**ApiTasksIdRawGet**](TasksApi.md#ApiTasksIdRawGet) | **Get** /api/tasks/{id}/raw | 
[**ApiTasksIdStatePost**](TasksApi.md#ApiTasksIdStatePost) | **Post** /api/tasks/{id}/state | 
[**ApiTasksPost**](TasksApi.md#ApiTasksPost) | **Post** /api/tasks | Create a TaskResource
[**ApiTasksRerunIdPost**](TasksApi.md#ApiTasksRerunIdPost) | **Post** /api/tasks/rerun/{id} | 


# **ApiTasksGet**
> ApiTasksGet()



Lists all of the tasks in the current Octopus installation. The results will be sorted from newest to oldest.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiTasksIdCancelPost**
> TaskResource ApiTasksIdCancelPost($id)



Marks the given task as canceled.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTasksIdDetailsGet**
> TaskDetailsServerResource ApiTasksIdDetailsGet($id)



Gets a single task by ID, including the full task log as a tree of activity elements.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**TaskDetailsServerResource**](TaskDetailsServerResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTasksIdGet**
> TaskResource ApiTasksIdGet($id)

Get a TaskResource by ID

Gets a single task by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the TaskResource to load | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTasksIdQueuedBehindGet**
> ApiTasksIdQueuedBehindGet()



Gets a list of tasks that this task is currently queued behind.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiTasksIdRawGet**
> ApiTasksIdRawGet($id)



Gets the full task log of a given resource as plain text. Useful when the log needs to be rendered to a console or sent as an email attachment.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiTasksIdStatePost**
> TaskResource ApiTasksIdStatePost($id)



Change the state of a task  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTasksPost**
> TaskResource ApiTasksPost($taskResource)

Create a TaskResource

Creates a new task.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskResource** | [**TaskResource**](TaskResource.md)| The TaskResource resource to create | [optional] 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTasksRerunIdPost**
> TaskResource ApiTasksRerunIdPost($id)



Creates a new task and executes it, using a given task as the input. Note that deployment tasks cannot be re-run.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

