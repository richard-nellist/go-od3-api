# \LetsEncryptApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiLetsencryptconfigurationDisablePost**](LetsEncryptApi.md#ApiLetsencryptconfigurationDisablePost) | **Post** /api/letsencryptconfiguration/disable | 
[**ApiLetsencryptconfigurationGet**](LetsEncryptApi.md#ApiLetsencryptconfigurationGet) | **Get** /api/letsencryptconfiguration | 
[**WellKnownAcmeChallengeTokenGet**](LetsEncryptApi.md#WellKnownAcmeChallengeTokenGet) | **Get** /.well-known/acme-challenge//{token} | 


# **ApiLetsencryptconfigurationDisablePost**
> ApiLetsencryptconfigurationDisablePost()



Disables integration with Let's Encrypt  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiLetsencryptconfigurationGet**
> ApiLetsencryptconfigurationGet()



Returns the current Let's Encrypt configuration  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **WellKnownAcmeChallengeTokenGet**
> WellKnownAcmeChallengeTokenGet($token)



Returns the computed HTTP challenge for a given token  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| LetsEncrypt response token | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

