/*
 * Kafka Management API
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * The version of the OpenAPI document: 1.15.0
 * Contact: rhosak-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.openshift.cloud.api.kas.models;

import java.util.Objects;
import java.util.Arrays;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import com.openshift.cloud.api.kas.models.SupportedKafkaBillingModel;
import com.openshift.cloud.api.kas.models.SupportedKafkaSize;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonTypeName;

/**
 * Supported Kafka instance type
 */
@ApiModel(description = "Supported Kafka instance type")
@JsonPropertyOrder({
  SupportedKafkaInstanceType.JSON_PROPERTY_ID,
  SupportedKafkaInstanceType.JSON_PROPERTY_DISPLAY_NAME,
  SupportedKafkaInstanceType.JSON_PROPERTY_SUPPORTED_BILLING_MODELS,
  SupportedKafkaInstanceType.JSON_PROPERTY_SIZES
})
@JsonTypeName("SupportedKafkaInstanceType")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class SupportedKafkaInstanceType {
  public static final String JSON_PROPERTY_ID = "id";
  private String id;

  public static final String JSON_PROPERTY_DISPLAY_NAME = "display_name";
  private String displayName;

  public static final String JSON_PROPERTY_SUPPORTED_BILLING_MODELS = "supported_billing_models";
  private List<SupportedKafkaBillingModel> supportedBillingModels = new ArrayList<>();

  public static final String JSON_PROPERTY_SIZES = "sizes";
  private List<SupportedKafkaSize> sizes = null;

  public SupportedKafkaInstanceType() { 
  }

  public SupportedKafkaInstanceType id(String id) {
    
    this.id = id;
    return this;
  }

   /**
   * Unique identifier of the Kafka instance type.
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Unique identifier of the Kafka instance type.")
  @JsonProperty(JSON_PROPERTY_ID)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getId() {
    return id;
  }


  @JsonProperty(JSON_PROPERTY_ID)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setId(String id) {
    this.id = id;
  }


  public SupportedKafkaInstanceType displayName(String displayName) {
    
    this.displayName = displayName;
    return this;
  }

   /**
   * Human readable name of the supported Kafka instance type
   * @return displayName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Human readable name of the supported Kafka instance type")
  @JsonProperty(JSON_PROPERTY_DISPLAY_NAME)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getDisplayName() {
    return displayName;
  }


  @JsonProperty(JSON_PROPERTY_DISPLAY_NAME)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDisplayName(String displayName) {
    this.displayName = displayName;
  }


  public SupportedKafkaInstanceType supportedBillingModels(List<SupportedKafkaBillingModel> supportedBillingModels) {
    
    this.supportedBillingModels = supportedBillingModels;
    return this;
  }

  public SupportedKafkaInstanceType addSupportedBillingModelsItem(SupportedKafkaBillingModel supportedBillingModelsItem) {
    this.supportedBillingModels.add(supportedBillingModelsItem);
    return this;
  }

   /**
   * A list of available kafka billing models for the instance type. Each kafka billing model item has a unique &#39;id&#39;
   * @return supportedBillingModels
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "A list of available kafka billing models for the instance type. Each kafka billing model item has a unique 'id'")
  @JsonProperty(JSON_PROPERTY_SUPPORTED_BILLING_MODELS)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public List<SupportedKafkaBillingModel> getSupportedBillingModels() {
    return supportedBillingModels;
  }


  @JsonProperty(JSON_PROPERTY_SUPPORTED_BILLING_MODELS)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setSupportedBillingModels(List<SupportedKafkaBillingModel> supportedBillingModels) {
    this.supportedBillingModels = supportedBillingModels;
  }


  public SupportedKafkaInstanceType sizes(List<SupportedKafkaSize> sizes) {
    
    this.sizes = sizes;
    return this;
  }

  public SupportedKafkaInstanceType addSizesItem(SupportedKafkaSize sizesItem) {
    if (this.sizes == null) {
      this.sizes = new ArrayList<>();
    }
    this.sizes.add(sizesItem);
    return this;
  }

   /**
   * A list of Kafka instance sizes available for this instance type
   * @return sizes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "A list of Kafka instance sizes available for this instance type")
  @JsonProperty(JSON_PROPERTY_SIZES)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public List<SupportedKafkaSize> getSizes() {
    return sizes;
  }


  @JsonProperty(JSON_PROPERTY_SIZES)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setSizes(List<SupportedKafkaSize> sizes) {
    this.sizes = sizes;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SupportedKafkaInstanceType supportedKafkaInstanceType = (SupportedKafkaInstanceType) o;
    return Objects.equals(this.id, supportedKafkaInstanceType.id) &&
        Objects.equals(this.displayName, supportedKafkaInstanceType.displayName) &&
        Objects.equals(this.supportedBillingModels, supportedKafkaInstanceType.supportedBillingModels) &&
        Objects.equals(this.sizes, supportedKafkaInstanceType.sizes);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, displayName, supportedBillingModels, sizes);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SupportedKafkaInstanceType {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    displayName: ").append(toIndentedString(displayName)).append("\n");
    sb.append("    supportedBillingModels: ").append(toIndentedString(supportedBillingModels)).append("\n");
    sb.append("    sizes: ").append(toIndentedString(sizes)).append("\n");
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

