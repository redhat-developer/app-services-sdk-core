package com.openshift.cloud.api.connector.models;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class ConnectorNamespaceRequestMeta extends ObjectMeta implements Parsable {
    /** Name-value string annotations for resource */
    private ConnectorResourceAnnotations annotations;
    /** Namespace name must match pattern `^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$`, or it may be empty to be auto-generated. */
    private String name;
    /**
     * Instantiates a new ConnectorNamespaceRequestMeta and sets the default values.
     * @return a void
     */
    @javax.annotation.Nullable
    public ConnectorNamespaceRequestMeta() {
        super();
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a ConnectorNamespaceRequestMeta
     */
    @javax.annotation.Nonnull
    public static ConnectorNamespaceRequestMeta createFromDiscriminatorValue(@javax.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new ConnectorNamespaceRequestMeta();
    }
    /**
     * Gets the annotations property value. Name-value string annotations for resource
     * @return a ConnectorResourceAnnotations
     */
    @javax.annotation.Nullable
    public ConnectorResourceAnnotations getAnnotations() {
        return this.annotations;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, java.util.function.Consumer<ParseNode>>
     */
    @javax.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(super.getFieldDeserializers());
        deserializerMap.put("annotations", (n) -> { this.setAnnotations(n.getObjectValue(ConnectorResourceAnnotations::createFromDiscriminatorValue)); });
        deserializerMap.put("name", (n) -> { this.setName(n.getStringValue()); });
        return deserializerMap;
    }
    /**
     * Gets the name property value. Namespace name must match pattern `^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$`, or it may be empty to be auto-generated.
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
        super.serialize(writer);
        writer.writeObjectValue("annotations", this.getAnnotations());
        writer.writeStringValue("name", this.getName());
    }
    /**
     * Sets the annotations property value. Name-value string annotations for resource
     * @param value Value to set for the annotations property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setAnnotations(@javax.annotation.Nullable final ConnectorResourceAnnotations value) {
        this.annotations = value;
    }
    /**
     * Sets the name property value. Namespace name must match pattern `^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$`, or it may be empty to be auto-generated.
     * @param value Value to set for the name property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setName(@javax.annotation.Nullable final String value) {
        this.name = value;
    }
}
