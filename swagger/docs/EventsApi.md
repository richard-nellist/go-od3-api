# \EventsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEventsCategoriesGet**](EventsApi.md#ApiEventsCategoriesGet) | **Get** /api/events/categories | 
[**ApiEventsDocumenttypesGet**](EventsApi.md#ApiEventsDocumenttypesGet) | **Get** /api/events/documenttypes | 
[**ApiEventsGet**](EventsApi.md#ApiEventsGet) | **Get** /api/events | 
[**ApiEventsGroupsGet**](EventsApi.md#ApiEventsGroupsGet) | **Get** /api/events/groups | 
[**ApiEventsIdGet**](EventsApi.md#ApiEventsIdGet) | **Get** /api/events/{id} | Get a EventResource by ID


# **ApiEventsCategoriesGet**
> ApiEventsCategoriesGet()



Lists event categories.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEventsDocumenttypesGet**
> ApiEventsDocumenttypesGet()



Lists subscription event document types.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEventsGet**
> ApiEventsGet()



List all of the the audit events collected to date. Events can be filtered by various criteria. Events will be ordered by the date of the event, descending.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiEventsGroupsGet**
> ApiEventsGroupsGet()



Lists subscription event groups.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEventsIdGet**
> EventResource ApiEventsIdGet($id)

Get a EventResource by ID

Gets a single event by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the EventResource to load | 

### Return type

[**EventResource**](EventResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

