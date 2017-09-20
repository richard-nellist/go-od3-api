# \InterruptionsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiInterruptionsGet**](InterruptionsApi.md#ApiInterruptionsGet) | **Get** /api/interruptions | 
[**ApiInterruptionsIdGet**](InterruptionsApi.md#ApiInterruptionsIdGet) | **Get** /api/interruptions/{id} | Get a InterruptionResource by ID
[**ApiInterruptionsIdResponsibleGet**](InterruptionsApi.md#ApiInterruptionsIdResponsibleGet) | **Get** /api/interruptions/{id}/responsible | 
[**ApiInterruptionsIdResponsiblePut**](InterruptionsApi.md#ApiInterruptionsIdResponsiblePut) | **Put** /api/interruptions/{id}/responsible | 
[**ApiInterruptionsIdSubmitPost**](InterruptionsApi.md#ApiInterruptionsIdSubmitPost) | **Post** /api/interruptions/{id}/submit | 


# **ApiInterruptionsGet**
> ApiInterruptionsGet()



Lists interruptions for user attention. The results will be sorted by date from most recently to least recently created.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiInterruptionsIdGet**
> InterruptionResource ApiInterruptionsIdGet($id)

Get a InterruptionResource by ID

Gets a single interruption by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the InterruptionResource to load | 

### Return type

[**InterruptionResource**](InterruptionResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiInterruptionsIdResponsibleGet**
> UserResource ApiInterruptionsIdResponsibleGet($id)



Gets the user that is currently responsible for this interruption.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**UserResource**](UserResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiInterruptionsIdResponsiblePut**
> UserResource ApiInterruptionsIdResponsiblePut($id)



Allows the current user to take responsibility for this interruption. Only users in one of the responsible teams on this interruption can take responsibility for it.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**UserResource**](UserResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiInterruptionsIdSubmitPost**
> InterruptionResource ApiInterruptionsIdSubmitPost($id)



Submits a dictionary of form values for the interruption. Only the user with responsibility for this interruption can submit this form.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**InterruptionResource**](InterruptionResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

