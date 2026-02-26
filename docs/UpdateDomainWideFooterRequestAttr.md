# UpdateDomainWideFooterRequestAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Html** | Pointer to **string** | Footer text in HTML format | [optional] 
**Plain** | Pointer to **string** | Footer text in PLAIN text format | [optional] 
**MboxExclude** | Pointer to **map[string]interface{}** | Array of mailboxes to exclude from domain wide footer | [optional] 

## Methods

### NewUpdateDomainWideFooterRequestAttr

`func NewUpdateDomainWideFooterRequestAttr() *UpdateDomainWideFooterRequestAttr`

NewUpdateDomainWideFooterRequestAttr instantiates a new UpdateDomainWideFooterRequestAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainWideFooterRequestAttrWithDefaults

`func NewUpdateDomainWideFooterRequestAttrWithDefaults() *UpdateDomainWideFooterRequestAttr`

NewUpdateDomainWideFooterRequestAttrWithDefaults instantiates a new UpdateDomainWideFooterRequestAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHtml

`func (o *UpdateDomainWideFooterRequestAttr) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *UpdateDomainWideFooterRequestAttr) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *UpdateDomainWideFooterRequestAttr) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *UpdateDomainWideFooterRequestAttr) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetPlain

`func (o *UpdateDomainWideFooterRequestAttr) GetPlain() string`

GetPlain returns the Plain field if non-nil, zero value otherwise.

### GetPlainOk

`func (o *UpdateDomainWideFooterRequestAttr) GetPlainOk() (*string, bool)`

GetPlainOk returns a tuple with the Plain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlain

`func (o *UpdateDomainWideFooterRequestAttr) SetPlain(v string)`

SetPlain sets Plain field to given value.

### HasPlain

`func (o *UpdateDomainWideFooterRequestAttr) HasPlain() bool`

HasPlain returns a boolean if a field has been set.

### GetMboxExclude

`func (o *UpdateDomainWideFooterRequestAttr) GetMboxExclude() map[string]interface{}`

GetMboxExclude returns the MboxExclude field if non-nil, zero value otherwise.

### GetMboxExcludeOk

`func (o *UpdateDomainWideFooterRequestAttr) GetMboxExcludeOk() (*map[string]interface{}, bool)`

GetMboxExcludeOk returns a tuple with the MboxExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMboxExclude

`func (o *UpdateDomainWideFooterRequestAttr) SetMboxExclude(v map[string]interface{})`

SetMboxExclude sets MboxExclude field to given value.

### HasMboxExclude

`func (o *UpdateDomainWideFooterRequestAttr) HasMboxExclude() bool`

HasMboxExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


