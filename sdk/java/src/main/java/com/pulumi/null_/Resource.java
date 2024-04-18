// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.null_;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.null_.ResourceArgs;
import com.pulumi.null_.Utilities;
import com.pulumi.null_.inputs.ResourceState;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.instance;
 * import com.pulumi.aws.InstanceArgs;
 * import com.pulumi.null.Resource;
 * import com.pulumi.null.ResourceArgs;
 * import com.pulumi.codegen.internal.KeyedValue;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         for (var i = 0; i &lt; 3; i++) {
 *             new Instance(&#34;cluster-&#34; + i, InstanceArgs.builder()            
 *                 .ami(&#34;ami-0dcc1e21636832c5d&#34;)
 *                 .instanceType(&#34;m5.large&#34;)
 *                 .build());
 * 
 *         
 * }
 *         // The primary use-case for the null resource is as a do-nothing container
 *         // for arbitrary actions taken by a provisioner.
 *         //
 *         // In this example, three EC2 instances are created and then a
 *         // null_resource instance is used to gather data about all three
 *         // and execute a single action that affects them all. Due to the triggers
 *         // map, the null_resource will be replaced each time the instance ids
 *         // change, and thus the remote-exec provisioner will be re-run.
 *         var clusterResource = new Resource(&#34;clusterResource&#34;, ResourceArgs.builder()        
 *             .triggers(Map.of(&#34;cluster_instance_ids&#34;, StdFunctions.join(JoinArgs.builder()
 *                 .separator(&#34;,&#34;)
 *                 .input(cluster.stream().map(element -&gt; element.id()).collect(toList()))
 *                 .build()).result()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="null:index/resource:Resource")
public class Resource extends com.pulumi.resources.CustomResource {
    /**
     * A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
     * 
     */
    @Export(name="triggers", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> triggers;

    /**
     * @return A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
     * 
     */
    public Output<Optional<Map<String,String>>> triggers() {
        return Codegen.optional(this.triggers);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Resource(String name) {
        this(name, ResourceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Resource(String name, @Nullable ResourceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Resource(String name, @Nullable ResourceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("null:index/resource:Resource", name, args == null ? ResourceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Resource(String name, Output<String> id, @Nullable ResourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("null:index/resource:Resource", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Resource get(String name, Output<String> id, @Nullable ResourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Resource(name, id, state, options);
    }
}
