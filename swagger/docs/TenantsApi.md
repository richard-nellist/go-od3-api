# \TenantsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTenantsAllGet**](TenantsApi.md#ApiTenantsAllGet) | **Get** /api/tenants/all | Get a list of TenantResources
[**ApiTenantsGet**](TenantsApi.md#ApiTenantsGet) | **Get** /api/tenants | Get a list of TenantResources
[**ApiTenantsIdDelete**](TenantsApi.md#ApiTenantsIdDelete) | **Delete** /api/tenants/{id} | Delete a TenantResource by ID
[**ApiTenantsIdGet**](TenantsApi.md#ApiTenantsIdGet) | **Get** /api/tenants/{id} | Get a TenantResource by ID
[**ApiTenantsIdLogoGet**](TenantsApi.md#ApiTenantsIdLogoGet) | **Get** /api/tenants/{id}/logo | 
[**ApiTenantsIdLogoPost**](TenantsApi.md#ApiTenantsIdLogoPost) | **Post** /api/tenants/{id}/logo | 
[**ApiTenantsIdLogoPut**](TenantsApi.md#ApiTenantsIdLogoPut) | **Put** /api/tenants/{id}/logo | 
[**ApiTenantsIdPut**](TenantsApi.md#ApiTenantsIdPut) | **Put** /api/tenants/{id} | Modify a TenantResource by ID
[**ApiTenantsIdVariablesGet**](TenantsApi.md#ApiTenantsIdVariablesGet) | **Get** /api/tenants/{id}/variables | 
[**ApiTenantsIdVariablesPost**](TenantsApi.md#ApiTenantsIdVariablesPost) | **Post** /api/tenants/{id}/variables | 
[**ApiTenantsIdVariablesPut**](TenantsApi.md#ApiTenantsIdVariablesPut) | **Put** /api/tenants/{id}/variables | 
[**ApiTenantsPost**](TenantsApi.md#ApiTenantsPost) | **Post** /api/tenants | Create a TenantResource
[**ApiTenantsTagTestGet**](TenantsApi.md#ApiTenantsTagTestGet) | **Get** /api/tenants/tag-test | 
[**ApiTenantsVariablesMissingGet**](TenantsApi.md#ApiTenantsVariablesMissingGet) | **Get** /api/tenants/variables-missing | 
[**ApiTenantsVariablesetTestGet**](TenantsApi.md#ApiTenantsVariablesetTestGet) | **Get** /api/tenants/variableset-test | 


# **ApiTenantsAllGet**
> []TenantResource ApiTenantsAllGet()

Get a list of TenantResources

Lists all of the tenants in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]TenantResource**](TenantResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTenantsGet**
> ResourceCollectionTenantResource ApiTenantsGet($skip, $take)

Get a list of TenantResources

Lists all of the tenants in the current Octopus installation. The results will be sorted alphabetically by name, and returned 30 at a time.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionTenantResource**](ResourceCollection[TenantResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTenantsIdDelete**
> TaskResource ApiTenantsIdDelete($id)

Delete a TenantResource by ID

Deletes an existing tenant.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the TenantResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTenantsIdGet**
> TenantResource ApiTenantsIdGet($id)

Get a TenantResource by ID

Gets a single tenant by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the TenantResource to load | 

### Return type

[**TenantResource**](TenantResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTenantsIdLogoGet**
> *os.File ApiTenantsIdLogoGet($id)



Gets the logo associated with the tenant.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiTenantsIdLogoPost**
> ApiTenantsIdLogoPost($id)



Updates the logo associated with the tenant.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiTenantsIdLogoPut**
> ApiTenantsIdLogoPut($id)



Updates the logo associated with the tenant.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiTenantsIdPut**
> TenantResource ApiTenantsIdPut($id, $tenantResource)

Modify a TenantResource by ID

Modifies an existing tenant.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the TenantResource to modify | 
 **tenantResource** | [**TenantResource**](TenantResource.md)| The TenantResource resource to create | [optional] 

### Return type

[**TenantResource**](TenantResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTenantsIdVariablesGet**
> TenantVariableResource ApiTenantsIdVariablesGet($id)



Gets all the available variables associated with the tenant.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**TenantVariableResource**](TenantVariableResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTenantsIdVariablesPost**
> TenantVariableResource ApiTenantsIdVariablesPost($id)



Updates the variables associated with the tenant.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**TenantVariableResource**](TenantVariableResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTenantsIdVariablesPut**
> TenantVariableResource ApiTenantsIdVariablesPut($id)



Updates the variables associated with the tenant.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**TenantVariableResource**](TenantVariableResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTenantsPost**
> TenantResource ApiTenantsPost($tenantResource)

Create a TenantResource

Creates a new tenant.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantResource** | [**TenantResource**](TenantResource.md)| The TenantResource resource to create | [optional] 

### Return type

[**TenantResource**](TenantResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTenantsTagTestGet**
> ApiTenantsTagTestGet()



Checks tenants for matching tags rule  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiTenantsVariablesMissingGet**
> ApiTenantsVariablesMissingGet()



Returns list of tenants who are missing required variables  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiTenantsVariablesetTestGet**
> ApiTenantsVariablesetTestGet()



Checks tenants for matching variable set rule  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

