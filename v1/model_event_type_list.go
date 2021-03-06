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

// EventTypeList struct for EventTypeList
type EventTypeList struct {
	EventTypes *[]EventType `json:"event_types,omitempty"`
}

// NewEventTypeList instantiates a new EventTypeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTypeList() *EventTypeList {
	this := EventTypeList{}
	return &this
}

// NewEventTypeListWithDefaults instantiates a new EventTypeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypeListWithDefaults() *EventTypeList {
	this := EventTypeList{}
	return &this
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *EventTypeList) GetEventTypes() []EventType {
	if o == nil || o.EventTypes == nil {
		var ret []EventType
		return ret
	}
	return *o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeList) GetEventTypesOk() (*[]EventType, bool) {
	if o == nil || o.EventTypes == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *EventTypeList) HasEventTypes() bool {
	if o != nil && o.EventTypes != nil {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []EventType and assigns it to the EventTypes field.
func (o *EventTypeList) SetEventTypes(v []EventType) {
	o.EventTypes = &v
}

func (o EventTypeList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventTypes != nil {
		toSerialize["event_types"] = o.EventTypes
	}
	return json.Marshal(toSerialize)
}

type NullableEventTypeList struct {
	value *EventTypeList
	isSet bool
}

func (v NullableEventTypeList) Get() *EventTypeList {
	return v.value
}

func (v *NullableEventTypeList) Set(val *EventTypeList) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeList) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeList(val *EventTypeList) *NullableEventTypeList {
	return &NullableEventTypeList{value: val, isSet: true}
}

func (v NullableEventTypeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


