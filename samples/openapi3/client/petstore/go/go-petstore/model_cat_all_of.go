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

// checks if the CatAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatAllOf{}

// CatAllOf struct for CatAllOf
type CatAllOf struct {
	Declawed *bool `json:"declawed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CatAllOf CatAllOf

// NewCatAllOf instantiates a new CatAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatAllOf() *CatAllOf {
	this := CatAllOf{}
	return &this
}

// NewCatAllOfWithDefaults instantiates a new CatAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatAllOfWithDefaults() *CatAllOf {
	this := CatAllOf{}
	return &this
}

// GetDeclawed returns the Declawed field value if set, zero value otherwise.
func (o *CatAllOf) GetDeclawed() bool {
	if o == nil || o.Declawed == nil {
		var ret bool
		return ret
	}
	return *o.Declawed
}

// GetDeclawedOk returns a tuple with the Declawed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatAllOf) GetDeclawedOk() (*bool, bool) {
	if o == nil || o.Declawed == nil {
		return nil, false
	}
	return o.Declawed, true
}

// HasDeclawed returns a boolean if a field has been set.
func (o *CatAllOf) HasDeclawed() bool {
	if o != nil && o.Declawed != nil {
		return true
	}

	return false
}

// SetDeclawed gets a reference to the given bool and assigns it to the Declawed field.
func (o *CatAllOf) SetDeclawed(v bool) {
	o.Declawed = &v
}

func (o CatAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Declawed != nil {
		toSerialize["declawed"] = o.Declawed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCatAllOf := _CatAllOf{}

	if err = json.Unmarshal(bytes, &varCatAllOf); err == nil {
		*o = CatAllOf(varCatAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "declawed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCatAllOf struct {
	value *CatAllOf
	isSet bool
}

func (v NullableCatAllOf) Get() *CatAllOf {
	return v.value
}

func (v *NullableCatAllOf) Set(val *CatAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCatAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCatAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatAllOf(val *CatAllOf) *NullableCatAllOf {
	return &NullableCatAllOf{value: val, isSet: true}
}

func (v NullableCatAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


