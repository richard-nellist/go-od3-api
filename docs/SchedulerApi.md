# \SchedulerApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSchedulerNameLogsGet**](SchedulerApi.md#ApiSchedulerNameLogsGet) | **Get** /api/scheduler/{name}/logs | 
[**ApiSchedulerStartGet**](SchedulerApi.md#ApiSchedulerStartGet) | **Get** /api/scheduler/start | 
[**ApiSchedulerStopGet**](SchedulerApi.md#ApiSchedulerStopGet) | **Get** /api/scheduler/stop | 
[**ApiSchedulerTriggerGet**](SchedulerApi.md#ApiSchedulerTriggerGet) | **Get** /api/scheduler/trigger | 


# **ApiSchedulerNameLogsGet**
> ApiSchedulerNameLogsGet($name)



Get the details of a scheduled task.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of the task | 

### Return type

void (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiSchedulerStartGet**
> ApiSchedulerStartGet()



Returns HTTP OK (200) when the Octopus Server scheduler has been started.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiSchedulerStopGet**
> ApiSchedulerStopGet()



Returns HTTP OK (200) when the Octopus Server scheduler has been stopped.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiSchedulerTriggerGet**
> ApiSchedulerTriggerGet()



Returns HTTP OK (200) when the Octopus Server scheduler has been triggered.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

