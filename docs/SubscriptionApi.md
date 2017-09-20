# \SubscriptionApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSubscriptionsAllGet**](SubscriptionApi.md#ApiSubscriptionsAllGet) | **Get** /api/subscriptions/all | Get a list of SubscriptionResources
[**ApiSubscriptionsGet**](SubscriptionApi.md#ApiSubscriptionsGet) | **Get** /api/subscriptions | Get a list of SubscriptionResources
[**ApiSubscriptionsIdDelete**](SubscriptionApi.md#ApiSubscriptionsIdDelete) | **Delete** /api/subscriptions/{id} | Delete a SubscriptionResource by ID
[**ApiSubscriptionsIdGet**](SubscriptionApi.md#ApiSubscriptionsIdGet) | **Get** /api/subscriptions/{id} | Get a SubscriptionResource by ID
[**ApiSubscriptionsIdPut**](SubscriptionApi.md#ApiSubscriptionsIdPut) | **Put** /api/subscriptions/{id} | Modify a SubscriptionResource by ID
[**ApiSubscriptionsPost**](SubscriptionApi.md#ApiSubscriptionsPost) | **Post** /api/subscriptions | Create a SubscriptionResource


# **ApiSubscriptionsAllGet**
> []SubscriptionResource ApiSubscriptionsAllGet()

Get a list of SubscriptionResources

Lists all the subscriptions in the Octopus Deploy installation.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]SubscriptionResource**](SubscriptionResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiSubscriptionsGet**
> ResourceCollectionSubscriptionResource ApiSubscriptionsGet($skip, $take)

Get a list of SubscriptionResources

Lists all of the subscriptions in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionSubscriptionResource**](ResourceCollection[SubscriptionResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiSubscriptionsIdDelete**
> TaskResource ApiSubscriptionsIdDelete($id)

Delete a SubscriptionResource by ID

Deletes an existing subscription.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the SubscriptionResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiSubscriptionsIdGet**
> SubscriptionResource ApiSubscriptionsIdGet($id)

Get a SubscriptionResource by ID

Get a subscription


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the SubscriptionResource to load | 

### Return type

[**SubscriptionResource**](SubscriptionResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiSubscriptionsIdPut**
> SubscriptionResource ApiSubscriptionsIdPut($id, $subscriptionResource)

Modify a SubscriptionResource by ID

Updates an existing subscription


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the SubscriptionResource to modify | 
 **subscriptionResource** | [**SubscriptionResource**](SubscriptionResource.md)| The SubscriptionResource resource to create | [optional] 

### Return type

[**SubscriptionResource**](SubscriptionResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiSubscriptionsPost**
> SubscriptionResource ApiSubscriptionsPost($subscriptionResource)

Create a SubscriptionResource

Creates a new subscription


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionResource** | [**SubscriptionResource**](SubscriptionResource.md)| The SubscriptionResource resource to create | [optional] 

### Return type

[**SubscriptionResource**](SubscriptionResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

