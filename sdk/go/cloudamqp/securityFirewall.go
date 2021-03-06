// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SecurityFirewall struct {
	s *pulumi.ResourceState
}

// NewSecurityFirewall registers a new resource with the given unique name, arguments, and options.
func NewSecurityFirewall(ctx *pulumi.Context,
	name string, args *SecurityFirewallArgs, opts ...pulumi.ResourceOpt) (*SecurityFirewall, error) {
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil || args.Rules == nil {
		return nil, errors.New("missing required argument 'Rules'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["instanceId"] = nil
		inputs["rules"] = nil
	} else {
		inputs["instanceId"] = args.InstanceId
		inputs["rules"] = args.Rules
	}
	s, err := ctx.RegisterResource("cloudamqp:index/securityFirewall:SecurityFirewall", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecurityFirewall{s: s}, nil
}

// GetSecurityFirewall gets an existing SecurityFirewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityFirewall(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecurityFirewallState, opts ...pulumi.ResourceOpt) (*SecurityFirewall, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["instanceId"] = state.InstanceId
		inputs["rules"] = state.Rules
	}
	s, err := ctx.ReadResource("cloudamqp:index/securityFirewall:SecurityFirewall", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecurityFirewall{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SecurityFirewall) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SecurityFirewall) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Instance identifier
func (r *SecurityFirewall) InstanceId() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["instanceId"])
}

func (r *SecurityFirewall) Rules() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["rules"])
}

// Input properties used for looking up and filtering SecurityFirewall resources.
type SecurityFirewallState struct {
	// Instance identifier
	InstanceId interface{}
	Rules interface{}
}

// The set of arguments for constructing a SecurityFirewall resource.
type SecurityFirewallArgs struct {
	// Instance identifier
	InstanceId interface{}
	Rules interface{}
}
