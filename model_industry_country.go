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

// IndustryCountry struct for IndustryCountry
type IndustryCountry struct {
	Data      *IndustryCountryData      `json:"data,omitempty"`
	Links     *IndustryCountryLinks     `json:"links,omitempty"`
	Countries *IndustryCountryCountries `json:"countries,omitempty"`
}

// NewIndustryCountry instantiates a new IndustryCountry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndustryCountry() *IndustryCountry {
	this := IndustryCountry{}
	return &this
}

// NewIndustryCountryWithDefaults instantiates a new IndustryCountry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndustryCountryWithDefaults() *IndustryCountry {
	this := IndustryCountry{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IndustryCountry) GetData() IndustryCountryData {
	if o == nil || o.Data == nil {
		var ret IndustryCountryData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndustryCountry) GetDataOk() (*IndustryCountryData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IndustryCountry) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given IndustryCountryData and assigns it to the Data field.
func (o *IndustryCountry) SetData(v IndustryCountryData) {
	o.Data = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IndustryCountry) GetLinks() IndustryCountryLinks {
	if o == nil || o.Links == nil {
		var ret IndustryCountryLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndustryCountry) GetLinksOk() (*IndustryCountryLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IndustryCountry) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given IndustryCountryLinks and assigns it to the Links field.
func (o *IndustryCountry) SetLinks(v IndustryCountryLinks) {
	o.Links = &v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *IndustryCountry) GetCountries() IndustryCountryCountries {
	if o == nil || o.Countries == nil {
		var ret IndustryCountryCountries
		return ret
	}
	return *o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndustryCountry) GetCountriesOk() (*IndustryCountryCountries, bool) {
	if o == nil || o.Countries == nil {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *IndustryCountry) HasCountries() bool {
	if o != nil && o.Countries != nil {
		return true
	}

	return false
}

// SetCountries gets a reference to the given IndustryCountryCountries and assigns it to the Countries field.
func (o *IndustryCountry) SetCountries(v IndustryCountryCountries) {
	o.Countries = &v
}

func (o IndustryCountry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Countries != nil {
		toSerialize["countries"] = o.Countries
	}
	return json.Marshal(toSerialize)
}

type NullableIndustryCountry struct {
	value *IndustryCountry
	isSet bool
}

func (v NullableIndustryCountry) Get() *IndustryCountry {
	return v.value
}

func (v *NullableIndustryCountry) Set(val *IndustryCountry) {
	v.value = val
	v.isSet = true
}

func (v NullableIndustryCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableIndustryCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndustryCountry(val *IndustryCountry) *NullableIndustryCountry {
	return &NullableIndustryCountry{value: val, isSet: true}
}

func (v NullableIndustryCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndustryCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
