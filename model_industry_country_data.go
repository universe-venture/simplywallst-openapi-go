/*
 * simply-wallst
 *
 * simply-wallst API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// IndustryCountryData struct for IndustryCountryData
type IndustryCountryData struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewIndustryCountryData instantiates a new IndustryCountryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndustryCountryData() *IndustryCountryData {
	this := IndustryCountryData{}
	return &this
}

// NewIndustryCountryDataWithDefaults instantiates a new IndustryCountryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndustryCountryDataWithDefaults() *IndustryCountryData {
	this := IndustryCountryData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IndustryCountryData) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndustryCountryData) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IndustryCountryData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *IndustryCountryData) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IndustryCountryData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndustryCountryData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IndustryCountryData) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IndustryCountryData) SetName(v string) {
	o.Name = &v
}

func (o IndustryCountryData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableIndustryCountryData struct {
	value *IndustryCountryData
	isSet bool
}

func (v NullableIndustryCountryData) Get() *IndustryCountryData {
	return v.value
}

func (v *NullableIndustryCountryData) Set(val *IndustryCountryData) {
	v.value = val
	v.isSet = true
}

func (v NullableIndustryCountryData) IsSet() bool {
	return v.isSet
}

func (v *NullableIndustryCountryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndustryCountryData(val *IndustryCountryData) *NullableIndustryCountryData {
	return &NullableIndustryCountryData{value: val, isSet: true}
}

func (v NullableIndustryCountryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndustryCountryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
