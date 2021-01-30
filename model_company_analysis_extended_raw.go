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

// CompanyAnalysisExtendedRaw struct for CompanyAnalysisExtendedRaw
type CompanyAnalysisExtendedRaw struct {
	Data *CompanyAnalysisExtendedRawData `json:"data,omitempty"`
}

// NewCompanyAnalysisExtendedRaw instantiates a new CompanyAnalysisExtendedRaw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyAnalysisExtendedRaw() *CompanyAnalysisExtendedRaw {
	this := CompanyAnalysisExtendedRaw{}
	return &this
}

// NewCompanyAnalysisExtendedRawWithDefaults instantiates a new CompanyAnalysisExtendedRaw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyAnalysisExtendedRawWithDefaults() *CompanyAnalysisExtendedRaw {
	this := CompanyAnalysisExtendedRaw{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRaw) GetData() CompanyAnalysisExtendedRawData {
	if o == nil || o.Data == nil {
		var ret CompanyAnalysisExtendedRawData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRaw) GetDataOk() (*CompanyAnalysisExtendedRawData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRaw) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CompanyAnalysisExtendedRawData and assigns it to the Data field.
func (o *CompanyAnalysisExtendedRaw) SetData(v CompanyAnalysisExtendedRawData) {
	o.Data = &v
}

func (o CompanyAnalysisExtendedRaw) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyAnalysisExtendedRaw struct {
	value *CompanyAnalysisExtendedRaw
	isSet bool
}

func (v NullableCompanyAnalysisExtendedRaw) Get() *CompanyAnalysisExtendedRaw {
	return v.value
}

func (v *NullableCompanyAnalysisExtendedRaw) Set(val *CompanyAnalysisExtendedRaw) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyAnalysisExtendedRaw) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyAnalysisExtendedRaw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyAnalysisExtendedRaw(val *CompanyAnalysisExtendedRaw) *NullableCompanyAnalysisExtendedRaw {
	return &NullableCompanyAnalysisExtendedRaw{value: val, isSet: true}
}

func (v NullableCompanyAnalysisExtendedRaw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyAnalysisExtendedRaw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
