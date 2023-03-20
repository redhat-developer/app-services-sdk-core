package com.openshift.cloud.api.kas.auth.models;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class AclBindingListPage extends List implements Parsable {
    /** The items property */
    private java.util.List<AclBinding> items;
    /**
     * Instantiates a new AclBindingListPage and sets the default values.
     * @return a void
     */
    @javax.annotation.Nullable
    public AclBindingListPage() {
        super();
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a AclBindingListPage
     */
    @javax.annotation.Nonnull
    public static AclBindingListPage createFromDiscriminatorValue(@javax.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new AclBindingListPage();
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, java.util.function.Consumer<ParseNode>>
     */
    @javax.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(super.getFieldDeserializers());
        deserializerMap.put("items", (n) -> { this.setItems(n.getCollectionOfObjectValues(AclBinding::createFromDiscriminatorValue)); });
        return deserializerMap;
    }
    /**
     * Gets the items property value. The items property
     * @return a AclBinding
     */
    @javax.annotation.Nullable
    public java.util.List<AclBinding> getItems() {
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
    public void setItems(@javax.annotation.Nullable final java.util.List<AclBinding> value) {
        this.items = value;
    }
}
