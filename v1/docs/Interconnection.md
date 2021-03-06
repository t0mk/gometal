# Interconnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ContactEmail** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Redundancy** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **int32** | The connection&#39;s speed in bps. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Ports** | Pointer to [**[]InterconnectionPort**](InterconnectionPort.md) |  | [optional] 
**Facility** | Pointer to [**Href**](Href.md) |  | [optional] 
**Organization** | Pointer to [**Href**](Href.md) |  | [optional] 
**Metro** | Pointer to [**Metro**](Metro.md) |  | [optional] 

## Methods

### NewInterconnection

`func NewInterconnection() *Interconnection`

NewInterconnection instantiates a new Interconnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectionWithDefaults

`func NewInterconnectionWithDefaults() *Interconnection`

NewInterconnectionWithDefaults instantiates a new Interconnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Interconnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Interconnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Interconnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Interconnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Interconnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Interconnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Interconnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Interconnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Interconnection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Interconnection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Interconnection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Interconnection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContactEmail

`func (o *Interconnection) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Interconnection) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Interconnection) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *Interconnection) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetStatus

`func (o *Interconnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Interconnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Interconnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Interconnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Interconnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Interconnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Interconnection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Interconnection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRedundancy

`func (o *Interconnection) GetRedundancy() string`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *Interconnection) GetRedundancyOk() (*string, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *Interconnection) SetRedundancy(v string)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *Interconnection) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetSpeed

`func (o *Interconnection) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *Interconnection) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *Interconnection) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *Interconnection) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *Interconnection) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Interconnection) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Interconnection) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Interconnection) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPorts

`func (o *Interconnection) GetPorts() []InterconnectionPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *Interconnection) GetPortsOk() (*[]InterconnectionPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *Interconnection) SetPorts(v []InterconnectionPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *Interconnection) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetFacility

`func (o *Interconnection) GetFacility() Href`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *Interconnection) GetFacilityOk() (*Href, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *Interconnection) SetFacility(v Href)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *Interconnection) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetOrganization

`func (o *Interconnection) GetOrganization() Href`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Interconnection) GetOrganizationOk() (*Href, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Interconnection) SetOrganization(v Href)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Interconnection) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetMetro

`func (o *Interconnection) GetMetro() Metro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *Interconnection) GetMetroOk() (*Metro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *Interconnection) SetMetro(v Metro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *Interconnection) HasMetro() bool`

HasMetro returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


