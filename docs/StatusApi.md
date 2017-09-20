# \StatusApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServerstatusExtensionsGet**](StatusApi.md#ApiServerstatusExtensionsGet) | **Get** /api/serverstatus/extensions | 
[**ApiServerstatusGcCollectPost**](StatusApi.md#ApiServerstatusGcCollectPost) | **Post** /api/serverstatus/gc-collect | 
[**ApiServerstatusLogsGet**](StatusApi.md#ApiServerstatusLogsGet) | **Get** /api/serverstatus/logs | 
[**ApiServerstatusNugetGet**](StatusApi.md#ApiServerstatusNugetGet) | **Get** /api/serverstatus/nuget | 
[**ApiServerstatusSystemInfoGet**](StatusApi.md#ApiServerstatusSystemInfoGet) | **Get** /api/serverstatus/system-info | 
[**ApiServerstatusSystemReportGet**](StatusApi.md#ApiServerstatusSystemReportGet) | **Get** /api/serverstatus/system-report | 
[**ApiServerstatusTimezonesGet**](StatusApi.md#ApiServerstatusTimezonesGet) | **Get** /api/serverstatus/timezones | 


# **ApiServerstatusExtensionsGet**
> []ExtensionsInfoResource ApiServerstatusExtensionsGet()



Provides statistics for the loaded server extensions.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]ExtensionsInfoResource**](ExtensionsInfoResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServerstatusGcCollectPost**
> ApiServerstatusGcCollectPost()



Forces a GC collect.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiServerstatusLogsGet**
> ApiServerstatusLogsGet()



Retrieves the most recent high-priority log messages from this execution of the Octopus Server process.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiServerstatusNugetGet**
> BuiltInFeedStatsResource ApiServerstatusNugetGet()



Provides statistics for the built-in package repository.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

[**BuiltInFeedStatsResource**](BuiltInFeedStatsResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServerstatusSystemInfoGet**
> ApiServerstatusSystemInfoGet()



Provides information about the Octopus Server process and the machine on which it is running.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiServerstatusSystemReportGet**
> ApiServerstatusSystemReportGet()



Creates a .zip archive containing an aggregate of the other system information APIs.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiServerstatusTimezonesGet**
> ApiServerstatusTimezonesGet()



Lists timezones supported by the server.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

