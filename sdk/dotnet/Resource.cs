// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Null
{
    /// <summary>
    /// ## Example Usage
    /// </summary>
    [NullResourceType("null:index/resource:Resource")]
    public partial class Resource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
        /// </summary>
        [Output("triggers")]
        public Output<ImmutableDictionary<string, string>?> Triggers { get; private set; } = null!;


        /// <summary>
        /// Create a Resource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Resource(string name, ResourceArgs? args = null, CustomResourceOptions? options = null)
            : base("null:index/resource:Resource", name, args ?? new ResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Resource(string name, Input<string> id, ResourceState? state = null, CustomResourceOptions? options = null)
            : base("null:index/resource:Resource", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Resource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Resource Get(string name, Input<string> id, ResourceState? state = null, CustomResourceOptions? options = null)
        {
            return new Resource(name, id, state, options);
        }
    }

    public sealed class ResourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("triggers")]
        private InputMap<string>? _triggers;

        /// <summary>
        /// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
        /// </summary>
        public InputMap<string> Triggers
        {
            get => _triggers ?? (_triggers = new InputMap<string>());
            set => _triggers = value;
        }

        public ResourceArgs()
        {
        }
        public static new ResourceArgs Empty => new ResourceArgs();
    }

    public sealed class ResourceState : global::Pulumi.ResourceArgs
    {
        [Input("triggers")]
        private InputMap<string>? _triggers;

        /// <summary>
        /// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
        /// </summary>
        public InputMap<string> Triggers
        {
            get => _triggers ?? (_triggers = new InputMap<string>());
            set => _triggers = value;
        }

        public ResourceState()
        {
        }
        public static new ResourceState Empty => new ResourceState();
    }
}
