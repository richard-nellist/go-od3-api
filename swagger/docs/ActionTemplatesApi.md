# \ActionTemplatesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiActiontemplatesAllGet**](ActionTemplatesApi.md#ApiActiontemplatesAllGet) | **Get** /api/actiontemplates/all | Get a list of ActionTemplateResources
[**ApiActiontemplatesGet**](ActionTemplatesApi.md#ApiActiontemplatesGet) | **Get** /api/actiontemplates | Get a list of ActionTemplateResources
[**ApiActiontemplatesIdActionsUpdatePost**](ActionTemplatesApi.md#ApiActiontemplatesIdActionsUpdatePost) | **Post** /api/actiontemplates/{id}/actionsUpdate | 
[**ApiActiontemplatesIdDelete**](ActionTemplatesApi.md#ApiActiontemplatesIdDelete) | **Delete** /api/actiontemplates/{id} | Delete a ActionTemplateResource by ID
[**ApiActiontemplatesIdGet**](ActionTemplatesApi.md#ApiActiontemplatesIdGet) | **Get** /api/actiontemplates/{id} | Get a ActionTemplateResource by ID
[**ApiActiontemplatesIdLogoPost**](ActionTemplatesApi.md#ApiActiontemplatesIdLogoPost) | **Post** /api/actiontemplates/{id}/logo | 
[**ApiActiontemplatesIdLogoPut**](ActionTemplatesApi.md#ApiActiontemplatesIdLogoPut) | **Put** /api/actiontemplates/{id}/logo | 
[**ApiActiontemplatesIdPut**](ActionTemplatesApi.md#ApiActiontemplatesIdPut) | **Put** /api/actiontemplates/{id} | Modify a ActionTemplateResource by ID
[**ApiActiontemplatesIdUsageGet**](ActionTemplatesApi.md#ApiActiontemplatesIdUsageGet) | **Get** /api/actiontemplates/{id}/usage | 
[**ApiActiontemplatesIdVersionsVersionGet**](ActionTemplatesApi.md#ApiActiontemplatesIdVersionsVersionGet) | **Get** /api/actiontemplates/{id}/versions/{version?} | 
[**ApiActiontemplatesPost**](ActionTemplatesApi.md#ApiActiontemplatesPost) | **Post** /api/actiontemplates | Create a ActionTemplateResource
[**ApiActiontemplatesSearchGet**](ActionTemplatesApi.md#ApiActiontemplatesSearchGet) | **Get** /api/actiontemplates/search | 
[**ApiActiontemplatesTypeOrIdLogoGet**](ActionTemplatesApi.md#ApiActiontemplatesTypeOrIdLogoGet) | **Get** /api/actiontemplates/{typeOrId}/logo | 
[**ApiActiontemplatesTypeOrIdVersionsVersionLogoGet**](ActionTemplatesApi.md#ApiActiontemplatesTypeOrIdVersionsVersionLogoGet) | **Get** /api/actiontemplates/{typeOrId}/versions/{version}/logo | 


# **ApiActiontemplatesAllGet**
> []ActionTemplateResource ApiActiontemplatesAllGet()

Get a list of ActionTemplateResources

Lists the all of the action templates in the current Octopus installation. The results will be sorted by name.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]ActionTemplateResource**](ActionTemplateResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiActiontemplatesGet**
> ResourceCollectionActionTemplateResource ApiActiontemplatesGet($skip, $take)

Get a list of ActionTemplateResources

Lists all of the action templates in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionActionTemplateResource**](ResourceCollection[ActionTemplateResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiActiontemplatesIdActionsUpdatePost**
> []ActionUpdateResultResource ApiActiontemplatesIdActionsUpdatePost($id)



Updates deployment actions to specific version of action template  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**[]ActionUpdateResultResource**](ActionUpdateResultResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiActiontemplatesIdDelete**
> TaskResource ApiActiontemplatesIdDelete($id)

Delete a ActionTemplateResource by ID

Deletes an existing action template and all its versions.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ActionTemplateResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiActiontemplatesIdGet**
> ActionTemplateResource ApiActiontemplatesIdGet($id)

Get a ActionTemplateResource by ID

Gets a single action template by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ActionTemplateResource to load | 

### Return type

[**ActionTemplateResource**](ActionTemplateResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiActiontemplatesIdLogoPost**
> ApiActiontemplatesIdLogoPost($id)



Updates the logo associated with the latest version of the action template.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiActiontemplatesIdLogoPut**
> ApiActiontemplatesIdLogoPut($id)



Updates the logo associated with the latest version of the action template.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiActiontemplatesIdPut**
> ActionTemplateResource ApiActiontemplatesIdPut($id, $actionTemplateResource)

Modify a ActionTemplateResource by ID

Modifies an existing action template.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the ActionTemplateResource to modify | 
 **actionTemplateResource** | [**ActionTemplateResource**](ActionTemplateResource.md)| The ActionTemplateResource resource to create | [optional] 

### Return type

[**ActionTemplateResource**](ActionTemplateResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiActiontemplatesIdUsageGet**
> []ActionTemplateUsageResource ApiActiontemplatesIdUsageGet($id)



Gets a list of all steps/deployment processes that use a given action template.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**[]ActionTemplateUsageResource**](ActionTemplateUsageResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiActiontemplatesIdVersionsVersionGet**
> ApiActiontemplatesIdVersionsVersionGet($version, $id)



Gets specific version of action template.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **int32**| ERROR - NOT DEFINED | 
 **id** | **string**| ID of the resource | 

### Return type

void (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiActiontemplatesPost**
> ActionTemplateResource ApiActiontemplatesPost($actionTemplateResource)

Create a ActionTemplateResource

Creates a new action template.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionTemplateResource** | [**ActionTemplateResource**](ActionTemplateResource.md)| The ActionTemplateResource resource to create | [optional] 

### Return type

[**ActionTemplateResource**](ActionTemplateResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiActiontemplatesSearchGet**
> []ActionTemplateSearchResource ApiActiontemplatesSearchGet()



Lists all available action templates including built-in, custom and community contributed step templates.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]ActionTemplateSearchResource**](ActionTemplateSearchResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiActiontemplatesTypeOrIdLogoGet**
> *os.File ApiActiontemplatesTypeOrIdLogoGet($typeOrId, $version, $id)



Gets the logo associated with the latest version of action template.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typeOrId** | **string**| Action Type or ID of the action type logo | 
 **version** | **int32**| Version of the action type logo | 
 **id** | **string**| ID of the resource | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiActiontemplatesTypeOrIdVersionsVersionLogoGet**
> *os.File ApiActiontemplatesTypeOrIdVersionsVersionLogoGet($typeOrId, $version, $id)



Gets the logo associated with specific version of the action template.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typeOrId** | **string**| Action Type or ID of the action type logo | 
 **version** | **int32**| Version of the action type logo | 
 **id** | **string**| ID of the resource | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

