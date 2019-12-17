// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.CloudAmqp
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("cloudamqp");

        /// <summary>
        /// Key used to authentication to the CloudAMQP Customer API
        /// </summary>
        public static string? Apikey { get; set; } = __config.Get("apikey") ?? Utilities.GetEnv("CLOUDAMQP_APIKEY");

        /// <summary>
        /// Base URL to CloudAMQP Customer website
        /// </summary>
        public static string? Baseurl { get; set; } = __config.Get("baseurl");

    }
    namespace ConfigTypes
    {
    }
}
