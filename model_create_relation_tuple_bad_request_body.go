/*
 * Ory Oathkeeper API
 *
 * Documentation for all of Ory Oathkeeper's APIs. 
 *
 * API version: v0.8.0-alpha.1
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateRelationTupleBadRequestBody CreateRelationTupleBadRequestBody CreateRelationTupleBadRequestBody CreateRelationTupleBadRequestBody CreateRelationTupleBadRequestBody CreateRelationTupleBadRequestBody CreateRelationTupleBadRequestBody create relation tuple bad request body
type CreateRelationTupleBadRequestBody struct {
	// code
	Code *int64 `json:"code,omitempty"`
	// details
	Details []map[string]interface{} `json:"details,omitempty"`
	// message
	Message *string `json:"message,omitempty"`
	// reason
	Reason *string `json:"reason,omitempty"`
	// request
	Request *string `json:"request,omitempty"`
	// status
	Status *string `json:"status,omitempty"`
}

// NewCreateRelationTupleBadRequestBody instantiates a new CreateRelationTupleBadRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRelationTupleBadRequestBody() *CreateRelationTupleBadRequestBody {
	this := CreateRelationTupleBadRequestBody{}
	return &this
}

// NewCreateRelationTupleBadRequestBodyWithDefaults instantiates a new CreateRelationTupleBadRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRelationTupleBadRequestBodyWithDefaults() *CreateRelationTupleBadRequestBody {
	this := CreateRelationTupleBadRequestBody{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateRelationTupleBadRequestBody) GetCode() int64 {
	if o == nil || o.Code == nil {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRelationTupleBadRequestBody) GetCodeOk() (*int64, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateRelationTupleBadRequestBody) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *CreateRelationTupleBadRequestBody) SetCode(v int64) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CreateRelationTupleBadRequestBody) GetDetails() []map[string]interface{} {
	if o == nil || o.Details == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRelationTupleBadRequestBody) GetDetailsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CreateRelationTupleBadRequestBody) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []map[string]interface{} and assigns it to the Details field.
func (o *CreateRelationTupleBadRequestBody) SetDetails(v []map[string]interface{}) {
	o.Details = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateRelationTupleBadRequestBody) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRelationTupleBadRequestBody) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateRelationTupleBadRequestBody) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateRelationTupleBadRequestBody) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CreateRelationTupleBadRequestBody) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRelationTupleBadRequestBody) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CreateRelationTupleBadRequestBody) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CreateRelationTupleBadRequestBody) SetReason(v string) {
	o.Reason = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *CreateRelationTupleBadRequestBody) GetRequest() string {
	if o == nil || o.Request == nil {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRelationTupleBadRequestBody) GetRequestOk() (*string, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *CreateRelationTupleBadRequestBody) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *CreateRelationTupleBadRequestBody) SetRequest(v string) {
	o.Request = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateRelationTupleBadRequestBody) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRelationTupleBadRequestBody) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateRelationTupleBadRequestBody) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateRelationTupleBadRequestBody) SetStatus(v string) {
	o.Status = &v
}

func (o CreateRelationTupleBadRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRelationTupleBadRequestBody struct {
	value *CreateRelationTupleBadRequestBody
	isSet bool
}

func (v NullableCreateRelationTupleBadRequestBody) Get() *CreateRelationTupleBadRequestBody {
	return v.value
}

func (v *NullableCreateRelationTupleBadRequestBody) Set(val *CreateRelationTupleBadRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRelationTupleBadRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRelationTupleBadRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRelationTupleBadRequestBody(val *CreateRelationTupleBadRequestBody) *NullableCreateRelationTupleBadRequestBody {
	return &NullableCreateRelationTupleBadRequestBody{value: val, isSet: true}
}

func (v NullableCreateRelationTupleBadRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRelationTupleBadRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


