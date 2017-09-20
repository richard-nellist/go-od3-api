# \ReportingApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiReportingDeploymentsCountedByWeekGet**](ReportingApi.md#ApiReportingDeploymentsCountedByWeekGet) | **Get** /api/reporting/deployments-counted-by-week | 
[**ApiReportingDeploymentsXmlGet**](ReportingApi.md#ApiReportingDeploymentsXmlGet) | **Get** /api/reporting/deployments/xml | 


# **ApiReportingDeploymentsCountedByWeekGet**
> []ReportDeploymentCountOverTimeResource ApiReportingDeploymentsCountedByWeekGet()



Provides a report summarizing the weekly deployments per project over the last 6 months  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]ReportDeploymentCountOverTimeResource**](ReportDeploymentCountOverTimeResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiReportingDeploymentsXmlGet**
> ApiReportingDeploymentsXmlGet()



Provides an XML report that contains all of the information about deployments  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

