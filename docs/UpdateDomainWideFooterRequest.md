# UpdateDomainWideFooterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | Pointer to [**UpdateDomainWideFooterRequestAttr**](UpdateDomainWideFooterRequestAttr.md) |  | [optional] 
**Items** | Pointer to **[]string** | contains a list of domain names where you want to update the footer | [optional] 

## Methods

### NewUpdateDomainWideFooterRequest

`func NewUpdateDomainWideFooterRequest() *UpdateDomainWideFooterRequest`

NewUpdateDomainWideFooterRequest instantiates a new UpdateDomainWideFooterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainWideFooterRequestWithDefaults

`func NewUpdateDomainWideFooterRequestWithDefaults() *UpdateDomainWideFooterRequest`

NewUpdateDomainWideFooterRequestWithDefaults instantiates a new UpdateDomainWideFooterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *UpdateDomainWideFooterRequest) GetAttr() UpdateDomainWideFooterRequestAttr`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *UpdateDomainWideFooterRequest) GetAttrOk() (*UpdateDomainWideFooterRequestAttr, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *UpdateDomainWideFooterRequest) SetAttr(v UpdateDomainWideFooterRequestAttr)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *UpdateDomainWideFooterRequest) HasAttr() bool`

HasAttr returns a boolean if a field has been set.

### GetItems

`func (o *UpdateDomainWideFooterRequest) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UpdateDomainWideFooterRequest) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UpdateDomainWideFooterRequest) SetItems(v []string)`

SetItems sets Items field to given value.

### HasItems

`func (o *UpdateDomainWideFooterRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


