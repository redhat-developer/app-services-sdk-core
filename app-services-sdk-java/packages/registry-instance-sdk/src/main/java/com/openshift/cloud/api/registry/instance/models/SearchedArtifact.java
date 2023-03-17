/*
 * Apicurio Registry API [v2]
 * Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 
 *
 * The version of the OpenAPI document: 2.4.x
 * Contact: apicurio@lists.jboss.org
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.openshift.cloud.api.registry.instance.models;

import java.util.Objects;
import java.util.Arrays;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import com.openshift.cloud.api.registry.instance.models.ArtifactState;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonTypeName;

/**
 * Models a single artifact from the result set returned when searching for artifacts.
 */
@ApiModel(description = "Models a single artifact from the result set returned when searching for artifacts.")
@JsonPropertyOrder({
  SearchedArtifact.JSON_PROPERTY_ID,
  SearchedArtifact.JSON_PROPERTY_NAME,
  SearchedArtifact.JSON_PROPERTY_DESCRIPTION,
  SearchedArtifact.JSON_PROPERTY_CREATED_ON,
  SearchedArtifact.JSON_PROPERTY_CREATED_BY,
  SearchedArtifact.JSON_PROPERTY_TYPE,
  SearchedArtifact.JSON_PROPERTY_LABELS,
  SearchedArtifact.JSON_PROPERTY_STATE,
  SearchedArtifact.JSON_PROPERTY_MODIFIED_ON,
  SearchedArtifact.JSON_PROPERTY_MODIFIED_BY,
  SearchedArtifact.JSON_PROPERTY_GROUP_ID
})
@JsonTypeName("SearchedArtifact")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class SearchedArtifact {
  public static final String JSON_PROPERTY_ID = "id";
  private String id;

  public static final String JSON_PROPERTY_NAME = "name";
  private String name;

  public static final String JSON_PROPERTY_DESCRIPTION = "description";
  private String description;

  public static final String JSON_PROPERTY_CREATED_ON = "createdOn";
  private OffsetDateTime createdOn;

  public static final String JSON_PROPERTY_CREATED_BY = "createdBy";
  private String createdBy;

  public static final String JSON_PROPERTY_TYPE = "type";
  private String type;

  public static final String JSON_PROPERTY_LABELS = "labels";
  private List<String> labels = null;

  public static final String JSON_PROPERTY_STATE = "state";
  private ArtifactState state;

  public static final String JSON_PROPERTY_MODIFIED_ON = "modifiedOn";
  private OffsetDateTime modifiedOn;

  public static final String JSON_PROPERTY_MODIFIED_BY = "modifiedBy";
  private String modifiedBy;

  public static final String JSON_PROPERTY_GROUP_ID = "groupId";
  private String groupId;

  public SearchedArtifact() { 
  }

  public SearchedArtifact id(String id) {
    
    this.id = id;
    return this;
  }

   /**
   * The ID of a single artifact.
   * @return id
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(example = "example-artifact", required = true, value = "The ID of a single artifact.")
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


  public SearchedArtifact name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_NAME)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getName() {
    return name;
  }


  @JsonProperty(JSON_PROPERTY_NAME)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setName(String name) {
    this.name = name;
  }


  public SearchedArtifact description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_DESCRIPTION)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getDescription() {
    return description;
  }


  @JsonProperty(JSON_PROPERTY_DESCRIPTION)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDescription(String description) {
    this.description = description;
  }


  public SearchedArtifact createdOn(OffsetDateTime createdOn) {
    
    this.createdOn = createdOn;
    return this;
  }

   /**
   * Get createdOn
   * @return createdOn
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")
  @JsonProperty(JSON_PROPERTY_CREATED_ON)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public OffsetDateTime getCreatedOn() {
    return createdOn;
  }


  @JsonProperty(JSON_PROPERTY_CREATED_ON)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setCreatedOn(OffsetDateTime createdOn) {
    this.createdOn = createdOn;
  }


  public SearchedArtifact createdBy(String createdBy) {
    
    this.createdBy = createdBy;
    return this;
  }

   /**
   * Get createdBy
   * @return createdBy
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")
  @JsonProperty(JSON_PROPERTY_CREATED_BY)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getCreatedBy() {
    return createdBy;
  }


  @JsonProperty(JSON_PROPERTY_CREATED_BY)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setCreatedBy(String createdBy) {
    this.createdBy = createdBy;
  }


  public SearchedArtifact type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(example = "AVRO", required = true, value = "")
  @JsonProperty(JSON_PROPERTY_TYPE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getType() {
    return type;
  }


  @JsonProperty(JSON_PROPERTY_TYPE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setType(String type) {
    this.type = type;
  }


  public SearchedArtifact labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public SearchedArtifact addLabelsItem(String labelsItem) {
    if (this.labels == null) {
      this.labels = new ArrayList<>();
    }
    this.labels.add(labelsItem);
    return this;
  }

   /**
   * Get labels
   * @return labels
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_LABELS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public List<String> getLabels() {
    return labels;
  }


  @JsonProperty(JSON_PROPERTY_LABELS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setLabels(List<String> labels) {
    this.labels = labels;
  }


  public SearchedArtifact state(ArtifactState state) {
    
    this.state = state;
    return this;
  }

   /**
   * Get state
   * @return state
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")
  @JsonProperty(JSON_PROPERTY_STATE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public ArtifactState getState() {
    return state;
  }


  @JsonProperty(JSON_PROPERTY_STATE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setState(ArtifactState state) {
    this.state = state;
  }


  public SearchedArtifact modifiedOn(OffsetDateTime modifiedOn) {
    
    this.modifiedOn = modifiedOn;
    return this;
  }

   /**
   * Get modifiedOn
   * @return modifiedOn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_MODIFIED_ON)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public OffsetDateTime getModifiedOn() {
    return modifiedOn;
  }


  @JsonProperty(JSON_PROPERTY_MODIFIED_ON)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setModifiedOn(OffsetDateTime modifiedOn) {
    this.modifiedOn = modifiedOn;
  }


  public SearchedArtifact modifiedBy(String modifiedBy) {
    
    this.modifiedBy = modifiedBy;
    return this;
  }

   /**
   * Get modifiedBy
   * @return modifiedBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_MODIFIED_BY)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getModifiedBy() {
    return modifiedBy;
  }


  @JsonProperty(JSON_PROPERTY_MODIFIED_BY)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setModifiedBy(String modifiedBy) {
    this.modifiedBy = modifiedBy;
  }


  public SearchedArtifact groupId(String groupId) {
    
    this.groupId = groupId;
    return this;
  }

   /**
   * An ID of a single artifact group.
   * @return groupId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "my-group", value = "An ID of a single artifact group.")
  @JsonProperty(JSON_PROPERTY_GROUP_ID)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getGroupId() {
    return groupId;
  }


  @JsonProperty(JSON_PROPERTY_GROUP_ID)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setGroupId(String groupId) {
    this.groupId = groupId;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SearchedArtifact searchedArtifact = (SearchedArtifact) o;
    return Objects.equals(this.id, searchedArtifact.id) &&
        Objects.equals(this.name, searchedArtifact.name) &&
        Objects.equals(this.description, searchedArtifact.description) &&
        Objects.equals(this.createdOn, searchedArtifact.createdOn) &&
        Objects.equals(this.createdBy, searchedArtifact.createdBy) &&
        Objects.equals(this.type, searchedArtifact.type) &&
        Objects.equals(this.labels, searchedArtifact.labels) &&
        Objects.equals(this.state, searchedArtifact.state) &&
        Objects.equals(this.modifiedOn, searchedArtifact.modifiedOn) &&
        Objects.equals(this.modifiedBy, searchedArtifact.modifiedBy) &&
        Objects.equals(this.groupId, searchedArtifact.groupId);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, description, createdOn, createdBy, type, labels, state, modifiedOn, modifiedBy, groupId);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SearchedArtifact {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    createdOn: ").append(toIndentedString(createdOn)).append("\n");
    sb.append("    createdBy: ").append(toIndentedString(createdBy)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    state: ").append(toIndentedString(state)).append("\n");
    sb.append("    modifiedOn: ").append(toIndentedString(modifiedOn)).append("\n");
    sb.append("    modifiedBy: ").append(toIndentedString(modifiedBy)).append("\n");
    sb.append("    groupId: ").append(toIndentedString(groupId)).append("\n");
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

