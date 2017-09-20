# DeploymentTemplateResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**IsDeploymentProcessModified** | **bool** |  | [optional] [default to null]
**IsVariableSetModified** | **bool** |  | [optional] [default to null]
**IsLibraryVariableSetModified** | **bool** |  | [optional] [default to null]
**DeploymentNotes** | **string** |  | [optional] [default to null]
**PromoteTo** | [**[]DeploymentPromotionTarget**](DeploymentPromotionTarget.md) |  | [optional] [default to null]
**TenantPromotions** | [**[]DeploymentPromomotionTenant**](DeploymentPromomotionTenant.md) |  | [optional] [default to null]
**LastModifiedOn** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**LastModifiedBy** | **string** |  | [optional] [default to null]
**Links** | **map[string]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


