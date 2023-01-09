/*
 * Red Hat Openshift SmartEvents Fleet Manager V2
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * The version of the OpenAPI document: 0.0.1
 * Contact: openbridge-dev@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.openshift.cloud.api.smartevents.models;

import java.util.Objects;
import java.util.Arrays;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import com.openshift.cloud.api.smartevents.models.ManagedResourceStatus;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.time.OffsetDateTime;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonTypeName;

/**
 * ProcessorResponse
 */
@JsonPropertyOrder({
  ProcessorResponse.JSON_PROPERTY_KIND,
  ProcessorResponse.JSON_PROPERTY_ID,
  ProcessorResponse.JSON_PROPERTY_HREF,
  ProcessorResponse.JSON_PROPERTY_SUBMITTED_AT,
  ProcessorResponse.JSON_PROPERTY_PUBLISHED_AT,
  ProcessorResponse.JSON_PROPERTY_MODIFIED_AT,
  ProcessorResponse.JSON_PROPERTY_STATUS,
  ProcessorResponse.JSON_PROPERTY_OWNER,
  ProcessorResponse.JSON_PROPERTY_NAME,
  ProcessorResponse.JSON_PROPERTY_FLOWS,
  ProcessorResponse.JSON_PROPERTY_STATUS_MESSAGE
})
@JsonTypeName("ProcessorResponse")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class ProcessorResponse {
  public static final String JSON_PROPERTY_KIND = "kind";
  private String kind;

  public static final String JSON_PROPERTY_ID = "id";
  private String id;

  public static final String JSON_PROPERTY_HREF = "href";
  private String href;

  public static final String JSON_PROPERTY_SUBMITTED_AT = "submitted_at";
  private OffsetDateTime submittedAt;

  public static final String JSON_PROPERTY_PUBLISHED_AT = "published_at";
  private OffsetDateTime publishedAt;

  public static final String JSON_PROPERTY_MODIFIED_AT = "modified_at";
  private OffsetDateTime modifiedAt;

  public static final String JSON_PROPERTY_STATUS = "status";
  private ManagedResourceStatus status;

  public static final String JSON_PROPERTY_OWNER = "owner";
  private String owner;

  public static final String JSON_PROPERTY_NAME = "name";
  private String name;

  public static final String JSON_PROPERTY_FLOWS = "flows";
  private Object flows;

  public static final String JSON_PROPERTY_STATUS_MESSAGE = "status_message";
  private String statusMessage;

  public ProcessorResponse() { 
  }

  public ProcessorResponse kind(String kind) {
    
    this.kind = kind;
    return this;
  }

   /**
   * The kind (type) of this resource
   * @return kind
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The kind (type) of this resource")
  @JsonProperty(JSON_PROPERTY_KIND)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getKind() {
    return kind;
  }


  @JsonProperty(JSON_PROPERTY_KIND)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setKind(String kind) {
    this.kind = kind;
  }


  public ProcessorResponse id(String id) {
    
    this.id = id;
    return this;
  }

   /**
   * The unique identifier of this resource
   * @return id
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The unique identifier of this resource")
  @JsonProperty(JSON_PROPERTY_ID)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getId() {
    return id;
  }


  @JsonProperty(JSON_PROPERTY_ID)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setId(String id) {
    this.id = id;
  }


  public ProcessorResponse href(String href) {
    
    this.href = href;
    return this;
  }

   /**
   * The URL of this resource, without the protocol
   * @return href
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(example = "example.com/resource", required = true, value = "The URL of this resource, without the protocol")
  @JsonProperty(JSON_PROPERTY_HREF)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getHref() {
    return href;
  }


  @JsonProperty(JSON_PROPERTY_HREF)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setHref(String href) {
    this.href = href;
  }


  public ProcessorResponse submittedAt(OffsetDateTime submittedAt) {
    
    this.submittedAt = submittedAt;
    return this;
  }

   /**
   * Get submittedAt
   * @return submittedAt
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(example = "2022-03-10T12:15:50-04:00", required = true, value = "")
  @JsonProperty(JSON_PROPERTY_SUBMITTED_AT)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public OffsetDateTime getSubmittedAt() {
    return submittedAt;
  }


  @JsonProperty(JSON_PROPERTY_SUBMITTED_AT)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setSubmittedAt(OffsetDateTime submittedAt) {
    this.submittedAt = submittedAt;
  }


  public ProcessorResponse publishedAt(OffsetDateTime publishedAt) {
    
    this.publishedAt = publishedAt;
    return this;
  }

   /**
   * Get publishedAt
   * @return publishedAt
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "2022-03-10T12:15:50-04:00", value = "")
  @JsonProperty(JSON_PROPERTY_PUBLISHED_AT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public OffsetDateTime getPublishedAt() {
    return publishedAt;
  }


  @JsonProperty(JSON_PROPERTY_PUBLISHED_AT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setPublishedAt(OffsetDateTime publishedAt) {
    this.publishedAt = publishedAt;
  }


  public ProcessorResponse modifiedAt(OffsetDateTime modifiedAt) {
    
    this.modifiedAt = modifiedAt;
    return this;
  }

   /**
   * Get modifiedAt
   * @return modifiedAt
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "2022-03-10T12:15:50-04:00", value = "")
  @JsonProperty(JSON_PROPERTY_MODIFIED_AT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public OffsetDateTime getModifiedAt() {
    return modifiedAt;
  }


  @JsonProperty(JSON_PROPERTY_MODIFIED_AT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setModifiedAt(OffsetDateTime modifiedAt) {
    this.modifiedAt = modifiedAt;
  }


  public ProcessorResponse status(ManagedResourceStatus status) {
    
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")
  @JsonProperty(JSON_PROPERTY_STATUS)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public ManagedResourceStatus getStatus() {
    return status;
  }


  @JsonProperty(JSON_PROPERTY_STATUS)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setStatus(ManagedResourceStatus status) {
    this.status = status;
  }


  public ProcessorResponse owner(String owner) {
    
    this.owner = owner;
    return this;
  }

   /**
   * The user that owns this resource
   * @return owner
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(example = "userName", required = true, value = "The user that owns this resource")
  @JsonProperty(JSON_PROPERTY_OWNER)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getOwner() {
    return owner;
  }


  @JsonProperty(JSON_PROPERTY_OWNER)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setOwner(String owner) {
    this.owner = owner;
  }


  public ProcessorResponse name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * The name of the processor
   * @return name
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(example = "processor1", required = true, value = "The name of the processor")
  @JsonProperty(JSON_PROPERTY_NAME)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getName() {
    return name;
  }


  @JsonProperty(JSON_PROPERTY_NAME)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setName(String name) {
    this.name = name;
  }


  public ProcessorResponse flows(Object flows) {
    
    this.flows = flows;
    return this;
  }

   /**
   * The Camel YAML DSL code, formatted as JSON, that defines the flows in the processor
   * @return flows
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The Camel YAML DSL code, formatted as JSON, that defines the flows in the processor")
  @JsonProperty(JSON_PROPERTY_FLOWS)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public Object getFlows() {
    return flows;
  }


  @JsonProperty(JSON_PROPERTY_FLOWS)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setFlows(Object flows) {
    this.flows = flows;
  }


  public ProcessorResponse statusMessage(String statusMessage) {
    
    this.statusMessage = statusMessage;
    return this;
  }

   /**
   * Get statusMessage
   * @return statusMessage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_STATUS_MESSAGE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getStatusMessage() {
    return statusMessage;
  }


  @JsonProperty(JSON_PROPERTY_STATUS_MESSAGE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setStatusMessage(String statusMessage) {
    this.statusMessage = statusMessage;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ProcessorResponse processorResponse = (ProcessorResponse) o;
    return Objects.equals(this.kind, processorResponse.kind) &&
        Objects.equals(this.id, processorResponse.id) &&
        Objects.equals(this.href, processorResponse.href) &&
        Objects.equals(this.submittedAt, processorResponse.submittedAt) &&
        Objects.equals(this.publishedAt, processorResponse.publishedAt) &&
        Objects.equals(this.modifiedAt, processorResponse.modifiedAt) &&
        Objects.equals(this.status, processorResponse.status) &&
        Objects.equals(this.owner, processorResponse.owner) &&
        Objects.equals(this.name, processorResponse.name) &&
        Objects.equals(this.flows, processorResponse.flows) &&
        Objects.equals(this.statusMessage, processorResponse.statusMessage);
  }

  @Override
  public int hashCode() {
    return Objects.hash(kind, id, href, submittedAt, publishedAt, modifiedAt, status, owner, name, flows, statusMessage);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ProcessorResponse {\n");
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    href: ").append(toIndentedString(href)).append("\n");
    sb.append("    submittedAt: ").append(toIndentedString(submittedAt)).append("\n");
    sb.append("    publishedAt: ").append(toIndentedString(publishedAt)).append("\n");
    sb.append("    modifiedAt: ").append(toIndentedString(modifiedAt)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    owner: ").append(toIndentedString(owner)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    flows: ").append(toIndentedString(flows)).append("\n");
    sb.append("    statusMessage: ").append(toIndentedString(statusMessage)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}
