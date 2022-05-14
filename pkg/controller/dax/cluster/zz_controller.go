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

package cluster

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/dax"
	svcsdk "github.com/aws/aws-sdk-go/service/dax"
	svcsdkapi "github.com/aws/aws-sdk-go/service/dax/daxiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/dax/v1alpha1"
	awsclient "github.com/crossplane/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an Cluster resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create Cluster in AWS"
	errUpdate        = "cannot update Cluster in AWS"
	errDescribe      = "failed to describe Cluster"
	errDelete        = "failed to delete Cluster"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.Cluster)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := awsclient.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.Cluster)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeClustersInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeClustersWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	resp = e.filterList(cr, resp)
	if len(resp.Clusters) == 0 {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateCluster(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, err := e.isUpToDate(cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.Cluster)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateClusterInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateClusterWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.Cluster.ActiveNodes != nil {
		cr.Status.AtProvider.ActiveNodes = resp.Cluster.ActiveNodes
	} else {
		cr.Status.AtProvider.ActiveNodes = nil
	}
	if resp.Cluster.ClusterArn != nil {
		cr.Status.AtProvider.ClusterARN = resp.Cluster.ClusterArn
	} else {
		cr.Status.AtProvider.ClusterARN = nil
	}
	if resp.Cluster.ClusterDiscoveryEndpoint != nil {
		f2 := &svcapitypes.Endpoint{}
		if resp.Cluster.ClusterDiscoveryEndpoint.Address != nil {
			f2.Address = resp.Cluster.ClusterDiscoveryEndpoint.Address
		}
		if resp.Cluster.ClusterDiscoveryEndpoint.Port != nil {
			f2.Port = resp.Cluster.ClusterDiscoveryEndpoint.Port
		}
		if resp.Cluster.ClusterDiscoveryEndpoint.URL != nil {
			f2.URL = resp.Cluster.ClusterDiscoveryEndpoint.URL
		}
		cr.Status.AtProvider.ClusterDiscoveryEndpoint = f2
	} else {
		cr.Status.AtProvider.ClusterDiscoveryEndpoint = nil
	}
	if resp.Cluster.ClusterEndpointEncryptionType != nil {
		cr.Spec.ForProvider.ClusterEndpointEncryptionType = resp.Cluster.ClusterEndpointEncryptionType
	} else {
		cr.Spec.ForProvider.ClusterEndpointEncryptionType = nil
	}
	if resp.Cluster.ClusterName != nil {
		cr.Status.AtProvider.ClusterName = resp.Cluster.ClusterName
	} else {
		cr.Status.AtProvider.ClusterName = nil
	}
	if resp.Cluster.Description != nil {
		cr.Spec.ForProvider.Description = resp.Cluster.Description
	} else {
		cr.Spec.ForProvider.Description = nil
	}
	if resp.Cluster.IamRoleArn != nil {
		cr.Status.AtProvider.IAMRoleARN = resp.Cluster.IamRoleArn
	} else {
		cr.Status.AtProvider.IAMRoleARN = nil
	}
	if resp.Cluster.NodeIdsToRemove != nil {
		f7 := []*string{}
		for _, f7iter := range resp.Cluster.NodeIdsToRemove {
			var f7elem string
			f7elem = *f7iter
			f7 = append(f7, &f7elem)
		}
		cr.Status.AtProvider.NodeIDsToRemove = f7
	} else {
		cr.Status.AtProvider.NodeIDsToRemove = nil
	}
	if resp.Cluster.NodeType != nil {
		cr.Spec.ForProvider.NodeType = resp.Cluster.NodeType
	} else {
		cr.Spec.ForProvider.NodeType = nil
	}
	if resp.Cluster.Nodes != nil {
		f9 := []*svcapitypes.Node{}
		for _, f9iter := range resp.Cluster.Nodes {
			f9elem := &svcapitypes.Node{}
			if f9iter.AvailabilityZone != nil {
				f9elem.AvailabilityZone = f9iter.AvailabilityZone
			}
			if f9iter.Endpoint != nil {
				f9elemf1 := &svcapitypes.Endpoint{}
				if f9iter.Endpoint.Address != nil {
					f9elemf1.Address = f9iter.Endpoint.Address
				}
				if f9iter.Endpoint.Port != nil {
					f9elemf1.Port = f9iter.Endpoint.Port
				}
				if f9iter.Endpoint.URL != nil {
					f9elemf1.URL = f9iter.Endpoint.URL
				}
				f9elem.Endpoint = f9elemf1
			}
			if f9iter.NodeCreateTime != nil {
				f9elem.NodeCreateTime = &metav1.Time{*f9iter.NodeCreateTime}
			}
			if f9iter.NodeId != nil {
				f9elem.NodeID = f9iter.NodeId
			}
			if f9iter.NodeStatus != nil {
				f9elem.NodeStatus = f9iter.NodeStatus
			}
			if f9iter.ParameterGroupStatus != nil {
				f9elem.ParameterGroupStatus = f9iter.ParameterGroupStatus
			}
			f9 = append(f9, f9elem)
		}
		cr.Status.AtProvider.Nodes = f9
	} else {
		cr.Status.AtProvider.Nodes = nil
	}
	if resp.Cluster.NotificationConfiguration != nil {
		f10 := &svcapitypes.NotificationConfiguration{}
		if resp.Cluster.NotificationConfiguration.TopicArn != nil {
			f10.TopicARN = resp.Cluster.NotificationConfiguration.TopicArn
		}
		if resp.Cluster.NotificationConfiguration.TopicStatus != nil {
			f10.TopicStatus = resp.Cluster.NotificationConfiguration.TopicStatus
		}
		cr.Status.AtProvider.NotificationConfiguration = f10
	} else {
		cr.Status.AtProvider.NotificationConfiguration = nil
	}
	if resp.Cluster.ParameterGroup != nil {
		f11 := &svcapitypes.ParameterGroupStatus_SDK{}
		if resp.Cluster.ParameterGroup.NodeIdsToReboot != nil {
			f11f0 := []*string{}
			for _, f11f0iter := range resp.Cluster.ParameterGroup.NodeIdsToReboot {
				var f11f0elem string
				f11f0elem = *f11f0iter
				f11f0 = append(f11f0, &f11f0elem)
			}
			f11.NodeIDsToReboot = f11f0
		}
		if resp.Cluster.ParameterGroup.ParameterApplyStatus != nil {
			f11.ParameterApplyStatus = resp.Cluster.ParameterGroup.ParameterApplyStatus
		}
		if resp.Cluster.ParameterGroup.ParameterGroupName != nil {
			f11.ParameterGroupName = resp.Cluster.ParameterGroup.ParameterGroupName
		}
		cr.Status.AtProvider.ParameterGroup = f11
	} else {
		cr.Status.AtProvider.ParameterGroup = nil
	}
	if resp.Cluster.PreferredMaintenanceWindow != nil {
		cr.Spec.ForProvider.PreferredMaintenanceWindow = resp.Cluster.PreferredMaintenanceWindow
	} else {
		cr.Spec.ForProvider.PreferredMaintenanceWindow = nil
	}
	if resp.Cluster.SSEDescription != nil {
		f13 := &svcapitypes.SSEDescription{}
		if resp.Cluster.SSEDescription.Status != nil {
			f13.Status = resp.Cluster.SSEDescription.Status
		}
		cr.Status.AtProvider.SSEDescription = f13
	} else {
		cr.Status.AtProvider.SSEDescription = nil
	}
	if resp.Cluster.SecurityGroups != nil {
		f14 := []*svcapitypes.SecurityGroupMembership{}
		for _, f14iter := range resp.Cluster.SecurityGroups {
			f14elem := &svcapitypes.SecurityGroupMembership{}
			if f14iter.SecurityGroupIdentifier != nil {
				f14elem.SecurityGroupIdentifier = f14iter.SecurityGroupIdentifier
			}
			if f14iter.Status != nil {
				f14elem.Status = f14iter.Status
			}
			f14 = append(f14, f14elem)
		}
		cr.Status.AtProvider.SecurityGroups = f14
	} else {
		cr.Status.AtProvider.SecurityGroups = nil
	}
	if resp.Cluster.Status != nil {
		cr.Status.AtProvider.Status = resp.Cluster.Status
	} else {
		cr.Status.AtProvider.Status = nil
	}
	if resp.Cluster.SubnetGroup != nil {
		cr.Status.AtProvider.SubnetGroup = resp.Cluster.SubnetGroup
	} else {
		cr.Status.AtProvider.SubnetGroup = nil
	}
	if resp.Cluster.TotalNodes != nil {
		cr.Status.AtProvider.TotalNodes = resp.Cluster.TotalNodes
	} else {
		cr.Status.AtProvider.TotalNodes = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.Cluster)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateClusterInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateClusterWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.Cluster)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteClusterInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteClusterWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.DAXAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		filterList:     nopFilterList,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.DAXAPI
	preObserve     func(context.Context, *svcapitypes.Cluster, *svcsdk.DescribeClustersInput) error
	postObserve    func(context.Context, *svcapitypes.Cluster, *svcsdk.DescribeClustersOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	filterList     func(*svcapitypes.Cluster, *svcsdk.DescribeClustersOutput) *svcsdk.DescribeClustersOutput
	lateInitialize func(*svcapitypes.ClusterParameters, *svcsdk.DescribeClustersOutput) error
	isUpToDate     func(*svcapitypes.Cluster, *svcsdk.DescribeClustersOutput) (bool, error)
	preCreate      func(context.Context, *svcapitypes.Cluster, *svcsdk.CreateClusterInput) error
	postCreate     func(context.Context, *svcapitypes.Cluster, *svcsdk.CreateClusterOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.Cluster, *svcsdk.DeleteClusterInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.Cluster, *svcsdk.DeleteClusterOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.Cluster, *svcsdk.UpdateClusterInput) error
	postUpdate     func(context.Context, *svcapitypes.Cluster, *svcsdk.UpdateClusterOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.Cluster, *svcsdk.DescribeClustersInput) error {
	return nil
}
func nopPostObserve(_ context.Context, _ *svcapitypes.Cluster, _ *svcsdk.DescribeClustersOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopFilterList(_ *svcapitypes.Cluster, list *svcsdk.DescribeClustersOutput) *svcsdk.DescribeClustersOutput {
	return list
}

func nopLateInitialize(*svcapitypes.ClusterParameters, *svcsdk.DescribeClustersOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.Cluster, *svcsdk.DescribeClustersOutput) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.Cluster, *svcsdk.CreateClusterInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.Cluster, _ *svcsdk.CreateClusterOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.Cluster, *svcsdk.DeleteClusterInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.Cluster, _ *svcsdk.DeleteClusterOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.Cluster, *svcsdk.UpdateClusterInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.Cluster, _ *svcsdk.UpdateClusterOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
