# MetroServerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metro** | Pointer to **string** | The metro ID or code to check the capacity in. | [optional] 
**Plan** | Pointer to **string** | The plan ID or slug to check the capacity of. | [optional] 
**Quantity** | Pointer to **string** | The number of servers to check the capacity of. | [optional] 

## Methods

### NewMetroServerInfo

`func NewMetroServerInfo() *MetroServerInfo`

NewMetroServerInfo instantiates a new MetroServerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroServerInfoWithDefaults

`func NewMetroServerInfoWithDefaults() *MetroServerInfo`

NewMetroServerInfoWithDefaults instantiates a new MetroServerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetro

`func (o *MetroServerInfo) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *MetroServerInfo) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *MetroServerInfo) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *MetroServerInfo) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetPlan

`func (o *MetroServerInfo) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *MetroServerInfo) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *MetroServerInfo) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *MetroServerInfo) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetQuantity

`func (o *MetroServerInfo) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MetroServerInfo) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MetroServerInfo) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *MetroServerInfo) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


