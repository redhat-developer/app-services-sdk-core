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
import com.openshift.cloud.api.serviceaccounts.models.SSOHealthResultErrorStackTrace;
import com.openshift.cloud.api.serviceaccounts.models.SSOHealthResultErrorSuppressed;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonTypeName;

/**
 * SSOHealthResultError
 */
@JsonPropertyOrder({
  SSOHealthResultError.JSON_PROPERTY_STACK_TRACE,
  SSOHealthResultError.JSON_PROPERTY_MESSAGE,
  SSOHealthResultError.JSON_PROPERTY_SUPPRESSED,
  SSOHealthResultError.JSON_PROPERTY_LOCALIZED_MESSAGE
})
@JsonTypeName("SSOHealthResult_error")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class SSOHealthResultError {
  public static final String JSON_PROPERTY_STACK_TRACE = "stackTrace";
  private List<SSOHealthResultErrorStackTrace> stackTrace = null;

  public static final String JSON_PROPERTY_MESSAGE = "message";
  private String message;

  public static final String JSON_PROPERTY_SUPPRESSED = "suppressed";
  private List<SSOHealthResultErrorSuppressed> suppressed = null;

  public static final String JSON_PROPERTY_LOCALIZED_MESSAGE = "localizedMessage";
  private String localizedMessage;

  public SSOHealthResultError() { 
  }

  public SSOHealthResultError stackTrace(List<SSOHealthResultErrorStackTrace> stackTrace) {
    
    this.stackTrace = stackTrace;
    return this;
  }

  public SSOHealthResultError addStackTraceItem(SSOHealthResultErrorStackTrace stackTraceItem) {
    if (this.stackTrace == null) {
      this.stackTrace = new ArrayList<>();
    }
    this.stackTrace.add(stackTraceItem);
    return this;
  }

   /**
   * Get stackTrace
   * @return stackTrace
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_STACK_TRACE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public List<SSOHealthResultErrorStackTrace> getStackTrace() {
    return stackTrace;
  }


  @JsonProperty(JSON_PROPERTY_STACK_TRACE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setStackTrace(List<SSOHealthResultErrorStackTrace> stackTrace) {
    this.stackTrace = stackTrace;
  }


  public SSOHealthResultError message(String message) {
    
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


  public SSOHealthResultError suppressed(List<SSOHealthResultErrorSuppressed> suppressed) {
    
    this.suppressed = suppressed;
    return this;
  }

  public SSOHealthResultError addSuppressedItem(SSOHealthResultErrorSuppressed suppressedItem) {
    if (this.suppressed == null) {
      this.suppressed = new ArrayList<>();
    }
    this.suppressed.add(suppressedItem);
    return this;
  }

   /**
   * Get suppressed
   * @return suppressed
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_SUPPRESSED)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public List<SSOHealthResultErrorSuppressed> getSuppressed() {
    return suppressed;
  }


  @JsonProperty(JSON_PROPERTY_SUPPRESSED)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setSuppressed(List<SSOHealthResultErrorSuppressed> suppressed) {
    this.suppressed = suppressed;
  }


  public SSOHealthResultError localizedMessage(String localizedMessage) {
    
    this.localizedMessage = localizedMessage;
    return this;
  }

   /**
   * Get localizedMessage
   * @return localizedMessage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_LOCALIZED_MESSAGE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getLocalizedMessage() {
    return localizedMessage;
  }


  @JsonProperty(JSON_PROPERTY_LOCALIZED_MESSAGE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setLocalizedMessage(String localizedMessage) {
    this.localizedMessage = localizedMessage;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SSOHealthResultError ssOHealthResultError = (SSOHealthResultError) o;
    return Objects.equals(this.stackTrace, ssOHealthResultError.stackTrace) &&
        Objects.equals(this.message, ssOHealthResultError.message) &&
        Objects.equals(this.suppressed, ssOHealthResultError.suppressed) &&
        Objects.equals(this.localizedMessage, ssOHealthResultError.localizedMessage);
  }

  @Override
  public int hashCode() {
    return Objects.hash(stackTrace, message, suppressed, localizedMessage);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SSOHealthResultError {\n");
    sb.append("    stackTrace: ").append(toIndentedString(stackTrace)).append("\n");
    sb.append("    message: ").append(toIndentedString(message)).append("\n");
    sb.append("    suppressed: ").append(toIndentedString(suppressed)).append("\n");
    sb.append("    localizedMessage: ").append(toIndentedString(localizedMessage)).append("\n");
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

