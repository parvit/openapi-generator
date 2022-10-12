/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
	"time"
)

// checks if the MixedPropertiesAndAdditionalPropertiesClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MixedPropertiesAndAdditionalPropertiesClass{}

// MixedPropertiesAndAdditionalPropertiesClass struct for MixedPropertiesAndAdditionalPropertiesClass
type MixedPropertiesAndAdditionalPropertiesClass struct {
	Uuid *string `json:"uuid,omitempty"`
	DateTime *time.Time `json:"dateTime,omitempty"`
	Map *map[string]Animal `json:"map,omitempty"`
}

// NewMixedPropertiesAndAdditionalPropertiesClass instantiates a new MixedPropertiesAndAdditionalPropertiesClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMixedPropertiesAndAdditionalPropertiesClass() *MixedPropertiesAndAdditionalPropertiesClass {
	this := MixedPropertiesAndAdditionalPropertiesClass{}
	return &this
}

// NewMixedPropertiesAndAdditionalPropertiesClassWithDefaults instantiates a new MixedPropertiesAndAdditionalPropertiesClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMixedPropertiesAndAdditionalPropertiesClassWithDefaults() *MixedPropertiesAndAdditionalPropertiesClass {
	this := MixedPropertiesAndAdditionalPropertiesClass{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *MixedPropertiesAndAdditionalPropertiesClass) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixedPropertiesAndAdditionalPropertiesClass) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *MixedPropertiesAndAdditionalPropertiesClass) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *MixedPropertiesAndAdditionalPropertiesClass) SetUuid(v string) {
	o.Uuid = &v
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *MixedPropertiesAndAdditionalPropertiesClass) GetDateTime() time.Time {
	if o == nil || o.DateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixedPropertiesAndAdditionalPropertiesClass) GetDateTimeOk() (*time.Time, bool) {
	if o == nil || o.DateTime == nil {
		return nil, false
	}
	return o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *MixedPropertiesAndAdditionalPropertiesClass) HasDateTime() bool {
	if o != nil && o.DateTime != nil {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given time.Time and assigns it to the DateTime field.
func (o *MixedPropertiesAndAdditionalPropertiesClass) SetDateTime(v time.Time) {
	o.DateTime = &v
}

// GetMap returns the Map field value if set, zero value otherwise.
func (o *MixedPropertiesAndAdditionalPropertiesClass) GetMap() map[string]Animal {
	if o == nil || o.Map == nil {
		var ret map[string]Animal
		return ret
	}
	return *o.Map
}

// GetMapOk returns a tuple with the Map field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixedPropertiesAndAdditionalPropertiesClass) GetMapOk() (*map[string]Animal, bool) {
	if o == nil || o.Map == nil {
		return nil, false
	}
	return o.Map, true
}

// HasMap returns a boolean if a field has been set.
func (o *MixedPropertiesAndAdditionalPropertiesClass) HasMap() bool {
	if o != nil && o.Map != nil {
		return true
	}

	return false
}

// SetMap gets a reference to the given map[string]Animal and assigns it to the Map field.
func (o *MixedPropertiesAndAdditionalPropertiesClass) SetMap(v map[string]Animal) {
	o.Map = &v
}

func (o MixedPropertiesAndAdditionalPropertiesClass) MarshalJSON() ([]byte, error) {
	toSerialize := o.ToMap()
	return json.Marshal(toSerialize)
}

func (o MixedPropertiesAndAdditionalPropertiesClass) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = *o.Uuid
	}
	if o.DateTime != nil {
		toSerialize["dateTime"] = *o.DateTime
	}
	if o.Map != nil {
		toSerialize["map"] = *o.Map
	}
	return toSerialize
}

type NullableMixedPropertiesAndAdditionalPropertiesClass struct {
	value *MixedPropertiesAndAdditionalPropertiesClass
	isSet bool
}

func (v NullableMixedPropertiesAndAdditionalPropertiesClass) Get() *MixedPropertiesAndAdditionalPropertiesClass {
	return v.value
}

func (v *NullableMixedPropertiesAndAdditionalPropertiesClass) Set(val *MixedPropertiesAndAdditionalPropertiesClass) {
	v.value = val
	v.isSet = true
}

func (v NullableMixedPropertiesAndAdditionalPropertiesClass) IsSet() bool {
	return v.isSet
}

func (v *NullableMixedPropertiesAndAdditionalPropertiesClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMixedPropertiesAndAdditionalPropertiesClass(val *MixedPropertiesAndAdditionalPropertiesClass) *NullableMixedPropertiesAndAdditionalPropertiesClass {
	return &NullableMixedPropertiesAndAdditionalPropertiesClass{value: val, isSet: true}
}

func (v NullableMixedPropertiesAndAdditionalPropertiesClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMixedPropertiesAndAdditionalPropertiesClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


