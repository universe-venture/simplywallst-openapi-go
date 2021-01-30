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

// CompanyDataScore struct for CompanyDataScore
type CompanyDataScore struct {
	Value      *int32  `json:"value,omitempty"`
	Income     *int32  `json:"income,omitempty"`
	Health     *int32  `json:"health,omitempty"`
	Past       *int32  `json:"past,omitempty"`
	Future     *int32  `json:"future,omitempty"`
	Management *int32  `json:"management,omitempty"`
	Misc       *int32  `json:"misc,omitempty"`
	Total      *int32  `json:"total,omitempty"`
	Sentence   *string `json:"sentence,omitempty"`
}

// NewCompanyDataScore instantiates a new CompanyDataScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyDataScore() *CompanyDataScore {
	this := CompanyDataScore{}
	return &this
}

// NewCompanyDataScoreWithDefaults instantiates a new CompanyDataScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyDataScoreWithDefaults() *CompanyDataScore {
	this := CompanyDataScore{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CompanyDataScore) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataScore) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CompanyDataScore) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *CompanyDataScore) SetValue(v int32) {
	o.Value = &v
}

// GetIncome returns the Income field value if set, zero value otherwise.
func (o *CompanyDataScore) GetIncome() int32 {
	if o == nil || o.Income == nil {
		var ret int32
		return ret
	}
	return *o.Income
}

// GetIncomeOk returns a tuple with the Income field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataScore) GetIncomeOk() (*int32, bool) {
	if o == nil || o.Income == nil {
		return nil, false
	}
	return o.Income, true
}

// HasIncome returns a boolean if a field has been set.
func (o *CompanyDataScore) HasIncome() bool {
	if o != nil && o.Income != nil {
		return true
	}

	return false
}

// SetIncome gets a reference to the given int32 and assigns it to the Income field.
func (o *CompanyDataScore) SetIncome(v int32) {
	o.Income = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *CompanyDataScore) GetHealth() int32 {
	if o == nil || o.Health == nil {
		var ret int32
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataScore) GetHealthOk() (*int32, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *CompanyDataScore) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given int32 and assigns it to the Health field.
func (o *CompanyDataScore) SetHealth(v int32) {
	o.Health = &v
}

// GetPast returns the Past field value if set, zero value otherwise.
func (o *CompanyDataScore) GetPast() int32 {
	if o == nil || o.Past == nil {
		var ret int32
		return ret
	}
	return *o.Past
}

// GetPastOk returns a tuple with the Past field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataScore) GetPastOk() (*int32, bool) {
	if o == nil || o.Past == nil {
		return nil, false
	}
	return o.Past, true
}

// HasPast returns a boolean if a field has been set.
func (o *CompanyDataScore) HasPast() bool {
	if o != nil && o.Past != nil {
		return true
	}

	return false
}

// SetPast gets a reference to the given int32 and assigns it to the Past field.
func (o *CompanyDataScore) SetPast(v int32) {
	o.Past = &v
}

// GetFuture returns the Future field value if set, zero value otherwise.
func (o *CompanyDataScore) GetFuture() int32 {
	if o == nil || o.Future == nil {
		var ret int32
		return ret
	}
	return *o.Future
}

// GetFutureOk returns a tuple with the Future field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataScore) GetFutureOk() (*int32, bool) {
	if o == nil || o.Future == nil {
		return nil, false
	}
	return o.Future, true
}

// HasFuture returns a boolean if a field has been set.
func (o *CompanyDataScore) HasFuture() bool {
	if o != nil && o.Future != nil {
		return true
	}

	return false
}

// SetFuture gets a reference to the given int32 and assigns it to the Future field.
func (o *CompanyDataScore) SetFuture(v int32) {
	o.Future = &v
}

// GetManagement returns the Management field value if set, zero value otherwise.
func (o *CompanyDataScore) GetManagement() int32 {
	if o == nil || o.Management == nil {
		var ret int32
		return ret
	}
	return *o.Management
}

// GetManagementOk returns a tuple with the Management field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataScore) GetManagementOk() (*int32, bool) {
	if o == nil || o.Management == nil {
		return nil, false
	}
	return o.Management, true
}

// HasManagement returns a boolean if a field has been set.
func (o *CompanyDataScore) HasManagement() bool {
	if o != nil && o.Management != nil {
		return true
	}

	return false
}

// SetManagement gets a reference to the given int32 and assigns it to the Management field.
func (o *CompanyDataScore) SetManagement(v int32) {
	o.Management = &v
}

// GetMisc returns the Misc field value if set, zero value otherwise.
func (o *CompanyDataScore) GetMisc() int32 {
	if o == nil || o.Misc == nil {
		var ret int32
		return ret
	}
	return *o.Misc
}

// GetMiscOk returns a tuple with the Misc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataScore) GetMiscOk() (*int32, bool) {
	if o == nil || o.Misc == nil {
		return nil, false
	}
	return o.Misc, true
}

// HasMisc returns a boolean if a field has been set.
func (o *CompanyDataScore) HasMisc() bool {
	if o != nil && o.Misc != nil {
		return true
	}

	return false
}

// SetMisc gets a reference to the given int32 and assigns it to the Misc field.
func (o *CompanyDataScore) SetMisc(v int32) {
	o.Misc = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CompanyDataScore) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataScore) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *CompanyDataScore) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *CompanyDataScore) SetTotal(v int32) {
	o.Total = &v
}

// GetSentence returns the Sentence field value if set, zero value otherwise.
func (o *CompanyDataScore) GetSentence() string {
	if o == nil || o.Sentence == nil {
		var ret string
		return ret
	}
	return *o.Sentence
}

// GetSentenceOk returns a tuple with the Sentence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataScore) GetSentenceOk() (*string, bool) {
	if o == nil || o.Sentence == nil {
		return nil, false
	}
	return o.Sentence, true
}

// HasSentence returns a boolean if a field has been set.
func (o *CompanyDataScore) HasSentence() bool {
	if o != nil && o.Sentence != nil {
		return true
	}

	return false
}

// SetSentence gets a reference to the given string and assigns it to the Sentence field.
func (o *CompanyDataScore) SetSentence(v string) {
	o.Sentence = &v
}

func (o CompanyDataScore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Income != nil {
		toSerialize["income"] = o.Income
	}
	if o.Health != nil {
		toSerialize["health"] = o.Health
	}
	if o.Past != nil {
		toSerialize["past"] = o.Past
	}
	if o.Future != nil {
		toSerialize["future"] = o.Future
	}
	if o.Management != nil {
		toSerialize["management"] = o.Management
	}
	if o.Misc != nil {
		toSerialize["misc"] = o.Misc
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Sentence != nil {
		toSerialize["sentence"] = o.Sentence
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyDataScore struct {
	value *CompanyDataScore
	isSet bool
}

func (v NullableCompanyDataScore) Get() *CompanyDataScore {
	return v.value
}

func (v *NullableCompanyDataScore) Set(val *CompanyDataScore) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyDataScore) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyDataScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyDataScore(val *CompanyDataScore) *NullableCompanyDataScore {
	return &NullableCompanyDataScore{value: val, isSet: true}
}

func (v NullableCompanyDataScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyDataScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
