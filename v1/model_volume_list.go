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

// VolumeList struct for VolumeList
type VolumeList struct {
	Volumes *[]Volume `json:"volumes,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
}

// NewVolumeList instantiates a new VolumeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeList() *VolumeList {
	this := VolumeList{}
	return &this
}

// NewVolumeListWithDefaults instantiates a new VolumeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeListWithDefaults() *VolumeList {
	this := VolumeList{}
	return &this
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *VolumeList) GetVolumes() []Volume {
	if o == nil || o.Volumes == nil {
		var ret []Volume
		return ret
	}
	return *o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeList) GetVolumesOk() (*[]Volume, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *VolumeList) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []Volume and assigns it to the Volumes field.
func (o *VolumeList) SetVolumes(v []Volume) {
	o.Volumes = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *VolumeList) GetMeta() Meta {
	if o == nil || o.Meta == nil {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeList) GetMetaOk() (*Meta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *VolumeList) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *VolumeList) SetMeta(v Meta) {
	o.Meta = &v
}

func (o VolumeList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableVolumeList struct {
	value *VolumeList
	isSet bool
}

func (v NullableVolumeList) Get() *VolumeList {
	return v.value
}

func (v *NullableVolumeList) Set(val *VolumeList) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeList) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeList(val *VolumeList) *NullableVolumeList {
	return &NullableVolumeList{value: val, isSet: true}
}

func (v NullableVolumeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


