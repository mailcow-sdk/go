# EditExternalIdentityProviderSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **[]interface{}** |  | [optional] [default to {"identity-provider"}]
**Attr** | Pointer to [**EditExternalIdentityProviderSettingsRequestAttr**](EditExternalIdentityProviderSettingsRequestAttr.md) |  | [optional] 

## Methods

### NewEditExternalIdentityProviderSettingsRequest

`func NewEditExternalIdentityProviderSettingsRequest() *EditExternalIdentityProviderSettingsRequest`

NewEditExternalIdentityProviderSettingsRequest instantiates a new EditExternalIdentityProviderSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditExternalIdentityProviderSettingsRequestWithDefaults

`func NewEditExternalIdentityProviderSettingsRequestWithDefaults() *EditExternalIdentityProviderSettingsRequest`

NewEditExternalIdentityProviderSettingsRequestWithDefaults instantiates a new EditExternalIdentityProviderSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EditExternalIdentityProviderSettingsRequest) GetItems() []interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EditExternalIdentityProviderSettingsRequest) GetItemsOk() (*[]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EditExternalIdentityProviderSettingsRequest) SetItems(v []interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *EditExternalIdentityProviderSettingsRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAttr

`func (o *EditExternalIdentityProviderSettingsRequest) GetAttr() EditExternalIdentityProviderSettingsRequestAttr`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *EditExternalIdentityProviderSettingsRequest) GetAttrOk() (*EditExternalIdentityProviderSettingsRequestAttr, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *EditExternalIdentityProviderSettingsRequest) SetAttr(v EditExternalIdentityProviderSettingsRequestAttr)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *EditExternalIdentityProviderSettingsRequest) HasAttr() bool`

HasAttr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


