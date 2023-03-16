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
 * SinkConnectorResponse
 */
@JsonPropertyOrder({
  SinkConnectorResponse.JSON_PROPERTY_KIND,
  SinkConnectorResponse.JSON_PROPERTY_ID,
  SinkConnectorResponse.JSON_PROPERTY_NAME,
  SinkConnectorResponse.JSON_PROPERTY_HREF,
  SinkConnectorResponse.JSON_PROPERTY_SUBMITTED_AT,
  SinkConnectorResponse.JSON_PROPERTY_PUBLISHED_AT,
  SinkConnectorResponse.JSON_PROPERTY_MODIFIED_AT,
  SinkConnectorResponse.JSON_PROPERTY_STATUS,
  SinkConnectorResponse.JSON_PROPERTY_OWNER,
  SinkConnectorResponse.JSON_PROPERTY_CONNECTOR_TYPE_ID,
  SinkConnectorResponse.JSON_PROPERTY_CONNECTOR,
  SinkConnectorResponse.JSON_PROPERTY_STATUS_MESSAGE,
  SinkConnectorResponse.JSON_PROPERTY_URI_DSL
})
@JsonTypeName("SinkConnectorResponse")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class SinkConnectorResponse {
  public static final String JSON_PROPERTY_KIND = "kind";
  private String kind;

  public static final String JSON_PROPERTY_ID = "id";
  private String id;

  public static final String JSON_PROPERTY_NAME = "name";
  private String name;

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

  public static final String JSON_PROPERTY_CONNECTOR_TYPE_ID = "connector_type_id";
  private String connectorTypeId;

  public static final String JSON_PROPERTY_CONNECTOR = "connector";
  private Object connector;

  public static final String JSON_PROPERTY_STATUS_MESSAGE = "status_message";
  private String statusMessage;

  public static final String JSON_PROPERTY_URI_DSL = "uri_dsl";
  private String uriDsl;

  public SinkConnectorResponse() { 
  }

  public SinkConnectorResponse kind(String kind) {
    
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


  public SinkConnectorResponse id(String id) {
    
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


  public SinkConnectorResponse name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * The name of this resource
   * @return name
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(example = "resourceName1", required = true, value = "The name of this resource")
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


  public SinkConnectorResponse href(String href) {
    
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


  public SinkConnectorResponse submittedAt(OffsetDateTime submittedAt) {
    
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


  public SinkConnectorResponse publishedAt(OffsetDateTime publishedAt) {
    
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


  public SinkConnectorResponse modifiedAt(OffsetDateTime modifiedAt) {
    
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


  public SinkConnectorResponse status(ManagedResourceStatus status) {
    
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


  public SinkConnectorResponse owner(String owner) {
    
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


  public SinkConnectorResponse connectorTypeId(String connectorTypeId) {
    
    this.connectorTypeId = connectorTypeId;
    return this;
  }

   /**
   * The connector type
   * @return connectorTypeId
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(example = "slack_sink_0.1", required = true, value = "The connector type")
  @JsonProperty(JSON_PROPERTY_CONNECTOR_TYPE_ID)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getConnectorTypeId() {
    return connectorTypeId;
  }


  @JsonProperty(JSON_PROPERTY_CONNECTOR_TYPE_ID)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setConnectorTypeId(String connectorTypeId) {
    this.connectorTypeId = connectorTypeId;
  }


  public SinkConnectorResponse connector(Object connector) {
    
    this.connector = connector;
    return this;
  }

   /**
   * The Connector configuration payload
   * @return connector
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The Connector configuration payload")
  @JsonProperty(JSON_PROPERTY_CONNECTOR)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public Object getConnector() {
    return connector;
  }


  @JsonProperty(JSON_PROPERTY_CONNECTOR)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setConnector(Object connector) {
    this.connector = connector;
  }


  public SinkConnectorResponse statusMessage(String statusMessage) {
    
    this.statusMessage = statusMessage;
    return this;
  }

   /**
   * A detailed status message in case there is a problem with the connector
   * @return statusMessage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "A detailed status message in case there is a problem with the connector")
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


  public SinkConnectorResponse uriDsl(String uriDsl) {
    
    this.uriDsl = uriDsl;
    return this;
  }

   /**
   * The URI to be used in Camel DSL to send data to this sink
   * @return uriDsl
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(example = "knative:my-id", required = true, value = "The URI to be used in Camel DSL to send data to this sink")
  @JsonProperty(JSON_PROPERTY_URI_DSL)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getUriDsl() {
    return uriDsl;
  }


  @JsonProperty(JSON_PROPERTY_URI_DSL)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setUriDsl(String uriDsl) {
    this.uriDsl = uriDsl;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SinkConnectorResponse sinkConnectorResponse = (SinkConnectorResponse) o;
    return Objects.equals(this.kind, sinkConnectorResponse.kind) &&
        Objects.equals(this.id, sinkConnectorResponse.id) &&
        Objects.equals(this.name, sinkConnectorResponse.name) &&
        Objects.equals(this.href, sinkConnectorResponse.href) &&
        Objects.equals(this.submittedAt, sinkConnectorResponse.submittedAt) &&
        Objects.equals(this.publishedAt, sinkConnectorResponse.publishedAt) &&
        Objects.equals(this.modifiedAt, sinkConnectorResponse.modifiedAt) &&
        Objects.equals(this.status, sinkConnectorResponse.status) &&
        Objects.equals(this.owner, sinkConnectorResponse.owner) &&
        Objects.equals(this.connectorTypeId, sinkConnectorResponse.connectorTypeId) &&
        Objects.equals(this.connector, sinkConnectorResponse.connector) &&
        Objects.equals(this.statusMessage, sinkConnectorResponse.statusMessage) &&
        Objects.equals(this.uriDsl, sinkConnectorResponse.uriDsl);
  }

  @Override
  public int hashCode() {
    return Objects.hash(kind, id, name, href, submittedAt, publishedAt, modifiedAt, status, owner, connectorTypeId, connector, statusMessage, uriDsl);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SinkConnectorResponse {\n");
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    href: ").append(toIndentedString(href)).append("\n");
    sb.append("    submittedAt: ").append(toIndentedString(submittedAt)).append("\n");
    sb.append("    publishedAt: ").append(toIndentedString(publishedAt)).append("\n");
    sb.append("    modifiedAt: ").append(toIndentedString(modifiedAt)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    owner: ").append(toIndentedString(owner)).append("\n");
    sb.append("    connectorTypeId: ").append(toIndentedString(connectorTypeId)).append("\n");
    sb.append("    connector: ").append(toIndentedString(connector)).append("\n");
    sb.append("    statusMessage: ").append(toIndentedString(statusMessage)).append("\n");
    sb.append("    uriDsl: ").append(toIndentedString(uriDsl)).append("\n");
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

