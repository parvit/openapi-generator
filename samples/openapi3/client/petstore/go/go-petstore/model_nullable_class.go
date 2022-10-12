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

// checks if the NullableClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NullableClass{}

// NullableClass struct for NullableClass
type NullableClass struct {
	IntegerProp NullableInt32 `json:"integer_prop,omitempty"`
	NumberProp NullableFloat32 `json:"number_prop,omitempty"`
	BooleanProp NullableBool `json:"boolean_prop,omitempty"`
	StringProp NullableString `json:"string_prop,omitempty"`
	DateProp NullableString `json:"date_prop,omitempty"`
	DatetimeProp NullableTime `json:"datetime_prop,omitempty"`
	ArrayNullableProp []map[string]interface{} `json:"array_nullable_prop,omitempty"`
	ArrayAndItemsNullableProp []*map[string]interface{} `json:"array_and_items_nullable_prop,omitempty"`
	ArrayItemsNullable []*map[string]interface{} `json:"array_items_nullable,omitempty"`
	ObjectNullableProp map[string]map[string]interface{} `json:"object_nullable_prop,omitempty"`
	ObjectAndItemsNullableProp map[string]map[string]interface{} `json:"object_and_items_nullable_prop,omitempty"`
	ObjectItemsNullable map[string]map[string]interface{} `json:"object_items_nullable,omitempty"`
}

// NewNullableClass instantiates a new NullableClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNullableClass() *NullableClass {
	this := NullableClass{}
	var booleanProp bool = false
	this.BooleanProp = *NewNullableBool(&booleanProp)
	return &this
}

// NewNullableClassWithDefaults instantiates a new NullableClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNullableClassWithDefaults() *NullableClass {
	this := NullableClass{}
	var booleanProp bool = false
	this.BooleanProp = *NewNullableBool(&booleanProp)
	return &this
}

// GetIntegerProp returns the IntegerProp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullableClass) GetIntegerProp() int32 {
	if o == nil || o.IntegerProp.Get() == nil {
		var ret int32
		return ret
	}
	return *o.IntegerProp.Get()
}

// GetIntegerPropOk returns a tuple with the IntegerProp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableClass) GetIntegerPropOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntegerProp.Get(), o.IntegerProp.IsSet()
}

// HasIntegerProp returns a boolean if a field has been set.
func (o *NullableClass) HasIntegerProp() bool {
	if o != nil && o.IntegerProp.IsSet() {
		return true
	}

	return false
}

// SetIntegerProp gets a reference to the given NullableInt32 and assigns it to the IntegerProp field.
func (o *NullableClass) SetIntegerProp(v int32) {
	o.IntegerProp.Set(&v)
}
// SetIntegerPropNil sets the value for IntegerProp to be an explicit nil
func (o *NullableClass) SetIntegerPropNil() {
	o.IntegerProp.Set(nil)
}

// UnsetIntegerProp ensures that no value is present for IntegerProp, not even an explicit nil
func (o *NullableClass) UnsetIntegerProp() {
	o.IntegerProp.Unset()
}

// GetNumberProp returns the NumberProp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullableClass) GetNumberProp() float32 {
	if o == nil || o.NumberProp.Get() == nil {
		var ret float32
		return ret
	}
	return *o.NumberProp.Get()
}

// GetNumberPropOk returns a tuple with the NumberProp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableClass) GetNumberPropOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumberProp.Get(), o.NumberProp.IsSet()
}

// HasNumberProp returns a boolean if a field has been set.
func (o *NullableClass) HasNumberProp() bool {
	if o != nil && o.NumberProp.IsSet() {
		return true
	}

	return false
}

// SetNumberProp gets a reference to the given NullableFloat32 and assigns it to the NumberProp field.
func (o *NullableClass) SetNumberProp(v float32) {
	o.NumberProp.Set(&v)
}
// SetNumberPropNil sets the value for NumberProp to be an explicit nil
func (o *NullableClass) SetNumberPropNil() {
	o.NumberProp.Set(nil)
}

// UnsetNumberProp ensures that no value is present for NumberProp, not even an explicit nil
func (o *NullableClass) UnsetNumberProp() {
	o.NumberProp.Unset()
}

// GetBooleanProp returns the BooleanProp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullableClass) GetBooleanProp() bool {
	if o == nil || o.BooleanProp.Get() == nil {
		var ret bool
		return ret
	}
	return *o.BooleanProp.Get()
}

// GetBooleanPropOk returns a tuple with the BooleanProp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableClass) GetBooleanPropOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.BooleanProp.Get(), o.BooleanProp.IsSet()
}

// HasBooleanProp returns a boolean if a field has been set.
func (o *NullableClass) HasBooleanProp() bool {
	if o != nil && o.BooleanProp.IsSet() {
		return true
	}

	return false
}

// SetBooleanProp gets a reference to the given NullableBool and assigns it to the BooleanProp field.
func (o *NullableClass) SetBooleanProp(v bool) {
	o.BooleanProp.Set(&v)
}
// SetBooleanPropNil sets the value for BooleanProp to be an explicit nil
func (o *NullableClass) SetBooleanPropNil() {
	o.BooleanProp.Set(nil)
}

// UnsetBooleanProp ensures that no value is present for BooleanProp, not even an explicit nil
func (o *NullableClass) UnsetBooleanProp() {
	o.BooleanProp.Unset()
}

// GetStringProp returns the StringProp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullableClass) GetStringProp() string {
	if o == nil || o.StringProp.Get() == nil {
		var ret string
		return ret
	}
	return *o.StringProp.Get()
}

// GetStringPropOk returns a tuple with the StringProp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableClass) GetStringPropOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StringProp.Get(), o.StringProp.IsSet()
}

// HasStringProp returns a boolean if a field has been set.
func (o *NullableClass) HasStringProp() bool {
	if o != nil && o.StringProp.IsSet() {
		return true
	}

	return false
}

// SetStringProp gets a reference to the given NullableString and assigns it to the StringProp field.
func (o *NullableClass) SetStringProp(v string) {
	o.StringProp.Set(&v)
}
// SetStringPropNil sets the value for StringProp to be an explicit nil
func (o *NullableClass) SetStringPropNil() {
	o.StringProp.Set(nil)
}

// UnsetStringProp ensures that no value is present for StringProp, not even an explicit nil
func (o *NullableClass) UnsetStringProp() {
	o.StringProp.Unset()
}

// GetDateProp returns the DateProp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullableClass) GetDateProp() string {
	if o == nil || o.DateProp.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateProp.Get()
}

// GetDatePropOk returns a tuple with the DateProp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableClass) GetDatePropOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateProp.Get(), o.DateProp.IsSet()
}

// HasDateProp returns a boolean if a field has been set.
func (o *NullableClass) HasDateProp() bool {
	if o != nil && o.DateProp.IsSet() {
		return true
	}

	return false
}

// SetDateProp gets a reference to the given NullableString and assigns it to the DateProp field.
func (o *NullableClass) SetDateProp(v string) {
	o.DateProp.Set(&v)
}
// SetDatePropNil sets the value for DateProp to be an explicit nil
func (o *NullableClass) SetDatePropNil() {
	o.DateProp.Set(nil)
}

// UnsetDateProp ensures that no value is present for DateProp, not even an explicit nil
func (o *NullableClass) UnsetDateProp() {
	o.DateProp.Unset()
}

// GetDatetimeProp returns the DatetimeProp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullableClass) GetDatetimeProp() time.Time {
	if o == nil || o.DatetimeProp.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DatetimeProp.Get()
}

// GetDatetimePropOk returns a tuple with the DatetimeProp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableClass) GetDatetimePropOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatetimeProp.Get(), o.DatetimeProp.IsSet()
}

// HasDatetimeProp returns a boolean if a field has been set.
func (o *NullableClass) HasDatetimeProp() bool {
	if o != nil && o.DatetimeProp.IsSet() {
		return true
	}

	return false
}

// SetDatetimeProp gets a reference to the given NullableTime and assigns it to the DatetimeProp field.
func (o *NullableClass) SetDatetimeProp(v time.Time) {
	o.DatetimeProp.Set(&v)
}
// SetDatetimePropNil sets the value for DatetimeProp to be an explicit nil
func (o *NullableClass) SetDatetimePropNil() {
	o.DatetimeProp.Set(nil)
}

// UnsetDatetimeProp ensures that no value is present for DatetimeProp, not even an explicit nil
func (o *NullableClass) UnsetDatetimeProp() {
	o.DatetimeProp.Unset()
}

// GetArrayNullableProp returns the ArrayNullableProp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullableClass) GetArrayNullableProp() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.ArrayNullableProp
}

// GetArrayNullablePropOk returns a tuple with the ArrayNullableProp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableClass) GetArrayNullablePropOk() ([]map[string]interface{}, bool) {
	if o == nil || o.ArrayNullableProp == nil {
		return nil, false
	}
	return o.ArrayNullableProp, true
}

// HasArrayNullableProp returns a boolean if a field has been set.
func (o *NullableClass) HasArrayNullableProp() bool {
	if o != nil && o.ArrayNullableProp != nil {
		return true
	}

	return false
}

// SetArrayNullableProp gets a reference to the given []map[string]interface{} and assigns it to the ArrayNullableProp field.
func (o *NullableClass) SetArrayNullableProp(v []map[string]interface{}) {
	o.ArrayNullableProp = v
}

// GetArrayAndItemsNullableProp returns the ArrayAndItemsNullableProp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullableClass) GetArrayAndItemsNullableProp() []*map[string]interface{} {
	if o == nil {
		var ret []*map[string]interface{}
		return ret
	}
	return o.ArrayAndItemsNullableProp
}

// GetArrayAndItemsNullablePropOk returns a tuple with the ArrayAndItemsNullableProp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableClass) GetArrayAndItemsNullablePropOk() ([]*map[string]interface{}, bool) {
	if o == nil || o.ArrayAndItemsNullableProp == nil {
		return nil, false
	}
	return o.ArrayAndItemsNullableProp, true
}

// HasArrayAndItemsNullableProp returns a boolean if a field has been set.
func (o *NullableClass) HasArrayAndItemsNullableProp() bool {
	if o != nil && o.ArrayAndItemsNullableProp != nil {
		return true
	}

	return false
}

// SetArrayAndItemsNullableProp gets a reference to the given []*map[string]interface{} and assigns it to the ArrayAndItemsNullableProp field.
func (o *NullableClass) SetArrayAndItemsNullableProp(v []*map[string]interface{}) {
	o.ArrayAndItemsNullableProp = v
}

// GetArrayItemsNullable returns the ArrayItemsNullable field value if set, zero value otherwise.
func (o *NullableClass) GetArrayItemsNullable() []*map[string]interface{} {
	if o == nil || o.ArrayItemsNullable == nil {
		var ret []*map[string]interface{}
		return ret
	}
	return o.ArrayItemsNullable
}

// GetArrayItemsNullableOk returns a tuple with the ArrayItemsNullable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetArrayItemsNullableOk() ([]*map[string]interface{}, bool) {
	if o == nil || o.ArrayItemsNullable == nil {
		return nil, false
	}
	return o.ArrayItemsNullable, true
}

// HasArrayItemsNullable returns a boolean if a field has been set.
func (o *NullableClass) HasArrayItemsNullable() bool {
	if o != nil && o.ArrayItemsNullable != nil {
		return true
	}

	return false
}

// SetArrayItemsNullable gets a reference to the given []*map[string]interface{} and assigns it to the ArrayItemsNullable field.
func (o *NullableClass) SetArrayItemsNullable(v []*map[string]interface{}) {
	o.ArrayItemsNullable = v
}

// GetObjectNullableProp returns the ObjectNullableProp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullableClass) GetObjectNullableProp() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ObjectNullableProp
}

// GetObjectNullablePropOk returns a tuple with the ObjectNullableProp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableClass) GetObjectNullablePropOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.ObjectNullableProp == nil {
		return nil, false
	}
	return o.ObjectNullableProp, true
}

// HasObjectNullableProp returns a boolean if a field has been set.
func (o *NullableClass) HasObjectNullableProp() bool {
	if o != nil && o.ObjectNullableProp != nil {
		return true
	}

	return false
}

// SetObjectNullableProp gets a reference to the given map[string]map[string]interface{} and assigns it to the ObjectNullableProp field.
func (o *NullableClass) SetObjectNullableProp(v map[string]map[string]interface{}) {
	o.ObjectNullableProp = v
}

// GetObjectAndItemsNullableProp returns the ObjectAndItemsNullableProp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullableClass) GetObjectAndItemsNullableProp() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ObjectAndItemsNullableProp
}

// GetObjectAndItemsNullablePropOk returns a tuple with the ObjectAndItemsNullableProp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableClass) GetObjectAndItemsNullablePropOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.ObjectAndItemsNullableProp == nil {
		return nil, false
	}
	return o.ObjectAndItemsNullableProp, true
}

// HasObjectAndItemsNullableProp returns a boolean if a field has been set.
func (o *NullableClass) HasObjectAndItemsNullableProp() bool {
	if o != nil && o.ObjectAndItemsNullableProp != nil {
		return true
	}

	return false
}

// SetObjectAndItemsNullableProp gets a reference to the given map[string]map[string]interface{} and assigns it to the ObjectAndItemsNullableProp field.
func (o *NullableClass) SetObjectAndItemsNullableProp(v map[string]map[string]interface{}) {
	o.ObjectAndItemsNullableProp = v
}

// GetObjectItemsNullable returns the ObjectItemsNullable field value if set, zero value otherwise.
func (o *NullableClass) GetObjectItemsNullable() map[string]map[string]interface{} {
	if o == nil || o.ObjectItemsNullable == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ObjectItemsNullable
}

// GetObjectItemsNullableOk returns a tuple with the ObjectItemsNullable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetObjectItemsNullableOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.ObjectItemsNullable == nil {
		return nil, false
	}
	return o.ObjectItemsNullable, true
}

// HasObjectItemsNullable returns a boolean if a field has been set.
func (o *NullableClass) HasObjectItemsNullable() bool {
	if o != nil && o.ObjectItemsNullable != nil {
		return true
	}

	return false
}

// SetObjectItemsNullable gets a reference to the given map[string]map[string]interface{} and assigns it to the ObjectItemsNullable field.
func (o *NullableClass) SetObjectItemsNullable(v map[string]map[string]interface{}) {
	o.ObjectItemsNullable = v
}

func (o NullableClass) MarshalJSON() ([]byte, error) {
	toSerialize := o.ToMap()
	return json.Marshal(toSerialize)
}

func (o NullableClass) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	if o.IntegerProp.IsSet() {
		toSerialize["integer_prop"] = *o.IntegerProp.Get()
	}
	if o.NumberProp.IsSet() {
		toSerialize["number_prop"] = *o.NumberProp.Get()
	}
	if o.BooleanProp.IsSet() {
		toSerialize["boolean_prop"] = *o.BooleanProp.Get()
	}
	if o.StringProp.IsSet() {
		toSerialize["string_prop"] = *o.StringProp.Get()
	}
	if o.DateProp.IsSet() {
		toSerialize["date_prop"] = *o.DateProp.Get()
	}
	if o.DatetimeProp.IsSet() {
		toSerialize["datetime_prop"] = *o.DatetimeProp.Get()
	}
	if o.ArrayNullableProp != nil {
		toSerialize["array_nullable_prop"] = *o.ArrayNullableProp
	}
	if o.ArrayAndItemsNullableProp != nil {
		toSerialize["array_and_items_nullable_prop"] = *o.ArrayAndItemsNullableProp
	}
	if o.ArrayItemsNullable != nil {
		toSerialize["array_items_nullable"] = *o.ArrayItemsNullable
	}
	if o.ObjectNullableProp != nil {
		toSerialize["object_nullable_prop"] = *o.ObjectNullableProp
	}
	if o.ObjectAndItemsNullableProp != nil {
		toSerialize["object_and_items_nullable_prop"] = *o.ObjectAndItemsNullableProp
	}
	if o.ObjectItemsNullable != nil {
		toSerialize["object_items_nullable"] = *o.ObjectItemsNullable
	}
	return toSerialize
}

type NullableNullableClass struct {
	value *NullableClass
	isSet bool
}

func (v NullableNullableClass) Get() *NullableClass {
	return v.value
}

func (v *NullableNullableClass) Set(val *NullableClass) {
	v.value = val
	v.isSet = true
}

func (v NullableNullableClass) IsSet() bool {
	return v.isSet
}

func (v *NullableNullableClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullableClass(val *NullableClass) *NullableNullableClass {
	return &NullableNullableClass{value: val, isSet: true}
}

func (v NullableNullableClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullableClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


