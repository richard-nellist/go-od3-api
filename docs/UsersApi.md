# \UsersApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiUsersAllGet**](UsersApi.md#ApiUsersAllGet) | **Get** /api/users/all | Get a list of UserResources
[**ApiUsersExternalSearchGet**](UsersApi.md#ApiUsersExternalSearchGet) | **Get** /api/users/external-search | 
[**ApiUsersGet**](UsersApi.md#ApiUsersGet) | **Get** /api/users | Get a list of UserResources
[**ApiUsersIdDelete**](UsersApi.md#ApiUsersIdDelete) | **Delete** /api/users/{id} | Delete a UserResource by ID
[**ApiUsersIdGet**](UsersApi.md#ApiUsersIdGet) | **Get** /api/users/{id} | Get a UserResource by ID
[**ApiUsersIdPermissionsExportGet**](UsersApi.md#ApiUsersIdPermissionsExportGet) | **Get** /api/users/{id}/permissions/export | 
[**ApiUsersIdPermissionsGet**](UsersApi.md#ApiUsersIdPermissionsGet) | **Get** /api/users/{id}/permissions | 
[**ApiUsersIdPut**](UsersApi.md#ApiUsersIdPut) | **Put** /api/users/{id} | 
[**ApiUsersIdentityMetadataGet**](UsersApi.md#ApiUsersIdentityMetadataGet) | **Get** /api/users/identity-metadata | 
[**ApiUsersInvitationsIdGet**](UsersApi.md#ApiUsersInvitationsIdGet) | **Get** /api/users/invitations/{id} | Get a InvitationResource by ID
[**ApiUsersInvitationsPost**](UsersApi.md#ApiUsersInvitationsPost) | **Post** /api/users/invitations | Create a InvitationResource
[**ApiUsersLoginPost**](UsersApi.md#ApiUsersLoginPost) | **Post** /api/users/login | 
[**ApiUsersLogoutPost**](UsersApi.md#ApiUsersLogoutPost) | **Post** /api/users/logout | 
[**ApiUsersMeGet**](UsersApi.md#ApiUsersMeGet) | **Get** /api/users/me | 
[**ApiUsersPost**](UsersApi.md#ApiUsersPost) | **Post** /api/users | 
[**ApiUsersRegisterPost**](UsersApi.md#ApiUsersRegisterPost) | **Post** /api/users/register | 


# **ApiUsersAllGet**
> []UserResource ApiUsersAllGet()

Get a list of UserResources

Lists all of the users in the current Octopus installation. The results will be sorted alphabetically by `Username`.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserResource**](UserResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersExternalSearchGet**
> ApiUsersExternalSearchGet()



Searches for users, using the authentication providers.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiUsersGet**
> ResourceCollectionUserResource ApiUsersGet($skip, $take)

Get a list of UserResources

Lists all of the users in the current Octopus installation, from all teams. The results will be sorted alphabetically by username.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionUserResource**](ResourceCollection[UserResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersIdDelete**
> TaskResource ApiUsersIdDelete($id)

Delete a UserResource by ID

Delete an existing user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the UserResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersIdGet**
> UserResource ApiUsersIdGet($id)

Get a UserResource by ID

Gets a single user by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the UserResource to load | 

### Return type

[**UserResource**](UserResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersIdPermissionsExportGet**
> ApiUsersIdPermissionsExportGet()



Gets list of permissions as a csv file. Available for the current authenticated user only.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiUsersIdPermissionsGet**
> UserPermissionSetResource ApiUsersIdPermissionsGet($id)



Gets summarized permission information. Available for the current authenticated user only.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**UserPermissionSetResource**](UserPermissionSetResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersIdPut**
> UserResource ApiUsersIdPut($id)



Modifies an existing user.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the resource | 

### Return type

[**UserResource**](UserResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersIdentityMetadataGet**
> ApiUsersIdentityMetadataGet()



Gets the metadata to describe the claims/fields used by authentication providers that support identities.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiUsersInvitationsIdGet**
> InvitationResource ApiUsersInvitationsIdGet($id)

Get a InvitationResource by ID

Gets an existing invitation by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the InvitationResource to load | 

### Return type

[**InvitationResource**](InvitationResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersInvitationsPost**
> InvitationResource ApiUsersInvitationsPost($invitationResource)

Create a InvitationResource

Invite a user to register.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationResource** | [**InvitationResource**](InvitationResource.md)| The InvitationResource resource to create | [optional] 

### Return type

[**InvitationResource**](InvitationResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersLoginPost**
> ApiUsersLoginPost()



Authenticates a user and returns a response with a cookie for the current user. This cookie can be submitted with future requests to avoid re-authentication.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersLogoutPost**
> ApiUsersLogoutPost()



Revokes the authentication cookie from the current session.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersMeGet**
> ApiUsersMeGet()



Gets information about the current user.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiUsersPost**
> ApiUsersPost()



Creates a new user.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


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

# **ApiUsersRegisterPost**
> ApiUsersRegisterPost()



Registers a new user and responds with an authentication cookie. Unless the first administrator user is being registered, an invitation code must be provided.  NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

