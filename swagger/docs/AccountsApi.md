# \AccountsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountsAccountIdWebsitesGet**](AccountsApi.md#ApiAccountsAccountIdWebsitesGet) | **Get** /api/accounts/{accountId}/websites | 
[**ApiAccountsAllGet**](AccountsApi.md#ApiAccountsAllGet) | **Get** /api/accounts/all | Get a list of AccountResources
[**ApiAccountsGet**](AccountsApi.md#ApiAccountsGet) | **Get** /api/accounts | Get a list of AccountResources
[**ApiAccountsIdCloudServicesGet**](AccountsApi.md#ApiAccountsIdCloudServicesGet) | **Get** /api/accounts/{id}/cloudServices | 
[**ApiAccountsIdDelete**](AccountsApi.md#ApiAccountsIdDelete) | **Delete** /api/accounts/{id} | Delete a AccountResource by ID
[**ApiAccountsIdGet**](AccountsApi.md#ApiAccountsIdGet) | **Get** /api/accounts/{id} | Get a AccountResource by ID
[**ApiAccountsIdPkGet**](AccountsApi.md#ApiAccountsIdPkGet) | **Get** /api/accounts/{id}/pk | 
[**ApiAccountsIdPut**](AccountsApi.md#ApiAccountsIdPut) | **Put** /api/accounts/{id} | Modify a AccountResource by ID
[**ApiAccountsIdResourceGroupsGet**](AccountsApi.md#ApiAccountsIdResourceGroupsGet) | **Get** /api/accounts/{id}/resourceGroups | 
[**ApiAccountsIdStorageAccountsGet**](AccountsApi.md#ApiAccountsIdStorageAccountsGet) | **Get** /api/accounts/{id}/storageAccounts | 
[**ApiAccountsPost**](AccountsApi.md#ApiAccountsPost) | **Post** /api/accounts | Create a AccountResource


# **ApiAccountsAccountIdWebsitesGet**
> []AzureWebSite ApiAccountsAccountIdWebsitesGet($accountId)



Lists the websites associated with an Azure account.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Account Id for Azure | 

### Return type

[**[]AzureWebSite**](AzureWebSite.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAccountsAllGet**
> []AccountResource ApiAccountsAllGet()

Get a list of AccountResources

Lists all of the accounts in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]AccountResource**](AccountResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAccountsGet**
> ResourceCollectionAccountResource ApiAccountsGet($skip, $take)

Get a list of AccountResources

Lists accounts in the current Octopus installation in pages. The results will be sorted alphabetically by name.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionAccountResource**](ResourceCollection[AccountResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAccountsIdCloudServicesGet**
> []AzureCloudServiceResource ApiAccountsIdCloudServicesGet($id)



Lists the cloud services associated with an Azure account.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**[]AzureCloudServiceResource**](AzureCloudServiceResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAccountsIdDelete**
> TaskResource ApiAccountsIdDelete($id)

Delete a AccountResource by ID

Deletes an existing account.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the AccountResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAccountsIdGet**
> AccountResource ApiAccountsIdGet($id)

Get a AccountResource by ID

Gets a single account by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the AccountResource to load | 

### Return type

[**AccountResource**](AccountResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAccountsIdPkGet**
> *os.File ApiAccountsIdPkGet($id)



Downloads the public key portion of the account's associated certificate, if present.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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
 - **Accept**: application/x-x509-ca-cert

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAccountsIdPut**
> AccountResource ApiAccountsIdPut($id, $accountResource)

Modify a AccountResource by ID

Modifies an existing account.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the AccountResource to modify | 
 **accountResource** | [**AccountResource**](AccountResource.md)| The AccountResource resource to create | [optional] 

### Return type

[**AccountResource**](AccountResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAccountsIdResourceGroupsGet**
> []AzureResourceGroupResource ApiAccountsIdResourceGroupsGet($id)



Lists the Resource Groups associated with an Azure account.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**[]AzureResourceGroupResource**](AzureResourceGroupResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAccountsIdStorageAccountsGet**
> []AzureCloudServiceResource ApiAccountsIdStorageAccountsGet($id)



Lists the storage accounts asssociated with an Azure account.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**[]AzureCloudServiceResource**](AzureCloudServiceResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAccountsPost**
> AccountResource ApiAccountsPost($accountResource)

Create a AccountResource

Creates a new account.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountResource** | [**AccountResource**](AccountResource.md)| The AccountResource resource to create | [optional] 

### Return type

[**AccountResource**](AccountResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

