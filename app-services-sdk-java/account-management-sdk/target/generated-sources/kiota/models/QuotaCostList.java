package com.openshift.cloud.api.accountmanagement.models;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class QuotaCostList extends List implements Parsable {
    /** The items property */
    private java.util.List<QuotaCost> items;
    /**
     * Instantiates a new QuotaCostList and sets the default values.
     * @return a void
     */
    @javax.annotation.Nullable
    public QuotaCostList() {
        super();
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a QuotaCostList
     */
    @javax.annotation.Nonnull
    public static QuotaCostList createFromDiscriminatorValue(@javax.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new QuotaCostList();
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, java.util.function.Consumer<ParseNode>>
     */
    @javax.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(super.getFieldDeserializers());
        deserializerMap.put("items", (n) -> { this.setItems(n.getCollectionOfObjectValues(QuotaCost::createFromDiscriminatorValue)); });
        return deserializerMap;
    }
    /**
     * Gets the items property value. The items property
     * @return a QuotaCost
     */
    @javax.annotation.Nullable
    public java.util.List<QuotaCost> getItems() {
        return this.items;
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
        writer.writeCollectionOfObjectValues("items", this.getItems());
    }
    /**
     * Sets the items property value. The items property
     * @param value Value to set for the items property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setItems(@javax.annotation.Nullable final java.util.List<QuotaCost> value) {
        this.items = value;
    }
}
