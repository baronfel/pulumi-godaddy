// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudamqp
{
    public partial class Notification : Pulumi.CustomResource
    {
        /// <summary>
        /// Instance identifier
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Type of the notification, valid options are: email, webhook, pagerduty, victorops, opsgenie, opsgenie-eu,
        /// slack
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Notification endpoint, where to send the notifcation
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a Notification resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Notification(string name, NotificationArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/notification:Notification", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Notification(string name, Input<string> id, NotificationState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/notification:Notification", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Notification resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Notification Get(string name, Input<string> id, NotificationState? state = null, CustomResourceOptions? options = null)
        {
            return new Notification(name, id, state, options);
        }
    }

    public sealed class NotificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance identifier
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// Type of the notification, valid options are: email, webhook, pagerduty, victorops, opsgenie, opsgenie-eu,
        /// slack
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Notification endpoint, where to send the notifcation
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public NotificationArgs()
        {
        }
    }

    public sealed class NotificationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance identifier
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// Type of the notification, valid options are: email, webhook, pagerduty, victorops, opsgenie, opsgenie-eu,
        /// slack
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Notification endpoint, where to send the notifcation
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public NotificationState()
        {
        }
    }
}