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

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// DatabaseParameters defines the desired state of Database
type DatabaseParameters struct {
	// Region is which region the Database will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The ID of the Data Catalog in which to create the database. If none is provided,
	// the Amazon Web Services account ID is used by default.
	CatalogID                *string `json:"catalogID,omitempty"`
	CustomDatabaseParameters `json:",inline"`
}

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DatabaseParameters `json:"forProvider"`
}

// DatabaseObservation defines the observed state of Database
type DatabaseObservation struct {
}

// DatabaseStatus defines the observed state of Database.
type DatabaseStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Database is the Schema for the Databases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseSpec   `json:"spec"`
	Status            DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

// Repository type metadata.
var (
	DatabaseKind             = "Database"
	DatabaseGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseKind}.String()
	DatabaseKindAPIVersion   = DatabaseKind + "." + GroupVersion.String()
	DatabaseGroupVersionKind = GroupVersion.WithKind(DatabaseKind)
)

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
