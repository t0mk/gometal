/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// BgpNeighborDataRoutesOut struct for BgpNeighborDataRoutesOut
type BgpNeighborDataRoutesOut struct {
	Route *string `json:"route,omitempty"`
	Exact *bool `json:"exact,omitempty"`
}

// NewBgpNeighborDataRoutesOut instantiates a new BgpNeighborDataRoutesOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpNeighborDataRoutesOut() *BgpNeighborDataRoutesOut {
	this := BgpNeighborDataRoutesOut{}
	return &this
}

// NewBgpNeighborDataRoutesOutWithDefaults instantiates a new BgpNeighborDataRoutesOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpNeighborDataRoutesOutWithDefaults() *BgpNeighborDataRoutesOut {
	this := BgpNeighborDataRoutesOut{}
	return &this
}

// GetRoute returns the Route field value if set, zero value otherwise.
func (o *BgpNeighborDataRoutesOut) GetRoute() string {
	if o == nil || o.Route == nil {
		var ret string
		return ret
	}
	return *o.Route
}

// GetRouteOk returns a tuple with the Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborDataRoutesOut) GetRouteOk() (*string, bool) {
	if o == nil || o.Route == nil {
		return nil, false
	}
	return o.Route, true
}

// HasRoute returns a boolean if a field has been set.
func (o *BgpNeighborDataRoutesOut) HasRoute() bool {
	if o != nil && o.Route != nil {
		return true
	}

	return false
}

// SetRoute gets a reference to the given string and assigns it to the Route field.
func (o *BgpNeighborDataRoutesOut) SetRoute(v string) {
	o.Route = &v
}

// GetExact returns the Exact field value if set, zero value otherwise.
func (o *BgpNeighborDataRoutesOut) GetExact() bool {
	if o == nil || o.Exact == nil {
		var ret bool
		return ret
	}
	return *o.Exact
}

// GetExactOk returns a tuple with the Exact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborDataRoutesOut) GetExactOk() (*bool, bool) {
	if o == nil || o.Exact == nil {
		return nil, false
	}
	return o.Exact, true
}

// HasExact returns a boolean if a field has been set.
func (o *BgpNeighborDataRoutesOut) HasExact() bool {
	if o != nil && o.Exact != nil {
		return true
	}

	return false
}

// SetExact gets a reference to the given bool and assigns it to the Exact field.
func (o *BgpNeighborDataRoutesOut) SetExact(v bool) {
	o.Exact = &v
}

func (o BgpNeighborDataRoutesOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Route != nil {
		toSerialize["route"] = o.Route
	}
	if o.Exact != nil {
		toSerialize["exact"] = o.Exact
	}
	return json.Marshal(toSerialize)
}

type NullableBgpNeighborDataRoutesOut struct {
	value *BgpNeighborDataRoutesOut
	isSet bool
}

func (v NullableBgpNeighborDataRoutesOut) Get() *BgpNeighborDataRoutesOut {
	return v.value
}

func (v *NullableBgpNeighborDataRoutesOut) Set(val *BgpNeighborDataRoutesOut) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpNeighborDataRoutesOut) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpNeighborDataRoutesOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpNeighborDataRoutesOut(val *BgpNeighborDataRoutesOut) *NullableBgpNeighborDataRoutesOut {
	return &NullableBgpNeighborDataRoutesOut{value: val, isSet: true}
}

func (v NullableBgpNeighborDataRoutesOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpNeighborDataRoutesOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


