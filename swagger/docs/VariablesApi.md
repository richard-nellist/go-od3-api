# \VariablesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiVariablesIdGet**](VariablesApi.md#ApiVariablesIdGet) | **Get** /api/variables/{id} | Get a VariableSetResource by ID
[**ApiVariablesIdPut**](VariablesApi.md#ApiVariablesIdPut) | **Put** /api/variables/{id} | 
[**ApiVariablesNamesGet**](VariablesApi.md#ApiVariablesNamesGet) | **Get** /api/variables/names | 
[**ApiVariablesNonPrintableCharsPost**](VariablesApi.md#ApiVariablesNonPrintableCharsPost) | **Post** /api/variables/nonPrintableChars | 


# **ApiVariablesIdGet**
> VariableSetResource ApiVariablesIdGet($id)

Get a VariableSetResource by ID

Gets a variable set by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the VariableSetResource to load | 

### Return type

[**VariableSetResource**](VariableSetResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiVariablesIdPut**
> VariableSetResource ApiVariablesIdPut($id)



Updates a variable set.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**VariableSetResource**](VariableSetResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiVariablesNamesGet**
> ApiVariablesNamesGet()



List the names of variables that can be used in deployment actions. If a project is specified, this will include variables in that project.  If a project environments filter is specified, project variables only scoped to the environment will be excluded.   NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiVariablesNonPrintableCharsPost**
> []int32 ApiVariablesNonPrintableCharsPost()



Get list of non-printable variable characters.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

**[]int32**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

