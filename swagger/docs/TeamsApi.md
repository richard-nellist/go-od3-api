# \TeamsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTeamsAllGet**](TeamsApi.md#ApiTeamsAllGet) | **Get** /api/teams/all | Get a list of TeamResources
[**ApiTeamsGet**](TeamsApi.md#ApiTeamsGet) | **Get** /api/teams | Get a list of TeamResources
[**ApiTeamsIdDelete**](TeamsApi.md#ApiTeamsIdDelete) | **Delete** /api/teams/{id} | Delete a TeamResource by ID
[**ApiTeamsIdGet**](TeamsApi.md#ApiTeamsIdGet) | **Get** /api/teams/{id} | Get a TeamResource by ID
[**ApiTeamsIdPut**](TeamsApi.md#ApiTeamsIdPut) | **Put** /api/teams/{id} | Modify a TeamResource by ID
[**ApiTeamsPost**](TeamsApi.md#ApiTeamsPost) | **Post** /api/teams | Create a TeamResource


# **ApiTeamsAllGet**
> []TeamResource ApiTeamsAllGet()

Get a list of TeamResources

Lists the name and ID of all of the teams in the current Octopus installation. The results will be sorted by name.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]TeamResource**](TeamResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTeamsGet**
> ResourceCollectionTeamResource ApiTeamsGet($skip, $take)

Get a list of TeamResources

Lists all of the teams in the current Octopus installation. The results will be sorted alphabetically by name.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| Number of items to skip | [optional] 
 **take** | **int32**| Number of items to take | [optional] 

### Return type

[**ResourceCollectionTeamResource**](ResourceCollection[TeamResource].md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTeamsIdDelete**
> TaskResource ApiTeamsIdDelete($id)

Delete a TeamResource by ID

Deletes an existing team.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the TeamResource to delete | 

### Return type

[**TaskResource**](TaskResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTeamsIdGet**
> TeamResource ApiTeamsIdGet($id)

Get a TeamResource by ID

Gets a team by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the TeamResource to load | 

### Return type

[**TeamResource**](TeamResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTeamsIdPut**
> TeamResource ApiTeamsIdPut($id, $teamResource)

Modify a TeamResource by ID

Modifies an existing team.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the TeamResource to modify | 
 **teamResource** | [**TeamResource**](TeamResource.md)| The TeamResource resource to create | [optional] 

### Return type

[**TeamResource**](TeamResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTeamsPost**
> TeamResource ApiTeamsPost($teamResource)

Create a TeamResource

Creates a team.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamResource** | [**TeamResource**](TeamResource.md)| The TeamResource resource to create | [optional] 

### Return type

[**TeamResource**](TeamResource.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APIKeyQuery](../README.md#APIKeyQuery), [NugetApiKeyHeader](../README.md#NugetApiKeyHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

