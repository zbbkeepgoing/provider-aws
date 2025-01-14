/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package route

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/ec2"

	svcapitypes "github.com/crossplane/provider-aws/apis/ec2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateCreateRouteInput returns a create input.
func GenerateCreateRouteInput(cr *svcapitypes.Route) *svcsdk.CreateRouteInput {
	res := &svcsdk.CreateRouteInput{}

	if cr.Spec.ForProvider.CarrierGatewayID != nil {
		res.SetCarrierGatewayId(*cr.Spec.ForProvider.CarrierGatewayID)
	}
	if cr.Spec.ForProvider.DestinationCIDRBlock != nil {
		res.SetDestinationCidrBlock(*cr.Spec.ForProvider.DestinationCIDRBlock)
	}
	if cr.Spec.ForProvider.DestinationIPv6CIDRBlock != nil {
		res.SetDestinationIpv6CidrBlock(*cr.Spec.ForProvider.DestinationIPv6CIDRBlock)
	}
	if cr.Spec.ForProvider.DestinationPrefixListID != nil {
		res.SetDestinationPrefixListId(*cr.Spec.ForProvider.DestinationPrefixListID)
	}
	if cr.Spec.ForProvider.EgressOnlyInternetGatewayID != nil {
		res.SetEgressOnlyInternetGatewayId(*cr.Spec.ForProvider.EgressOnlyInternetGatewayID)
	}
	if cr.Spec.ForProvider.LocalGatewayID != nil {
		res.SetLocalGatewayId(*cr.Spec.ForProvider.LocalGatewayID)
	}
	if cr.Spec.ForProvider.NetworkInterfaceID != nil {
		res.SetNetworkInterfaceId(*cr.Spec.ForProvider.NetworkInterfaceID)
	}
	if cr.Spec.ForProvider.VPCEndpointID != nil {
		res.SetVpcEndpointId(*cr.Spec.ForProvider.VPCEndpointID)
	}

	return res
}

// GenerateDeleteRouteInput returns a deletion input.
func GenerateDeleteRouteInput(cr *svcapitypes.Route) *svcsdk.DeleteRouteInput {
	res := &svcsdk.DeleteRouteInput{}

	if cr.Spec.ForProvider.DestinationCIDRBlock != nil {
		res.SetDestinationCidrBlock(*cr.Spec.ForProvider.DestinationCIDRBlock)
	}
	if cr.Spec.ForProvider.DestinationIPv6CIDRBlock != nil {
		res.SetDestinationIpv6CidrBlock(*cr.Spec.ForProvider.DestinationIPv6CIDRBlock)
	}
	if cr.Spec.ForProvider.DestinationPrefixListID != nil {
		res.SetDestinationPrefixListId(*cr.Spec.ForProvider.DestinationPrefixListID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "UNKNOWN"
}
