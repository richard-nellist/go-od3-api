# \UserRolesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiUserrolesAllGet**](UserRolesApi.md#ApiUserrolesAllGet) | **Get** /api/userroles/all | Get a list of UserRoleResources
[**ApiUserrolesGet**](UserRolesApi.md#ApiUserrolesGet) | **Get** /api/userroles | Get a list of UserRoleResources
[**ApiUserrolesIdDelete**](UserRolesApi.md#ApiUserrolesIdDelete) | **Delete** /api/userroles/{id} | Delete a UserRoleResource by ID
[**ApiUserrolesIdGet**](UserRolesApi.md#ApiUserrolesIdGet) | **Get** /api/userroles/{id} | Get a UserRoleResource by ID
[**ApiUserrolesIdPut**](UserRolesApi.md#ApiUserrolesIdPut) | **Put** /api/userroles/{id} | Modify a UserRoleResource by ID
[**ApiUserrolesPost**](UserRolesApi.md#ApiUserrolesPost) | **Post** /api/userroles | Create a UserRoleResource


# **ApiUserrolesAllGet**
> []UserRoleResource ApiUserrolesAllGet()

Get a list of UserRoleResources

Lists all of the user roles in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserRoleResource**](UserRoleResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUserrolesGet**
> ResourceCollectionUserRoleResource ApiUserrolesGet($skip, $take)

Get a list of UserRoleResources

Lists all of the user roles in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionUserRoleResource**](ResourceCollection[UserRoleResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUserrolesIdDelete**
> TaskResource ApiUserrolesIdDelete($id)

Delete a UserRoleResource by ID

Deletes an existing user role.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the UserRoleResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUserrolesIdGet**
> UserRoleResource ApiUserrolesIdGet($id)

Get a UserRoleResource by ID

Gets a single user role by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the UserRoleResource to load | 

### Return type

[**UserRoleResource**](UserRoleResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUserrolesIdPut**
> UserRoleResource ApiUserrolesIdPut($id, $userRoleResource)

Modify a UserRoleResource by ID

Modifies an existing user role definition.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the UserRoleResource to modify | 
 **userRoleResource** | [**UserRoleResource**](UserRoleResource.md)| The UserRoleResource resource to create | [optional] 

### Return type

[**UserRoleResource**](UserRoleResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUserrolesPost**
> UserRoleResource ApiUserrolesPost($userRoleResource)

Create a UserRoleResource

Creates a custom user role definition.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRoleResource** | [**UserRoleResource**](UserRoleResource.md)| The UserRoleResource resource to create | [optional] 

### Return type

[**UserRoleResource**](UserRoleResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

