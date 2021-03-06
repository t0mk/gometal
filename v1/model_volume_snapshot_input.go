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

// VolumeSnapshotInput struct for VolumeSnapshotInput
type VolumeSnapshotInput struct {
	DeviceId string `json:"device_id"`
}

// NewVolumeSnapshotInput instantiates a new VolumeSnapshotInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeSnapshotInput(deviceId string) *VolumeSnapshotInput {
	this := VolumeSnapshotInput{}
	this.DeviceId = deviceId
	return &this
}

// NewVolumeSnapshotInputWithDefaults instantiates a new VolumeSnapshotInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeSnapshotInputWithDefaults() *VolumeSnapshotInput {
	this := VolumeSnapshotInput{}
	return &this
}

// GetDeviceId returns the DeviceId field value
func (o *VolumeSnapshotInput) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *VolumeSnapshotInput) GetDeviceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *VolumeSnapshotInput) SetDeviceId(v string) {
	o.DeviceId = v
}

func (o VolumeSnapshotInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["device_id"] = o.DeviceId
	}
	return json.Marshal(toSerialize)
}

type NullableVolumeSnapshotInput struct {
	value *VolumeSnapshotInput
	isSet bool
}

func (v NullableVolumeSnapshotInput) Get() *VolumeSnapshotInput {
	return v.value
}

func (v *NullableVolumeSnapshotInput) Set(val *VolumeSnapshotInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeSnapshotInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeSnapshotInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeSnapshotInput(val *VolumeSnapshotInput) *NullableVolumeSnapshotInput {
	return &NullableVolumeSnapshotInput{value: val, isSet: true}
}

func (v NullableVolumeSnapshotInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeSnapshotInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


