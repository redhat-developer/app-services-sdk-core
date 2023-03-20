package com.openshift.cloud.api.accountmanagement.models;

import com.microsoft.kiota.serialization.AdditionalDataHolder;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class LandingPageEstateItem implements AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private Map<String, Object> additionalData;
    /** The count property */
    private Long count;
    /** The id property */
    private String id;
    /** The shape property */
    private LandingPageLinkShape shape;
    /**
     * Instantiates a new LandingPageEstateItem and sets the default values.
     * @return a void
     */
    @javax.annotation.Nullable
    public LandingPageEstateItem() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a LandingPageEstateItem
     */
    @javax.annotation.Nonnull
    public static LandingPageEstateItem createFromDiscriminatorValue(@javax.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new LandingPageEstateItem();
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
     * Gets the count property value. The count property
     * @return a int64
     */
    @javax.annotation.Nullable
    public Long getCount() {
        return this.count;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, java.util.function.Consumer<ParseNode>>
     */
    @javax.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(3);
        deserializerMap.put("count", (n) -> { this.setCount(n.getLongValue()); });
        deserializerMap.put("id", (n) -> { this.setId(n.getStringValue()); });
        deserializerMap.put("shape", (n) -> { this.setShape(n.getObjectValue(LandingPageLinkShape::createFromDiscriminatorValue)); });
        return deserializerMap;
    }
    /**
     * Gets the id property value. The id property
     * @return a string
     */
    @javax.annotation.Nullable
    public String getId() {
        return this.id;
    }
    /**
     * Gets the shape property value. The shape property
     * @return a LandingPageLinkShape
     */
    @javax.annotation.Nullable
    public LandingPageLinkShape getShape() {
        return this.shape;
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    @javax.annotation.Nonnull
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeLongValue("count", this.getCount());
        writer.writeStringValue("id", this.getId());
        writer.writeObjectValue("shape", this.getShape());
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
     * Sets the count property value. The count property
     * @param value Value to set for the count property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setCount(@javax.annotation.Nullable final Long value) {
        this.count = value;
    }
    /**
     * Sets the id property value. The id property
     * @param value Value to set for the id property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setId(@javax.annotation.Nullable final String value) {
        this.id = value;
    }
    /**
     * Sets the shape property value. The shape property
     * @param value Value to set for the shape property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setShape(@javax.annotation.Nullable final LandingPageLinkShape value) {
        this.shape = value;
    }
}
