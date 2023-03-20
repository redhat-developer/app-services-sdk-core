package com.openshift.cloud.api.registry.instance.models;

import com.microsoft.kiota.ApiException;
import com.microsoft.kiota.serialization.AdditionalDataHolder;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
/**
 * All error responses, whether `4xx` or `5xx` will include one of these as the responsebody.
 */
public class Error extends ApiException implements AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private Map<String, Object> additionalData;
    /** Full details about the error.  This might contain a server stack trace, for example. */
    private String detail;
    /** The server-side error code. */
    private Integer errorCode;
    /** The short error message. */
    private String message;
    /** The error name - typically the classname of the exception thrown by the server. */
    private String name;
    /**
     * Instantiates a new Error and sets the default values.
     * @return a void
     */
    @javax.annotation.Nullable
    public Error() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a Error
     */
    @javax.annotation.Nonnull
    public static Error createFromDiscriminatorValue(@javax.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new Error();
    }
    /**
     * Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a Map<String, Object>
     */
    @javax.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return this.additionalData;
    }
    /**
     * Gets the detail property value. Full details about the error.  This might contain a server stack trace, for example.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getDetail() {
        return this.detail;
    }
    /**
     * Gets the error_code property value. The server-side error code.
     * @return a integer
     */
    @javax.annotation.Nullable
    public Integer getErrorCode() {
        return this.errorCode;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, java.util.function.Consumer<ParseNode>>
     */
    @javax.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(4);
        deserializerMap.put("detail", (n) -> { this.setDetail(n.getStringValue()); });
        deserializerMap.put("error_code", (n) -> { this.setErrorCode(n.getIntegerValue()); });
        deserializerMap.put("message", (n) -> { this.setMessage(n.getStringValue()); });
        deserializerMap.put("name", (n) -> { this.setName(n.getStringValue()); });
        return deserializerMap;
    }
    /**
     * Gets the message property value. The short error message.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getMessage() {
        return this.message;
    }
    /**
     * Gets the name property value. The error name - typically the classname of the exception thrown by the server.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getName() {
        return this.name;
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    @javax.annotation.Nonnull
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("detail", this.getDetail());
        writer.writeIntegerValue("error_code", this.getErrorCode());
        writer.writeStringValue("message", this.getMessage());
        writer.writeStringValue("name", this.getName());
        writer.writeAdditionalData(this.getAdditionalData());
    }
    /**
     * Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setAdditionalData(@javax.annotation.Nullable final Map<String, Object> value) {
        this.additionalData = value;
    }
    /**
     * Sets the detail property value. Full details about the error.  This might contain a server stack trace, for example.
     * @param value Value to set for the detail property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setDetail(@javax.annotation.Nullable final String value) {
        this.detail = value;
    }
    /**
     * Sets the error_code property value. The server-side error code.
     * @param value Value to set for the errorCode property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setErrorCode(@javax.annotation.Nullable final Integer value) {
        this.errorCode = value;
    }
    /**
     * Sets the message property value. The short error message.
     * @param value Value to set for the message property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setMessage(@javax.annotation.Nullable final String value) {
        this.message = value;
    }
    /**
     * Sets the name property value. The error name - typically the classname of the exception thrown by the server.
     * @param value Value to set for the name property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setName(@javax.annotation.Nullable final String value) {
        this.name = value;
    }
}
