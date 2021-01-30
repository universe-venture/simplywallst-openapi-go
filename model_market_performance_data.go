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

// MarketPerformanceData struct for MarketPerformanceData
type MarketPerformanceData struct {
	Top   *map[string]MarketPerformanceDate `json:"top,omitempty"`
	Worst *map[string]MarketPerformanceDate `json:"worst,omitempty"`
}

// NewMarketPerformanceData instantiates a new MarketPerformanceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketPerformanceData() *MarketPerformanceData {
	this := MarketPerformanceData{}
	return &this
}

// NewMarketPerformanceDataWithDefaults instantiates a new MarketPerformanceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketPerformanceDataWithDefaults() *MarketPerformanceData {
	this := MarketPerformanceData{}
	return &this
}

// GetTop returns the Top field value if set, zero value otherwise.
func (o *MarketPerformanceData) GetTop() map[string]MarketPerformanceDate {
	if o == nil || o.Top == nil {
		var ret map[string]MarketPerformanceDate
		return ret
	}
	return *o.Top
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketPerformanceData) GetTopOk() (*map[string]MarketPerformanceDate, bool) {
	if o == nil || o.Top == nil {
		return nil, false
	}
	return o.Top, true
}

// HasTop returns a boolean if a field has been set.
func (o *MarketPerformanceData) HasTop() bool {
	if o != nil && o.Top != nil {
		return true
	}

	return false
}

// SetTop gets a reference to the given map[string]MarketPerformanceDate and assigns it to the Top field.
func (o *MarketPerformanceData) SetTop(v map[string]MarketPerformanceDate) {
	o.Top = &v
}

// GetWorst returns the Worst field value if set, zero value otherwise.
func (o *MarketPerformanceData) GetWorst() map[string]MarketPerformanceDate {
	if o == nil || o.Worst == nil {
		var ret map[string]MarketPerformanceDate
		return ret
	}
	return *o.Worst
}

// GetWorstOk returns a tuple with the Worst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketPerformanceData) GetWorstOk() (*map[string]MarketPerformanceDate, bool) {
	if o == nil || o.Worst == nil {
		return nil, false
	}
	return o.Worst, true
}

// HasWorst returns a boolean if a field has been set.
func (o *MarketPerformanceData) HasWorst() bool {
	if o != nil && o.Worst != nil {
		return true
	}

	return false
}

// SetWorst gets a reference to the given map[string]MarketPerformanceDate and assigns it to the Worst field.
func (o *MarketPerformanceData) SetWorst(v map[string]MarketPerformanceDate) {
	o.Worst = &v
}

func (o MarketPerformanceData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Top != nil {
		toSerialize["top"] = o.Top
	}
	if o.Worst != nil {
		toSerialize["worst"] = o.Worst
	}
	return json.Marshal(toSerialize)
}

type NullableMarketPerformanceData struct {
	value *MarketPerformanceData
	isSet bool
}

func (v NullableMarketPerformanceData) Get() *MarketPerformanceData {
	return v.value
}

func (v *NullableMarketPerformanceData) Set(val *MarketPerformanceData) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketPerformanceData) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketPerformanceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketPerformanceData(val *MarketPerformanceData) *NullableMarketPerformanceData {
	return &NullableMarketPerformanceData{value: val, isSet: true}
}

func (v NullableMarketPerformanceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketPerformanceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
