# \CertificatesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCertificatesGet**](CertificatesApi.md#ApiCertificatesGet) | **Get** /api/certificates | 
[**ApiCertificatesIdArchivePost**](CertificatesApi.md#ApiCertificatesIdArchivePost) | **Post** /api/certificates/{id}/archive | 
[**ApiCertificatesIdDelete**](CertificatesApi.md#ApiCertificatesIdDelete) | **Delete** /api/certificates/{id} | Delete a CertificateResource by ID
[**ApiCertificatesIdExportGet**](CertificatesApi.md#ApiCertificatesIdExportGet) | **Get** /api/certificates/{id}/export | 
[**ApiCertificatesIdOrThumbprintICertificateGlobalcertificateAzureGet**](CertificatesApi.md#ApiCertificatesIdOrThumbprintICertificateGlobalcertificateAzureGet) | **Get** /api/certificates/(?&lt;idOrThumbprint&gt;(?i)^(?!(certificate-global|certificate-azure)).*) | 
[**ApiCertificatesIdPut**](CertificatesApi.md#ApiCertificatesIdPut) | **Put** /api/certificates/{id} | Modify a CertificateResource by ID
[**ApiCertificatesIdReplacePost**](CertificatesApi.md#ApiCertificatesIdReplacePost) | **Post** /api/certificates/{id}/replace | 
[**ApiCertificatesIdUnarchivePost**](CertificatesApi.md#ApiCertificatesIdUnarchivePost) | **Post** /api/certificates/{id}/unarchive | 
[**ApiCertificatesIdUsageGet**](CertificatesApi.md#ApiCertificatesIdUsageGet) | **Get** /api/certificates/{id}/usage | 
[**ApiCertificatesPost**](CertificatesApi.md#ApiCertificatesPost) | **Post** /api/certificates | Create a CertificateResource


# **ApiCertificatesGet**
> ApiCertificatesGet()



Lists X.509 certificates managed by Octopus  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiCertificatesIdArchivePost**
> ApiCertificatesIdArchivePost($id)



Archives a certificate  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiCertificatesIdDelete**
> TaskResource ApiCertificatesIdDelete($id)

Delete a CertificateResource by ID

Permanently deletes a certificate


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the CertificateResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCertificatesIdExportGet**
> *os.File ApiCertificatesIdExportGet($id)



Exports the certificate  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCertificatesIdOrThumbprintICertificateGlobalcertificateAzureGet**
> ApiCertificatesIdOrThumbprintICertificateGlobalcertificateAzureGet()



Get a certificate by ID or thumbprint  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiCertificatesIdPut**
> CertificateResource ApiCertificatesIdPut($id, $certificateResource)

Modify a CertificateResource by ID

Modifies an existing certificate


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the CertificateResource to modify | 
 **certificateResource** | [**CertificateResource**](CertificateResource.md)| The CertificateResource resource to create | [optional] 

### Return type

[**CertificateResource**](CertificateResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCertificatesIdReplacePost**
> Certificate ApiCertificatesIdReplacePost($id)



Replaces a certificate with another  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**Certificate**](Certificate.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCertificatesIdUnarchivePost**
> ApiCertificatesIdUnarchivePost($id)



Un-archives a certificate  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiCertificatesIdUsageGet**
> ApiCertificatesIdUsageGet()



Get the usages of a certificate  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiCertificatesPost**
> CertificateResource ApiCertificatesPost($certificateResource)

Create a CertificateResource

Adds a new certificate


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateResource** | [**CertificateResource**](CertificateResource.md)| The CertificateResource resource to create | [optional] 

### Return type

[**CertificateResource**](CertificateResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

