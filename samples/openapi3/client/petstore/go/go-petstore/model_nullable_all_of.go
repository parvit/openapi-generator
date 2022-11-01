/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
)

// checks if the NullableAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NullableAllOf{}

// NullableAllOf struct for NullableAllOf
type NullableAllOf struct {
	Child NullableNullableAllOfChild `json:"child,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NullableAllOf NullableAllOf

// NewNullableAllOf instantiates a new NullableAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNullableAllOf() *NullableAllOf {
	this := NullableAllOf{}
	return &this
}

// NewNullableAllOfWithDefaults instantiates a new NullableAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNullableAllOfWithDefaults() *NullableAllOf {
	this := NullableAllOf{}
	return &this
}

// GetChild returns the Child field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullableAllOf) GetChild() NullableAllOfChild {
	if o == nil || isNil(o.Child.Get()) {
		var ret NullableAllOfChild
		return ret
	}
	return *o.Child.Get()
}

// GetChildOk returns a tuple with the Child field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableAllOf) GetChildOk() (*NullableAllOfChild, bool) {
	if o == nil {
    return nil, false
	}
	return o.Child.Get(), o.Child.IsSet()
}

// HasChild returns a boolean if a field has been set.
func (o *NullableAllOf) HasChild() bool {
	if o != nil && o.Child.IsSet() {
		return true
	}

	return false
}

// SetChild gets a reference to the given NullableNullableAllOfChild and assigns it to the Child field.
func (o *NullableAllOf) SetChild(v NullableAllOfChild) {
	o.Child.Set(&v)
}
// SetChildNil sets the value for Child to be an explicit nil
func (o *NullableAllOf) SetChildNil() {
	o.Child.Set(nil)
}

// UnsetChild ensures that no value is present for Child, not even an explicit nil
func (o *NullableAllOf) UnsetChild() {
	o.Child.Unset()
}

func (o NullableAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NullableAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Child.IsSet() {
		toSerialize["child"] = o.Child.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NullableAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNullableAllOf := _NullableAllOf{}

	if err = json.Unmarshal(bytes, &varNullableAllOf); err == nil {
		*o = NullableAllOf(varNullableAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "child")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNullableAllOf struct {
	value *NullableAllOf
	isSet bool
}

func (v NullableNullableAllOf) Get() *NullableAllOf {
	return v.value
}

func (v *NullableNullableAllOf) Set(val *NullableAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNullableAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNullableAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullableAllOf(val *NullableAllOf) *NullableNullableAllOf {
	return &NullableNullableAllOf{value: val, isSet: true}
}

func (v NullableNullableAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullableAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


