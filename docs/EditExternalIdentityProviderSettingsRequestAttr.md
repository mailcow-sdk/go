# EditExternalIdentityProviderSettingsRequestAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authsource** | Pointer to **string** | Specifies the type of the Identity Provider | [optional] 
**ServerUrl** | Pointer to **string** | The base URL of your Keycloak server. Required if &#x60;authsource&#x60; is keycloak. | [optional] 
**Realm** | Pointer to **string** | The Keycloak realm where the mailcow client is configured. Required if &#x60;authsource&#x60; is keycloak. | [optional] 
**ClientId** | Pointer to **string** | The Client ID assigned to mailcow Client in OIDC Provider. Required if &#x60;authsource&#x60; is keycloak or generic-oidc. | [optional] 
**ClientSecret** | Pointer to **string** | The Client Secret assigned to mailcow Client in OIDC Provider. Required if &#x60;authsource&#x60; is keycloak or generic-oidc. | [optional] 
**RedirectUrl** | Pointer to **string** | The redirect URL that OIDC Provider will use after authentication. Required if &#x60;authsource&#x60; is keycloak or generic-oidc. | [optional] 
**RedirectUrlExtra** | Pointer to **[]interface{}** | Additional redirect URLs that OIDC Provider can use after authentication if valid. | [optional] 
**Version** | Pointer to **string** | Specifies the Keycloak version. Required if &#x60;authsource&#x60; is keycloak. | [optional] 
**DefaultTemplate** | Pointer to **string** | (Optional) If no matching Attribute Mapping exists for a User, the default template will be used for creating the mailbox, but not for updating the mailbox. | [optional] 
**Mappers** | Pointer to **[]interface{}** | (Optional) Attribute values used to match a mailbox template. Each element corresponds to the respective index in the templates array (i.e., the first element matches the first element of templates, the second matches the second, and so on). | [optional] 
**Templates** | Pointer to **[]interface{}** | (Optional) Defines the mailbox templates to be assigned. Each element corresponds to the respective index in the &#x60;mappers&#x60; array. | [optional] 
**IgnoreSslError** | Pointer to **bool** | If enabled, SSL certificate validation is bypassed | [optional] [default to false]
**MailpasswordFlow** | Pointer to **bool** | If enabled, mailcow will attempt to validate user credentials using the Keycloak Admin REST API instead of relying solely on the Authorization Code Flow. | [optional] [default to false]
**PeriodicSync** | Pointer to **bool** | If enabled, mailcow periodically performs a full sync of all users from Keycloak or LDAP. | [optional] [default to false]
**ImportUsers** | Pointer to **bool** | If enabled, new users are automatically imported from Keycloak or LDAP into mailcow. | [optional] [default to false]
**SyncInterval** | Pointer to **float32** | Defines the time interval (in minutes) for periodic synchronization and user imports. | [optional] [default to 15]
**Host** | Pointer to **string** | The address of your LDAP server. You can provide a single hostname or a comma-separated list of hosts for fallback in case the primary server is unreachable. Required if &#x60;authsource&#x60; is ldap. | [optional] 
**Port** | Pointer to **string** | The port used to connect to the LDAP server. Required if &#x60;authsource&#x60; is ldap. | [optional] 
**UseSsl** | Pointer to **bool** | enable LDAPS connection. If Port is set to 389 it will be overriden to 636. | [optional] [default to false]
**UseTls** | Pointer to **bool** | enable TLS connection. TLS is recommended over SSL. SSL Ports cannot be used. | [optional] [default to false]
**Basedn** | Pointer to **string** | The Distinguished Name (DN) from which searches will be performed. Required if &#x60;authsource&#x60; is ldap. | [optional] 
**UsernameField** | Pointer to **string** | The LDAP attribute used to identify users during authentication. Required if &#x60;authsource&#x60; is ldap. | [optional] [default to "mail"]
**Filter** | Pointer to **string** | An optional LDAP search filter to refine which users can authenticate. | [optional] 
**AttributeField** | Pointer to **string** | Specifies an LDAP attribute that holds a specific value which can be mapped to a mailbox template using the Attribute Mapping section. Required if &#x60;authsource&#x60; is ldap. | [optional] 
**Binddn** | Pointer to **string** | The Distinguished Name (DN) of the LDAP user that will be used to authenticate and perform LDAP searches. This account should have sufficient permissions to read the required attributes. Required if &#x60;authsource&#x60; is ldap. | [optional] 
**Bindpass** | Pointer to **string** | The password for the Bind DN user. It is required for authentication when connecting to the LDAP server. Required if &#x60;authsource&#x60; is ldap. | [optional] 
**AuthorizeUrl** | Pointer to **string** | The OIDC provider&#39;s authorization server URL. Required if &#x60;authsource&#x60; is generic-oidc. | [optional] 
**TokenUrl** | Pointer to **string** | The OIDC provider&#39;s token server URL. Required if &#x60;authsource&#x60; is generic-oidc. | [optional] 
**UserinfoUrl** | Pointer to **string** | The OIDC provider&#39;s user info server URL. Required if &#x60;authsource&#x60; is generic-oidc. | [optional] 
**ClientScopes** | Pointer to **string** | Specifies the OIDC scopes requested during authentication. | [optional] [default to "openid profile email mailcow_template"]

## Methods

### NewEditExternalIdentityProviderSettingsRequestAttr

`func NewEditExternalIdentityProviderSettingsRequestAttr() *EditExternalIdentityProviderSettingsRequestAttr`

NewEditExternalIdentityProviderSettingsRequestAttr instantiates a new EditExternalIdentityProviderSettingsRequestAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditExternalIdentityProviderSettingsRequestAttrWithDefaults

`func NewEditExternalIdentityProviderSettingsRequestAttrWithDefaults() *EditExternalIdentityProviderSettingsRequestAttr`

NewEditExternalIdentityProviderSettingsRequestAttrWithDefaults instantiates a new EditExternalIdentityProviderSettingsRequestAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthsource

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetAuthsource() string`

GetAuthsource returns the Authsource field if non-nil, zero value otherwise.

### GetAuthsourceOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetAuthsourceOk() (*string, bool)`

GetAuthsourceOk returns a tuple with the Authsource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthsource

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetAuthsource(v string)`

SetAuthsource sets Authsource field to given value.

### HasAuthsource

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasAuthsource() bool`

HasAuthsource returns a boolean if a field has been set.

### GetServerUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.

### HasServerUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasServerUrl() bool`

HasServerUrl returns a boolean if a field has been set.

### GetRealm

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetClientId

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetRedirectUrlExtra

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetRedirectUrlExtra() []interface{}`

GetRedirectUrlExtra returns the RedirectUrlExtra field if non-nil, zero value otherwise.

### GetRedirectUrlExtraOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetRedirectUrlExtraOk() (*[]interface{}, bool)`

GetRedirectUrlExtraOk returns a tuple with the RedirectUrlExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrlExtra

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetRedirectUrlExtra(v []interface{})`

SetRedirectUrlExtra sets RedirectUrlExtra field to given value.

### HasRedirectUrlExtra

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasRedirectUrlExtra() bool`

HasRedirectUrlExtra returns a boolean if a field has been set.

### GetVersion

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDefaultTemplate

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetDefaultTemplate() string`

GetDefaultTemplate returns the DefaultTemplate field if non-nil, zero value otherwise.

### GetDefaultTemplateOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetDefaultTemplateOk() (*string, bool)`

GetDefaultTemplateOk returns a tuple with the DefaultTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTemplate

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetDefaultTemplate(v string)`

SetDefaultTemplate sets DefaultTemplate field to given value.

### HasDefaultTemplate

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasDefaultTemplate() bool`

HasDefaultTemplate returns a boolean if a field has been set.

### GetMappers

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetMappers() []interface{}`

GetMappers returns the Mappers field if non-nil, zero value otherwise.

### GetMappersOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetMappersOk() (*[]interface{}, bool)`

GetMappersOk returns a tuple with the Mappers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappers

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetMappers(v []interface{})`

SetMappers sets Mappers field to given value.

### HasMappers

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasMappers() bool`

HasMappers returns a boolean if a field has been set.

### GetTemplates

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetTemplates() []interface{}`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetTemplatesOk() (*[]interface{}, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetTemplates(v []interface{})`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetIgnoreSslError

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetIgnoreSslError() bool`

GetIgnoreSslError returns the IgnoreSslError field if non-nil, zero value otherwise.

### GetIgnoreSslErrorOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetIgnoreSslErrorOk() (*bool, bool)`

GetIgnoreSslErrorOk returns a tuple with the IgnoreSslError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSslError

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetIgnoreSslError(v bool)`

SetIgnoreSslError sets IgnoreSslError field to given value.

### HasIgnoreSslError

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasIgnoreSslError() bool`

HasIgnoreSslError returns a boolean if a field has been set.

### GetMailpasswordFlow

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetMailpasswordFlow() bool`

GetMailpasswordFlow returns the MailpasswordFlow field if non-nil, zero value otherwise.

### GetMailpasswordFlowOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetMailpasswordFlowOk() (*bool, bool)`

GetMailpasswordFlowOk returns a tuple with the MailpasswordFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailpasswordFlow

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetMailpasswordFlow(v bool)`

SetMailpasswordFlow sets MailpasswordFlow field to given value.

### HasMailpasswordFlow

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasMailpasswordFlow() bool`

HasMailpasswordFlow returns a boolean if a field has been set.

### GetPeriodicSync

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetPeriodicSync() bool`

GetPeriodicSync returns the PeriodicSync field if non-nil, zero value otherwise.

### GetPeriodicSyncOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetPeriodicSyncOk() (*bool, bool)`

GetPeriodicSyncOk returns a tuple with the PeriodicSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicSync

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetPeriodicSync(v bool)`

SetPeriodicSync sets PeriodicSync field to given value.

### HasPeriodicSync

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasPeriodicSync() bool`

HasPeriodicSync returns a boolean if a field has been set.

### GetImportUsers

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetImportUsers() bool`

GetImportUsers returns the ImportUsers field if non-nil, zero value otherwise.

### GetImportUsersOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetImportUsersOk() (*bool, bool)`

GetImportUsersOk returns a tuple with the ImportUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportUsers

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetImportUsers(v bool)`

SetImportUsers sets ImportUsers field to given value.

### HasImportUsers

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasImportUsers() bool`

HasImportUsers returns a boolean if a field has been set.

### GetSyncInterval

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetSyncInterval() float32`

GetSyncInterval returns the SyncInterval field if non-nil, zero value otherwise.

### GetSyncIntervalOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetSyncIntervalOk() (*float32, bool)`

GetSyncIntervalOk returns a tuple with the SyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncInterval

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetSyncInterval(v float32)`

SetSyncInterval sets SyncInterval field to given value.

### HasSyncInterval

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasSyncInterval() bool`

HasSyncInterval returns a boolean if a field has been set.

### GetHost

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUseSsl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetUseTls

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetBasedn

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetBasedn() string`

GetBasedn returns the Basedn field if non-nil, zero value otherwise.

### GetBasednOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetBasednOk() (*string, bool)`

GetBasednOk returns a tuple with the Basedn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasedn

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetBasedn(v string)`

SetBasedn sets Basedn field to given value.

### HasBasedn

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasBasedn() bool`

HasBasedn returns a boolean if a field has been set.

### GetUsernameField

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetUsernameField() string`

GetUsernameField returns the UsernameField field if non-nil, zero value otherwise.

### GetUsernameFieldOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetUsernameFieldOk() (*string, bool)`

GetUsernameFieldOk returns a tuple with the UsernameField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameField

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetUsernameField(v string)`

SetUsernameField sets UsernameField field to given value.

### HasUsernameField

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasUsernameField() bool`

HasUsernameField returns a boolean if a field has been set.

### GetFilter

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetAttributeField

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetAttributeField() string`

GetAttributeField returns the AttributeField field if non-nil, zero value otherwise.

### GetAttributeFieldOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetAttributeFieldOk() (*string, bool)`

GetAttributeFieldOk returns a tuple with the AttributeField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeField

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetAttributeField(v string)`

SetAttributeField sets AttributeField field to given value.

### HasAttributeField

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasAttributeField() bool`

HasAttributeField returns a boolean if a field has been set.

### GetBinddn

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.

### HasBinddn

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.

### GetBindpass

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetBindpass() string`

GetBindpass returns the Bindpass field if non-nil, zero value otherwise.

### GetBindpassOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetBindpassOk() (*string, bool)`

GetBindpassOk returns a tuple with the Bindpass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpass

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetBindpass(v string)`

SetBindpass sets Bindpass field to given value.

### HasBindpass

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasBindpass() bool`

HasBindpass returns a boolean if a field has been set.

### GetAuthorizeUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetAuthorizeUrl() string`

GetAuthorizeUrl returns the AuthorizeUrl field if non-nil, zero value otherwise.

### GetAuthorizeUrlOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetAuthorizeUrlOk() (*string, bool)`

GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetAuthorizeUrl(v string)`

SetAuthorizeUrl sets AuthorizeUrl field to given value.

### HasAuthorizeUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasAuthorizeUrl() bool`

HasAuthorizeUrl returns a boolean if a field has been set.

### GetTokenUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### GetUserinfoUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetUserinfoUrl() string`

GetUserinfoUrl returns the UserinfoUrl field if non-nil, zero value otherwise.

### GetUserinfoUrlOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetUserinfoUrlOk() (*string, bool)`

GetUserinfoUrlOk returns a tuple with the UserinfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserinfoUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetUserinfoUrl(v string)`

SetUserinfoUrl sets UserinfoUrl field to given value.

### HasUserinfoUrl

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasUserinfoUrl() bool`

HasUserinfoUrl returns a boolean if a field has been set.

### GetClientScopes

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetClientScopes() string`

GetClientScopes returns the ClientScopes field if non-nil, zero value otherwise.

### GetClientScopesOk

`func (o *EditExternalIdentityProviderSettingsRequestAttr) GetClientScopesOk() (*string, bool)`

GetClientScopesOk returns a tuple with the ClientScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientScopes

`func (o *EditExternalIdentityProviderSettingsRequestAttr) SetClientScopes(v string)`

SetClientScopes sets ClientScopes field to given value.

### HasClientScopes

`func (o *EditExternalIdentityProviderSettingsRequestAttr) HasClientScopes() bool`

HasClientScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


