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

// CompanyAnalysisExtendedRawDataMembersData struct for CompanyAnalysisExtendedRawDataMembersData
type CompanyAnalysisExtendedRawDataMembersData struct {
	Id                         *int32   `json:"id,omitempty"`
	ProId                      *int32   `json:"pro_id,omitempty"`
	Salutation                 *string  `json:"salutation,omitempty"`
	Prefix                     *string  `json:"prefix,omitempty"`
	Suffix                     *string  `json:"suffix,omitempty"`
	FirstName                  *string  `json:"first_name,omitempty"`
	MiddleName                 *string  `json:"middle_name,omitempty"`
	LastName                   *string  `json:"last_name,omitempty"`
	YearBorn                   *int32   `json:"year_born,omitempty"`
	Title                      *string  `json:"title,omitempty"`
	ProTitle                   *string  `json:"pro_title,omitempty"`
	Rank                       *int32   `json:"rank,omitempty"`
	ProRank                    *int32   `json:"pro_rank,omitempty"`
	BoardRank                  *int32   `json:"board_rank,omitempty"`
	Biography                  *string  `json:"biography,omitempty"`
	HoldingDate                *float32 `json:"holding_date,omitempty"`
	SharesHeld                 *float32 `json:"shares_held,omitempty"`
	PercentOfSharesOutstanding *float32 `json:"percent_of_shares_outstanding,omitempty"`
	OptionsHeld                *int32   `json:"options_held,omitempty"`
	SharesChanged              *float32 `json:"shares_changed,omitempty"`
	PercentSharesChanged       *float32 `json:"percent_shares_changed,omitempty"`
	RankSharesHeld             *int32   `json:"rank_shares_held,omitempty"`
	RankSharesBought           *int32   `json:"rank_shares_bought,omitempty"`
	RankSharesSold             *int32   `json:"rank_shares_sold,omitempty"`
	Age                        *int32   `json:"age,omitempty"`
	ShortTitle                 *string  `json:"short_title,omitempty"`
	Tenure                     *float32 `json:"tenure,omitempty"`
	StartDate                  *float32 `json:"start_date,omitempty"`
}

// NewCompanyAnalysisExtendedRawDataMembersData instantiates a new CompanyAnalysisExtendedRawDataMembersData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyAnalysisExtendedRawDataMembersData() *CompanyAnalysisExtendedRawDataMembersData {
	this := CompanyAnalysisExtendedRawDataMembersData{}
	return &this
}

// NewCompanyAnalysisExtendedRawDataMembersDataWithDefaults instantiates a new CompanyAnalysisExtendedRawDataMembersData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyAnalysisExtendedRawDataMembersDataWithDefaults() *CompanyAnalysisExtendedRawDataMembersData {
	this := CompanyAnalysisExtendedRawDataMembersData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetId(v int32) {
	o.Id = &v
}

// GetProId returns the ProId field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetProId() int32 {
	if o == nil || o.ProId == nil {
		var ret int32
		return ret
	}
	return *o.ProId
}

// GetProIdOk returns a tuple with the ProId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetProIdOk() (*int32, bool) {
	if o == nil || o.ProId == nil {
		return nil, false
	}
	return o.ProId, true
}

// HasProId returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasProId() bool {
	if o != nil && o.ProId != nil {
		return true
	}

	return false
}

// SetProId gets a reference to the given int32 and assigns it to the ProId field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetProId(v int32) {
	o.ProId = &v
}

// GetSalutation returns the Salutation field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetSalutation() string {
	if o == nil || o.Salutation == nil {
		var ret string
		return ret
	}
	return *o.Salutation
}

// GetSalutationOk returns a tuple with the Salutation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetSalutationOk() (*string, bool) {
	if o == nil || o.Salutation == nil {
		return nil, false
	}
	return o.Salutation, true
}

// HasSalutation returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasSalutation() bool {
	if o != nil && o.Salutation != nil {
		return true
	}

	return false
}

// SetSalutation gets a reference to the given string and assigns it to the Salutation field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetSalutation(v string) {
	o.Salutation = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetPrefix(v string) {
	o.Prefix = &v
}

// GetSuffix returns the Suffix field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetSuffix() string {
	if o == nil || o.Suffix == nil {
		var ret string
		return ret
	}
	return *o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetSuffixOk() (*string, bool) {
	if o == nil || o.Suffix == nil {
		return nil, false
	}
	return o.Suffix, true
}

// HasSuffix returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasSuffix() bool {
	if o != nil && o.Suffix != nil {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given string and assigns it to the Suffix field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetSuffix(v string) {
	o.Suffix = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetFirstName(v string) {
	o.FirstName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetMiddleName() string {
	if o == nil || o.MiddleName == nil {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetMiddleNameOk() (*string, bool) {
	if o == nil || o.MiddleName == nil {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasMiddleName() bool {
	if o != nil && o.MiddleName != nil {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetLastName(v string) {
	o.LastName = &v
}

// GetYearBorn returns the YearBorn field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetYearBorn() int32 {
	if o == nil || o.YearBorn == nil {
		var ret int32
		return ret
	}
	return *o.YearBorn
}

// GetYearBornOk returns a tuple with the YearBorn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetYearBornOk() (*int32, bool) {
	if o == nil || o.YearBorn == nil {
		return nil, false
	}
	return o.YearBorn, true
}

// HasYearBorn returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasYearBorn() bool {
	if o != nil && o.YearBorn != nil {
		return true
	}

	return false
}

// SetYearBorn gets a reference to the given int32 and assigns it to the YearBorn field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetYearBorn(v int32) {
	o.YearBorn = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetTitle(v string) {
	o.Title = &v
}

// GetProTitle returns the ProTitle field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetProTitle() string {
	if o == nil || o.ProTitle == nil {
		var ret string
		return ret
	}
	return *o.ProTitle
}

// GetProTitleOk returns a tuple with the ProTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetProTitleOk() (*string, bool) {
	if o == nil || o.ProTitle == nil {
		return nil, false
	}
	return o.ProTitle, true
}

// HasProTitle returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasProTitle() bool {
	if o != nil && o.ProTitle != nil {
		return true
	}

	return false
}

// SetProTitle gets a reference to the given string and assigns it to the ProTitle field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetProTitle(v string) {
	o.ProTitle = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetRank() int32 {
	if o == nil || o.Rank == nil {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetRankOk() (*int32, bool) {
	if o == nil || o.Rank == nil {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasRank() bool {
	if o != nil && o.Rank != nil {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetRank(v int32) {
	o.Rank = &v
}

// GetProRank returns the ProRank field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetProRank() int32 {
	if o == nil || o.ProRank == nil {
		var ret int32
		return ret
	}
	return *o.ProRank
}

// GetProRankOk returns a tuple with the ProRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetProRankOk() (*int32, bool) {
	if o == nil || o.ProRank == nil {
		return nil, false
	}
	return o.ProRank, true
}

// HasProRank returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasProRank() bool {
	if o != nil && o.ProRank != nil {
		return true
	}

	return false
}

// SetProRank gets a reference to the given int32 and assigns it to the ProRank field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetProRank(v int32) {
	o.ProRank = &v
}

// GetBoardRank returns the BoardRank field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetBoardRank() int32 {
	if o == nil || o.BoardRank == nil {
		var ret int32
		return ret
	}
	return *o.BoardRank
}

// GetBoardRankOk returns a tuple with the BoardRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetBoardRankOk() (*int32, bool) {
	if o == nil || o.BoardRank == nil {
		return nil, false
	}
	return o.BoardRank, true
}

// HasBoardRank returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasBoardRank() bool {
	if o != nil && o.BoardRank != nil {
		return true
	}

	return false
}

// SetBoardRank gets a reference to the given int32 and assigns it to the BoardRank field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetBoardRank(v int32) {
	o.BoardRank = &v
}

// GetBiography returns the Biography field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetBiography() string {
	if o == nil || o.Biography == nil {
		var ret string
		return ret
	}
	return *o.Biography
}

// GetBiographyOk returns a tuple with the Biography field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetBiographyOk() (*string, bool) {
	if o == nil || o.Biography == nil {
		return nil, false
	}
	return o.Biography, true
}

// HasBiography returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasBiography() bool {
	if o != nil && o.Biography != nil {
		return true
	}

	return false
}

// SetBiography gets a reference to the given string and assigns it to the Biography field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetBiography(v string) {
	o.Biography = &v
}

// GetHoldingDate returns the HoldingDate field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetHoldingDate() float32 {
	if o == nil || o.HoldingDate == nil {
		var ret float32
		return ret
	}
	return *o.HoldingDate
}

// GetHoldingDateOk returns a tuple with the HoldingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetHoldingDateOk() (*float32, bool) {
	if o == nil || o.HoldingDate == nil {
		return nil, false
	}
	return o.HoldingDate, true
}

// HasHoldingDate returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasHoldingDate() bool {
	if o != nil && o.HoldingDate != nil {
		return true
	}

	return false
}

// SetHoldingDate gets a reference to the given float32 and assigns it to the HoldingDate field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetHoldingDate(v float32) {
	o.HoldingDate = &v
}

// GetSharesHeld returns the SharesHeld field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetSharesHeld() float32 {
	if o == nil || o.SharesHeld == nil {
		var ret float32
		return ret
	}
	return *o.SharesHeld
}

// GetSharesHeldOk returns a tuple with the SharesHeld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetSharesHeldOk() (*float32, bool) {
	if o == nil || o.SharesHeld == nil {
		return nil, false
	}
	return o.SharesHeld, true
}

// HasSharesHeld returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasSharesHeld() bool {
	if o != nil && o.SharesHeld != nil {
		return true
	}

	return false
}

// SetSharesHeld gets a reference to the given float32 and assigns it to the SharesHeld field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetSharesHeld(v float32) {
	o.SharesHeld = &v
}

// GetPercentOfSharesOutstanding returns the PercentOfSharesOutstanding field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetPercentOfSharesOutstanding() float32 {
	if o == nil || o.PercentOfSharesOutstanding == nil {
		var ret float32
		return ret
	}
	return *o.PercentOfSharesOutstanding
}

// GetPercentOfSharesOutstandingOk returns a tuple with the PercentOfSharesOutstanding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetPercentOfSharesOutstandingOk() (*float32, bool) {
	if o == nil || o.PercentOfSharesOutstanding == nil {
		return nil, false
	}
	return o.PercentOfSharesOutstanding, true
}

// HasPercentOfSharesOutstanding returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasPercentOfSharesOutstanding() bool {
	if o != nil && o.PercentOfSharesOutstanding != nil {
		return true
	}

	return false
}

// SetPercentOfSharesOutstanding gets a reference to the given float32 and assigns it to the PercentOfSharesOutstanding field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetPercentOfSharesOutstanding(v float32) {
	o.PercentOfSharesOutstanding = &v
}

// GetOptionsHeld returns the OptionsHeld field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetOptionsHeld() int32 {
	if o == nil || o.OptionsHeld == nil {
		var ret int32
		return ret
	}
	return *o.OptionsHeld
}

// GetOptionsHeldOk returns a tuple with the OptionsHeld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetOptionsHeldOk() (*int32, bool) {
	if o == nil || o.OptionsHeld == nil {
		return nil, false
	}
	return o.OptionsHeld, true
}

// HasOptionsHeld returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasOptionsHeld() bool {
	if o != nil && o.OptionsHeld != nil {
		return true
	}

	return false
}

// SetOptionsHeld gets a reference to the given int32 and assigns it to the OptionsHeld field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetOptionsHeld(v int32) {
	o.OptionsHeld = &v
}

// GetSharesChanged returns the SharesChanged field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetSharesChanged() float32 {
	if o == nil || o.SharesChanged == nil {
		var ret float32
		return ret
	}
	return *o.SharesChanged
}

// GetSharesChangedOk returns a tuple with the SharesChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetSharesChangedOk() (*float32, bool) {
	if o == nil || o.SharesChanged == nil {
		return nil, false
	}
	return o.SharesChanged, true
}

// HasSharesChanged returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasSharesChanged() bool {
	if o != nil && o.SharesChanged != nil {
		return true
	}

	return false
}

// SetSharesChanged gets a reference to the given float32 and assigns it to the SharesChanged field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetSharesChanged(v float32) {
	o.SharesChanged = &v
}

// GetPercentSharesChanged returns the PercentSharesChanged field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetPercentSharesChanged() float32 {
	if o == nil || o.PercentSharesChanged == nil {
		var ret float32
		return ret
	}
	return *o.PercentSharesChanged
}

// GetPercentSharesChangedOk returns a tuple with the PercentSharesChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetPercentSharesChangedOk() (*float32, bool) {
	if o == nil || o.PercentSharesChanged == nil {
		return nil, false
	}
	return o.PercentSharesChanged, true
}

// HasPercentSharesChanged returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasPercentSharesChanged() bool {
	if o != nil && o.PercentSharesChanged != nil {
		return true
	}

	return false
}

// SetPercentSharesChanged gets a reference to the given float32 and assigns it to the PercentSharesChanged field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetPercentSharesChanged(v float32) {
	o.PercentSharesChanged = &v
}

// GetRankSharesHeld returns the RankSharesHeld field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetRankSharesHeld() int32 {
	if o == nil || o.RankSharesHeld == nil {
		var ret int32
		return ret
	}
	return *o.RankSharesHeld
}

// GetRankSharesHeldOk returns a tuple with the RankSharesHeld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetRankSharesHeldOk() (*int32, bool) {
	if o == nil || o.RankSharesHeld == nil {
		return nil, false
	}
	return o.RankSharesHeld, true
}

// HasRankSharesHeld returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasRankSharesHeld() bool {
	if o != nil && o.RankSharesHeld != nil {
		return true
	}

	return false
}

// SetRankSharesHeld gets a reference to the given int32 and assigns it to the RankSharesHeld field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetRankSharesHeld(v int32) {
	o.RankSharesHeld = &v
}

// GetRankSharesBought returns the RankSharesBought field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetRankSharesBought() int32 {
	if o == nil || o.RankSharesBought == nil {
		var ret int32
		return ret
	}
	return *o.RankSharesBought
}

// GetRankSharesBoughtOk returns a tuple with the RankSharesBought field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetRankSharesBoughtOk() (*int32, bool) {
	if o == nil || o.RankSharesBought == nil {
		return nil, false
	}
	return o.RankSharesBought, true
}

// HasRankSharesBought returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasRankSharesBought() bool {
	if o != nil && o.RankSharesBought != nil {
		return true
	}

	return false
}

// SetRankSharesBought gets a reference to the given int32 and assigns it to the RankSharesBought field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetRankSharesBought(v int32) {
	o.RankSharesBought = &v
}

// GetRankSharesSold returns the RankSharesSold field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetRankSharesSold() int32 {
	if o == nil || o.RankSharesSold == nil {
		var ret int32
		return ret
	}
	return *o.RankSharesSold
}

// GetRankSharesSoldOk returns a tuple with the RankSharesSold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetRankSharesSoldOk() (*int32, bool) {
	if o == nil || o.RankSharesSold == nil {
		return nil, false
	}
	return o.RankSharesSold, true
}

// HasRankSharesSold returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasRankSharesSold() bool {
	if o != nil && o.RankSharesSold != nil {
		return true
	}

	return false
}

// SetRankSharesSold gets a reference to the given int32 and assigns it to the RankSharesSold field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetRankSharesSold(v int32) {
	o.RankSharesSold = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetAge() int32 {
	if o == nil || o.Age == nil {
		var ret int32
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetAgeOk() (*int32, bool) {
	if o == nil || o.Age == nil {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasAge() bool {
	if o != nil && o.Age != nil {
		return true
	}

	return false
}

// SetAge gets a reference to the given int32 and assigns it to the Age field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetAge(v int32) {
	o.Age = &v
}

// GetShortTitle returns the ShortTitle field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetShortTitle() string {
	if o == nil || o.ShortTitle == nil {
		var ret string
		return ret
	}
	return *o.ShortTitle
}

// GetShortTitleOk returns a tuple with the ShortTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetShortTitleOk() (*string, bool) {
	if o == nil || o.ShortTitle == nil {
		return nil, false
	}
	return o.ShortTitle, true
}

// HasShortTitle returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasShortTitle() bool {
	if o != nil && o.ShortTitle != nil {
		return true
	}

	return false
}

// SetShortTitle gets a reference to the given string and assigns it to the ShortTitle field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetShortTitle(v string) {
	o.ShortTitle = &v
}

// GetTenure returns the Tenure field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetTenure() float32 {
	if o == nil || o.Tenure == nil {
		var ret float32
		return ret
	}
	return *o.Tenure
}

// GetTenureOk returns a tuple with the Tenure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetTenureOk() (*float32, bool) {
	if o == nil || o.Tenure == nil {
		return nil, false
	}
	return o.Tenure, true
}

// HasTenure returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasTenure() bool {
	if o != nil && o.Tenure != nil {
		return true
	}

	return false
}

// SetTenure gets a reference to the given float32 and assigns it to the Tenure field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetTenure(v float32) {
	o.Tenure = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetStartDate() float32 {
	if o == nil || o.StartDate == nil {
		var ret float32
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) GetStartDateOk() (*float32, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataMembersData) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given float32 and assigns it to the StartDate field.
func (o *CompanyAnalysisExtendedRawDataMembersData) SetStartDate(v float32) {
	o.StartDate = &v
}

func (o CompanyAnalysisExtendedRawDataMembersData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ProId != nil {
		toSerialize["pro_id"] = o.ProId
	}
	if o.Salutation != nil {
		toSerialize["salutation"] = o.Salutation
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	if o.Suffix != nil {
		toSerialize["suffix"] = o.Suffix
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.MiddleName != nil {
		toSerialize["middle_name"] = o.MiddleName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.YearBorn != nil {
		toSerialize["year_born"] = o.YearBorn
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.ProTitle != nil {
		toSerialize["pro_title"] = o.ProTitle
	}
	if o.Rank != nil {
		toSerialize["rank"] = o.Rank
	}
	if o.ProRank != nil {
		toSerialize["pro_rank"] = o.ProRank
	}
	if o.BoardRank != nil {
		toSerialize["board_rank"] = o.BoardRank
	}
	if o.Biography != nil {
		toSerialize["biography"] = o.Biography
	}
	if o.HoldingDate != nil {
		toSerialize["holding_date"] = o.HoldingDate
	}
	if o.SharesHeld != nil {
		toSerialize["shares_held"] = o.SharesHeld
	}
	if o.PercentOfSharesOutstanding != nil {
		toSerialize["percent_of_shares_outstanding"] = o.PercentOfSharesOutstanding
	}
	if o.OptionsHeld != nil {
		toSerialize["options_held"] = o.OptionsHeld
	}
	if o.SharesChanged != nil {
		toSerialize["shares_changed"] = o.SharesChanged
	}
	if o.PercentSharesChanged != nil {
		toSerialize["percent_shares_changed"] = o.PercentSharesChanged
	}
	if o.RankSharesHeld != nil {
		toSerialize["rank_shares_held"] = o.RankSharesHeld
	}
	if o.RankSharesBought != nil {
		toSerialize["rank_shares_bought"] = o.RankSharesBought
	}
	if o.RankSharesSold != nil {
		toSerialize["rank_shares_sold"] = o.RankSharesSold
	}
	if o.Age != nil {
		toSerialize["age"] = o.Age
	}
	if o.ShortTitle != nil {
		toSerialize["short_title"] = o.ShortTitle
	}
	if o.Tenure != nil {
		toSerialize["tenure"] = o.Tenure
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyAnalysisExtendedRawDataMembersData struct {
	value *CompanyAnalysisExtendedRawDataMembersData
	isSet bool
}

func (v NullableCompanyAnalysisExtendedRawDataMembersData) Get() *CompanyAnalysisExtendedRawDataMembersData {
	return v.value
}

func (v *NullableCompanyAnalysisExtendedRawDataMembersData) Set(val *CompanyAnalysisExtendedRawDataMembersData) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyAnalysisExtendedRawDataMembersData) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyAnalysisExtendedRawDataMembersData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyAnalysisExtendedRawDataMembersData(val *CompanyAnalysisExtendedRawDataMembersData) *NullableCompanyAnalysisExtendedRawDataMembersData {
	return &NullableCompanyAnalysisExtendedRawDataMembersData{value: val, isSet: true}
}

func (v NullableCompanyAnalysisExtendedRawDataMembersData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyAnalysisExtendedRawDataMembersData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
