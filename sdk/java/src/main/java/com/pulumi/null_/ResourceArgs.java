// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.null_;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ResourceArgs extends com.pulumi.resources.ResourceArgs {

    public static final ResourceArgs Empty = new ResourceArgs();

    /**
     * A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
     * 
     */
    @Import(name="triggers")
    private @Nullable Output<Map<String,String>> triggers;

    /**
     * @return A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
     * 
     */
    public Optional<Output<Map<String,String>>> triggers() {
        return Optional.ofNullable(this.triggers);
    }

    private ResourceArgs() {}

    private ResourceArgs(ResourceArgs $) {
        this.triggers = $.triggers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResourceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResourceArgs $;

        public Builder() {
            $ = new ResourceArgs();
        }

        public Builder(ResourceArgs defaults) {
            $ = new ResourceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param triggers A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
         * 
         * @return builder
         * 
         */
        public Builder triggers(@Nullable Output<Map<String,String>> triggers) {
            $.triggers = triggers;
            return this;
        }

        /**
         * @param triggers A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
         * 
         * @return builder
         * 
         */
        public Builder triggers(Map<String,String> triggers) {
            return triggers(Output.of(triggers));
        }

        public ResourceArgs build() {
            return $;
        }
    }

}
