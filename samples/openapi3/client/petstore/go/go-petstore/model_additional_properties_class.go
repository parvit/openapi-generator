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

// checks if the AdditionalPropertiesClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalPropertiesClass{}

// AdditionalPropertiesClass struct for AdditionalPropertiesClass
type AdditionalPropertiesClass struct {
	MapProperty *map[string]string `json:"map_property,omitempty"`
	MapOfMapProperty *map[string]map[string]string `json:"map_of_map_property,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdditionalPropertiesClass AdditionalPropertiesClass

// NewAdditionalPropertiesClass instantiates a new AdditionalPropertiesClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalPropertiesClass() *AdditionalPropertiesClass {
	this := AdditionalPropertiesClass{}
	return &this
}

// NewAdditionalPropertiesClassWithDefaults instantiates a new AdditionalPropertiesClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalPropertiesClassWithDefaults() *AdditionalPropertiesClass {
	this := AdditionalPropertiesClass{}
	return &this
}

// GetMapProperty returns the MapProperty field value if set, zero value otherwise.
func (o *AdditionalPropertiesClass) GetMapProperty() map[string]string {
	if o == nil || o.MapProperty == nil {
		var ret map[string]string
		return ret
	}
	return *o.MapProperty
}

// GetMapPropertyOk returns a tuple with the MapProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalPropertiesClass) GetMapPropertyOk() (*map[string]string, bool) {
	if o == nil || o.MapProperty == nil {
		return nil, false
	}
	return o.MapProperty, true
}

// HasMapProperty returns a boolean if a field has been set.
func (o *AdditionalPropertiesClass) HasMapProperty() bool {
	if o != nil && o.MapProperty != nil {
		return true
	}

	return false
}

// SetMapProperty gets a reference to the given map[string]string and assigns it to the MapProperty field.
func (o *AdditionalPropertiesClass) SetMapProperty(v map[string]string) {
	o.MapProperty = &v
}

// GetMapOfMapProperty returns the MapOfMapProperty field value if set, zero value otherwise.
func (o *AdditionalPropertiesClass) GetMapOfMapProperty() map[string]map[string]string {
	if o == nil || o.MapOfMapProperty == nil {
		var ret map[string]map[string]string
		return ret
	}
	return *o.MapOfMapProperty
}

// GetMapOfMapPropertyOk returns a tuple with the MapOfMapProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalPropertiesClass) GetMapOfMapPropertyOk() (*map[string]map[string]string, bool) {
	if o == nil || o.MapOfMapProperty == nil {
		return nil, false
	}
	return o.MapOfMapProperty, true
}

// HasMapOfMapProperty returns a boolean if a field has been set.
func (o *AdditionalPropertiesClass) HasMapOfMapProperty() bool {
	if o != nil && o.MapOfMapProperty != nil {
		return true
	}

	return false
}

// SetMapOfMapProperty gets a reference to the given map[string]map[string]string and assigns it to the MapOfMapProperty field.
func (o *AdditionalPropertiesClass) SetMapOfMapProperty(v map[string]map[string]string) {
	o.MapOfMapProperty = &v
}

func (o AdditionalPropertiesClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalPropertiesClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MapProperty != nil {
		toSerialize["map_property"] = o.MapProperty
	}
	if o.MapOfMapProperty != nil {
		toSerialize["map_of_map_property"] = o.MapOfMapProperty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdditionalPropertiesClass) UnmarshalJSON(bytes []byte) (err error) {
	varAdditionalPropertiesClass := _AdditionalPropertiesClass{}

	if err = json.Unmarshal(bytes, &varAdditionalPropertiesClass); err == nil {
		*o = AdditionalPropertiesClass(varAdditionalPropertiesClass)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "map_property")
		delete(additionalProperties, "map_of_map_property")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdditionalPropertiesClass struct {
	value *AdditionalPropertiesClass
	isSet bool
}

func (v NullableAdditionalPropertiesClass) Get() *AdditionalPropertiesClass {
	return v.value
}

func (v *NullableAdditionalPropertiesClass) Set(val *AdditionalPropertiesClass) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalPropertiesClass) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalPropertiesClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalPropertiesClass(val *AdditionalPropertiesClass) *NullableAdditionalPropertiesClass {
	return &NullableAdditionalPropertiesClass{value: val, isSet: true}
}

func (v NullableAdditionalPropertiesClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalPropertiesClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


