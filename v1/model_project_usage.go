/*
 * Metal API
 *
 * This is the API for Equinix Metal Product. Interact with your devices, user account, and projects.
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// ProjectUsage struct for ProjectUsage
type ProjectUsage struct {
	Facility *string `json:"facility,omitempty"`
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`
	Plan *string `json:"plan,omitempty"`
	PlanVersion *string `json:"plan_version,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	Unit *string `json:"unit,omitempty"`
	Price *string `json:"price,omitempty"`
	Total *string `json:"total,omitempty"`
}

// NewProjectUsage instantiates a new ProjectUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectUsage() *ProjectUsage {
	this := ProjectUsage{}
	return &this
}

// NewProjectUsageWithDefaults instantiates a new ProjectUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectUsageWithDefaults() *ProjectUsage {
	this := ProjectUsage{}
	return &this
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *ProjectUsage) GetFacility() string {
	if o == nil || o.Facility == nil {
		var ret string
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetFacilityOk() (*string, bool) {
	if o == nil || o.Facility == nil {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *ProjectUsage) HasFacility() bool {
	if o != nil && o.Facility != nil {
		return true
	}

	return false
}

// SetFacility gets a reference to the given string and assigns it to the Facility field.
func (o *ProjectUsage) SetFacility(v string) {
	o.Facility = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProjectUsage) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProjectUsage) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProjectUsage) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectUsage) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectUsage) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectUsage) SetName(v string) {
	o.Name = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *ProjectUsage) GetPlan() string {
	if o == nil || o.Plan == nil {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetPlanOk() (*string, bool) {
	if o == nil || o.Plan == nil {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *ProjectUsage) HasPlan() bool {
	if o != nil && o.Plan != nil {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *ProjectUsage) SetPlan(v string) {
	o.Plan = &v
}

// GetPlanVersion returns the PlanVersion field value if set, zero value otherwise.
func (o *ProjectUsage) GetPlanVersion() string {
	if o == nil || o.PlanVersion == nil {
		var ret string
		return ret
	}
	return *o.PlanVersion
}

// GetPlanVersionOk returns a tuple with the PlanVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetPlanVersionOk() (*string, bool) {
	if o == nil || o.PlanVersion == nil {
		return nil, false
	}
	return o.PlanVersion, true
}

// HasPlanVersion returns a boolean if a field has been set.
func (o *ProjectUsage) HasPlanVersion() bool {
	if o != nil && o.PlanVersion != nil {
		return true
	}

	return false
}

// SetPlanVersion gets a reference to the given string and assigns it to the PlanVersion field.
func (o *ProjectUsage) SetPlanVersion(v string) {
	o.PlanVersion = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ProjectUsage) GetQuantity() string {
	if o == nil || o.Quantity == nil {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetQuantityOk() (*string, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ProjectUsage) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *ProjectUsage) SetQuantity(v string) {
	o.Quantity = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ProjectUsage) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ProjectUsage) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *ProjectUsage) SetUnit(v string) {
	o.Unit = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProjectUsage) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetPriceOk() (*string, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ProjectUsage) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *ProjectUsage) SetPrice(v string) {
	o.Price = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ProjectUsage) GetTotal() string {
	if o == nil || o.Total == nil {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsage) GetTotalOk() (*string, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ProjectUsage) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *ProjectUsage) SetTotal(v string) {
	o.Total = &v
}

func (o ProjectUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Facility != nil {
		toSerialize["facility"] = o.Facility
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Plan != nil {
		toSerialize["plan"] = o.Plan
	}
	if o.PlanVersion != nil {
		toSerialize["plan_version"] = o.PlanVersion
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableProjectUsage struct {
	value *ProjectUsage
	isSet bool
}

func (v NullableProjectUsage) Get() *ProjectUsage {
	return v.value
}

func (v *NullableProjectUsage) Set(val *ProjectUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectUsage(val *ProjectUsage) *NullableProjectUsage {
	return &NullableProjectUsage{value: val, isSet: true}
}

func (v NullableProjectUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

