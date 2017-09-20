# \CommunityActionTemplatesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCommunityactiontemplatesGet**](CommunityActionTemplatesApi.md#ApiCommunityactiontemplatesGet) | **Get** /api/communityactiontemplates | Get a list of CommunityActionTemplateResources
[**ApiCommunityactiontemplatesIdActiontemplateGet**](CommunityActionTemplatesApi.md#ApiCommunityactiontemplatesIdActiontemplateGet) | **Get** /api/communityactiontemplates/{id}/actiontemplate | 
[**ApiCommunityactiontemplatesIdGet**](CommunityActionTemplatesApi.md#ApiCommunityactiontemplatesIdGet) | **Get** /api/communityactiontemplates/{id} | Get a CommunityActionTemplateResource by ID
[**ApiCommunityactiontemplatesIdInstallationPost**](CommunityActionTemplatesApi.md#ApiCommunityactiontemplatesIdInstallationPost) | **Post** /api/communityactiontemplates/{id}/installation | 
[**ApiCommunityactiontemplatesIdInstallationPut**](CommunityActionTemplatesApi.md#ApiCommunityactiontemplatesIdInstallationPut) | **Put** /api/communityactiontemplates/{id}/installation | 
[**ApiCommunityactiontemplatesIdLogoGet**](CommunityActionTemplatesApi.md#ApiCommunityactiontemplatesIdLogoGet) | **Get** /api/communityactiontemplates/{id}/logo | 


# **ApiCommunityactiontemplatesGet**
> ResourceCollectionCommunityActionTemplateResource ApiCommunityactiontemplatesGet($skip, $take)

Get a list of CommunityActionTemplateResources




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionCommunityActionTemplateResource**](ResourceCollection[CommunityActionTemplateResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCommunityactiontemplatesIdActiontemplateGet**
> ActionTemplateResource ApiCommunityactiontemplatesIdActiontemplateGet($id)



Gets installed version of the template.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**ActionTemplateResource**](ActionTemplateResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCommunityactiontemplatesIdGet**
> CommunityActionTemplateResource ApiCommunityactiontemplatesIdGet($id)

Get a CommunityActionTemplateResource by ID

Gets a single community step template.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the CommunityActionTemplateResource to load | 

### Return type

[**CommunityActionTemplateResource**](CommunityActionTemplateResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCommunityactiontemplatesIdInstallationPost**
> ActionTemplateResource ApiCommunityactiontemplatesIdInstallationPost($id)



Installs community step template.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**ActionTemplateResource**](ActionTemplateResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCommunityactiontemplatesIdInstallationPut**
> ActionTemplateResource ApiCommunityactiontemplatesIdInstallationPut($id)



Updates installed community step template to the lastest version.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**ActionTemplateResource**](ActionTemplateResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCommunityactiontemplatesIdLogoGet**
> *os.File ApiCommunityactiontemplatesIdLogoGet($id)



Gets the logo associated with the community step template.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

