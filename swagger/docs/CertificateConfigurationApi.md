# \CertificateConfigurationApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCertificatesIdICertificateGlobalcertificateAzureGet**](CertificateConfigurationApi.md#ApiCertificatesIdICertificateGlobalcertificateAzureGet) | **Get** /api/certificates/(?&lt;id&gt;(?i)^(certificate-global|certificate-azure)) | Get a CertificateConfigurationResource by ID
[**ApiConfigurationCertificatesGet**](CertificateConfigurationApi.md#ApiConfigurationCertificatesGet) | **Get** /api/configuration/certificates | Get a list of CertificateConfigurationResources
[**ApiConfigurationCertificatesIdGet**](CertificateConfigurationApi.md#ApiConfigurationCertificatesIdGet) | **Get** /api/configuration/certificates/{id} | Get a CertificateConfigurationResource by ID
[**ApiConfigurationCertificatesIdPublicCerGet**](CertificateConfigurationApi.md#ApiConfigurationCertificatesIdPublicCerGet) | **Get** /api/configuration/certificates/{id}/public-cer | 


# **ApiCertificatesIdICertificateGlobalcertificateAzureGet**
> CertificateConfigurationResource ApiCertificatesIdICertificateGlobalcertificateAzureGet($id)

Get a CertificateConfigurationResource by ID

Gets a certificate by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the CertificateConfigurationResource to load | 

### Return type

[**CertificateConfigurationResource**](CertificateConfigurationResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiConfigurationCertificatesGet**
> ResourceCollectionCertificateConfigurationResource ApiConfigurationCertificatesGet($skip, $take)

Get a list of CertificateConfigurationResources

Lists all of the X509 certificates in the current Octopus installation.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionCertificateConfigurationResource**](ResourceCollection[CertificateConfigurationResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiConfigurationCertificatesIdGet**
> CertificateConfigurationResource ApiConfigurationCertificatesIdGet($id)

Get a CertificateConfigurationResource by ID

Gets a certificate by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the CertificateConfigurationResource to load | 

### Return type

[**CertificateConfigurationResource**](CertificateConfigurationResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiConfigurationCertificatesIdPublicCerGet**
> *os.File ApiConfigurationCertificatesIdPublicCerGet($id)



Downloads the public portion of the certificate in .cer format  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

