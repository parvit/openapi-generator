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

// checks if the FileSchemaTestClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSchemaTestClass{}

// FileSchemaTestClass struct for FileSchemaTestClass
type FileSchemaTestClass struct {
	File *File `json:"file,omitempty"`
	Files []File `json:"files,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FileSchemaTestClass FileSchemaTestClass

// NewFileSchemaTestClass instantiates a new FileSchemaTestClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSchemaTestClass() *FileSchemaTestClass {
	this := FileSchemaTestClass{}
	return &this
}

// NewFileSchemaTestClassWithDefaults instantiates a new FileSchemaTestClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSchemaTestClassWithDefaults() *FileSchemaTestClass {
	this := FileSchemaTestClass{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *FileSchemaTestClass) GetFile() File {
	if o == nil || o.File == nil {
		var ret File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSchemaTestClass) GetFileOk() (*File, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *FileSchemaTestClass) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given File and assigns it to the File field.
func (o *FileSchemaTestClass) SetFile(v File) {
	o.File = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *FileSchemaTestClass) GetFiles() []File {
	if o == nil || o.Files == nil {
		var ret []File
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSchemaTestClass) GetFilesOk() ([]File, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *FileSchemaTestClass) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []File and assigns it to the Files field.
func (o *FileSchemaTestClass) SetFiles(v []File) {
	o.Files = v
}

func (o FileSchemaTestClass) MarshalJSON() ([]byte, error) {
	toSerialize := o.ToMap()
	return json.Marshal(toSerialize)
}

func (o FileSchemaTestClass) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	if o.File != nil {
		toSerialize["file"] = *o.File
	}
	if o.Files != nil {
		toSerialize["files"] = *o.Files
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize
}

func (o *FileSchemaTestClass) UnmarshalJSON(bytes []byte) (err error) {
	varFileSchemaTestClass := _FileSchemaTestClass{}

	if err = json.Unmarshal(bytes, &varFileSchemaTestClass); err == nil {
		*o = FileSchemaTestClass(varFileSchemaTestClass)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "file")
		delete(additionalProperties, "files")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileSchemaTestClass struct {
	value *FileSchemaTestClass
	isSet bool
}

func (v NullableFileSchemaTestClass) Get() *FileSchemaTestClass {
	return v.value
}

func (v *NullableFileSchemaTestClass) Set(val *FileSchemaTestClass) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSchemaTestClass) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSchemaTestClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSchemaTestClass(val *FileSchemaTestClass) *NullableFileSchemaTestClass {
	return &NullableFileSchemaTestClass{value: val, isSet: true}
}

func (v NullableFileSchemaTestClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSchemaTestClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


