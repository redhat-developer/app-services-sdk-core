/*
 * Apicurio Registry API [v2]
 *
 * Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 
 *
 * API version: 2.4.x
 * Contact: apicurio@lists.jboss.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
	"time"
)

// Comment struct for Comment
type Comment struct {
	Value string `json:"value"`
	CreatedOn time.Time `json:"createdOn"`
	CreatedBy string `json:"createdBy"`
	CommentId string `json:"commentId"`
}

// NewComment instantiates a new Comment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComment(value string, createdOn time.Time, createdBy string, commentId string) *Comment {
	this := Comment{}
	this.Value = value
	this.CreatedOn = createdOn
	this.CreatedBy = createdBy
	this.CommentId = commentId
	return &this
}

// NewCommentWithDefaults instantiates a new Comment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentWithDefaults() *Comment {
	this := Comment{}
	return &this
}

// GetValue returns the Value field value
func (o *Comment) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Comment) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Comment) SetValue(v string) {
	o.Value = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *Comment) GetCreatedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *Comment) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *Comment) SetCreatedOn(v time.Time) {
	o.CreatedOn = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Comment) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Comment) GetCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Comment) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCommentId returns the CommentId field value
func (o *Comment) GetCommentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommentId
}

// GetCommentIdOk returns a tuple with the CommentId field value
// and a boolean to check if the value has been set.
func (o *Comment) GetCommentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommentId, true
}

// SetCommentId sets field value
func (o *Comment) SetCommentId(v string) {
	o.CommentId = v
}

func (o Comment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["commentId"] = o.CommentId
	}
	return json.Marshal(toSerialize)
}

type NullableComment struct {
	value *Comment
	isSet bool
}

func (v NullableComment) Get() *Comment {
	return v.value
}

func (v *NullableComment) Set(val *Comment) {
	v.value = val
	v.isSet = true
}

func (v NullableComment) IsSet() bool {
	return v.isSet
}

func (v *NullableComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComment(val *Comment) *NullableComment {
	return &NullableComment{value: val, isSet: true}
}

func (v NullableComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

