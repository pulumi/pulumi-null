// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Null
{
    public static class GetDataSource
    {
        public static Task<GetDataSourceResult> InvokeAsync(GetDataSourceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDataSourceResult>("null:index/getDataSource:getDataSource", args ?? new GetDataSourceArgs(), options.WithDefaults());

        public static Output<GetDataSourceResult> Invoke(GetDataSourceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataSourceResult>("null:index/getDataSource:getDataSource", args ?? new GetDataSourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataSourceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// If set, its literal value will be stored and returned. If not, its value defaults to `"default"`. This argument exists primarily for testing and has little practical use.
        /// </summary>
        [Input("hasComputedDefault")]
        public string? HasComputedDefault { get; set; }

        [Input("inputs")]
        private Dictionary<string, string>? _inputs;

        /// <summary>
        /// A map of arbitrary strings that is copied into the `outputs` attribute, and accessible directly for interpolation.
        /// </summary>
        public Dictionary<string, string> Inputs
        {
            get => _inputs ?? (_inputs = new Dictionary<string, string>());
            set => _inputs = value;
        }

        public GetDataSourceArgs()
        {
        }
        public static new GetDataSourceArgs Empty => new GetDataSourceArgs();
    }

    public sealed class GetDataSourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// If set, its literal value will be stored and returned. If not, its value defaults to `"default"`. This argument exists primarily for testing and has little practical use.
        /// </summary>
        [Input("hasComputedDefault")]
        public Input<string>? HasComputedDefault { get; set; }

        [Input("inputs")]
        private InputMap<string>? _inputs;

        /// <summary>
        /// A map of arbitrary strings that is copied into the `outputs` attribute, and accessible directly for interpolation.
        /// </summary>
        public InputMap<string> Inputs
        {
            get => _inputs ?? (_inputs = new InputMap<string>());
            set => _inputs = value;
        }

        public GetDataSourceInvokeArgs()
        {
        }
        public static new GetDataSourceInvokeArgs Empty => new GetDataSourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataSourceResult
    {
        /// <summary>
        /// If set, its literal value will be stored and returned. If not, its value defaults to `"default"`. This argument exists primarily for testing and has little practical use.
        /// </summary>
        public readonly string HasComputedDefault;
        /// <summary>
        /// This attribute is only present for some legacy compatibility issues and should not be used. It will be removed in a future version.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A map of arbitrary strings that is copied into the `outputs` attribute, and accessible directly for interpolation.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Inputs;
        /// <summary>
        /// After the data source is "read", a copy of the `inputs` map.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Outputs;
        public readonly string Random;

        [OutputConstructor]
        private GetDataSourceResult(
            string hasComputedDefault,

            string id,

            ImmutableDictionary<string, string>? inputs,

            ImmutableDictionary<string, string> outputs,

            string random)
        {
            HasComputedDefault = hasComputedDefault;
            Id = id;
            Inputs = inputs;
            Outputs = outputs;
            Random = random;
        }
    }
}
