/*
 * sso.redhat.com API documentation
 * This is the API documentation for sso.redhat.com
 *
 * The version of the OpenAPI document: 5.0.19-SNAPSHOT
 * Contact: it-user-team-list@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.openshift.cloud.api.serviceaccounts.models;

import java.util.Objects;
import java.util.Arrays;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import com.openshift.cloud.api.serviceaccounts.models.SSOHealthResultError;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonTypeName;

/**
 * SSOHealthResult
 */
@JsonPropertyOrder({
  SSOHealthResult.JSON_PROPERTY_MESSAGE,
  SSOHealthResult.JSON_PROPERTY_ERROR,
  SSOHealthResult.JSON_PROPERTY_DETAILS,
  SSOHealthResult.JSON_PROPERTY_TIME,
  SSOHealthResult.JSON_PROPERTY_HEALTHY
})
@JsonTypeName("SSOHealthResult")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class SSOHealthResult {
  public static final String JSON_PROPERTY_MESSAGE = "message";
  private String message;

  public static final String JSON_PROPERTY_ERROR = "error";
  private SSOHealthResultError error;

  public static final String JSON_PROPERTY_DETAILS = "details";
  private Map<String, Object> details = null;

  public static final String JSON_PROPERTY_TIME = "time";
  private Long time;

  public static final String JSON_PROPERTY_HEALTHY = "healthy";
  private Boolean healthy;

  public SSOHealthResult() { 
  }

  public SSOHealthResult message(String message) {
    
    this.message = message;
    return this;
  }

   /**
   * Get message
   * @return message
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_MESSAGE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getMessage() {
    return message;
  }


  @JsonProperty(JSON_PROPERTY_MESSAGE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setMessage(String message) {
    this.message = message;
  }


  public SSOHealthResult error(SSOHealthResultError error) {
    
    this.error = error;
    return this;
  }

   /**
   * Get error
   * @return error
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_ERROR)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public SSOHealthResultError getError() {
    return error;
  }


  @JsonProperty(JSON_PROPERTY_ERROR)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setError(SSOHealthResultError error) {
    this.error = error;
  }


  public SSOHealthResult details(Map<String, Object> details) {
    
    this.details = details;
    return this;
  }

  public SSOHealthResult putDetailsItem(String key, Object detailsItem) {
    if (this.details == null) {
      this.details = new HashMap<>();
    }
    this.details.put(key, detailsItem);
    return this;
  }

   /**
   * Get details
   * @return details
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_DETAILS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Map<String, Object> getDetails() {
    return details;
  }


  @JsonProperty(JSON_PROPERTY_DETAILS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDetails(Map<String, Object> details) {
    this.details = details;
  }


  public SSOHealthResult time(Long time) {
    
    this.time = time;
    return this;
  }

   /**
   * Get time
   * @return time
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_TIME)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Long getTime() {
    return time;
  }


  @JsonProperty(JSON_PROPERTY_TIME)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setTime(Long time) {
    this.time = time;
  }


  public SSOHealthResult healthy(Boolean healthy) {
    
    this.healthy = healthy;
    return this;
  }

   /**
   * Get healthy
   * @return healthy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_HEALTHY)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Boolean getHealthy() {
    return healthy;
  }


  @JsonProperty(JSON_PROPERTY_HEALTHY)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setHealthy(Boolean healthy) {
    this.healthy = healthy;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SSOHealthResult ssOHealthResult = (SSOHealthResult) o;
    return Objects.equals(this.message, ssOHealthResult.message) &&
        Objects.equals(this.error, ssOHealthResult.error) &&
        Objects.equals(this.details, ssOHealthResult.details) &&
        Objects.equals(this.time, ssOHealthResult.time) &&
        Objects.equals(this.healthy, ssOHealthResult.healthy);
  }

  @Override
  public int hashCode() {
    return Objects.hash(message, error, details, time, healthy);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SSOHealthResult {\n");
    sb.append("    message: ").append(toIndentedString(message)).append("\n");
    sb.append("    error: ").append(toIndentedString(error)).append("\n");
    sb.append("    details: ").append(toIndentedString(details)).append("\n");
    sb.append("    time: ").append(toIndentedString(time)).append("\n");
    sb.append("    healthy: ").append(toIndentedString(healthy)).append("\n");
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
