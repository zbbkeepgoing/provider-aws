/*
Copyright 2019 The Crossplane Authors.

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

package v1beta1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Tag represents user-provided metadata that can be associated
type Tag struct {

	// The key name that can be used to look up or retrieve the associated value.
	Key string `json:"key"`

	// The value associated with this tag.
	Value string `json:"value"`
}

// DomainValidationOption validate domain ownership.
type DomainValidationOption struct {
	// Additinal Fully qualified domain name (FQDN),that to secure with an ACM certificate.
	// +immutable
	DomainName string `json:"domainName"`

	// Method to validate certificate
	// +immutable
	ValidationDomain string `json:"validationDomain"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CertificateParameters `json:"forProvider"`
}

// CertificateExternalStatus keeps the state of external resource
type CertificateExternalStatus struct {
	// String that contains the ARN of the issued certificate. This must be of the
	CertificateARN string `json:"certificateARN,omitempty"`

	// Flag to check eligibility for renewal status
	// +kubebuilder:validation:Enum=ELIGIBLE;INELIGIBLE
	RenewalEligibility string `json:"renewalEligibility,omitempty"`

	// Status of the certificate
	// +kubebuilder:validation:Enum=PENDING_VALIDATION;ISSUED;INACTIVE;EXPIRED;VALIDATION_TIMED_OUT;REVOKED;FAILED
	Status string `json:"status,omitempty"`

	// Type of the certificate
	// +kubebuilder:validation:Enum=IMPORTED;AMAZON_ISSUED;PRIVATE
	Type string `json:"type,omitempty"`

	// Contains the CNAME record that you add to your DNS database for domain
	// validation. For more information, see Use DNS to Validate Domain Ownership
	// (https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-dns.html).
	// Note: The CNAME information that you need does not include the name of your
	// domain. If you include your domain name in the DNS database CNAME record,
	// validation fails. For example, if the name is
	// "_a79865eb4cd1a6ab990a45779b4e0b96.yourdomain.com", only
	// "_a79865eb4cd1a6ab990a45779b4e0b96" must be used.
	ResourceRecord *ResourceRecord `json:"resourceRecord,omitempty"`
}

// An CertificateStatus represents the observed state of an Certificate manager.
type CertificateStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CertificateExternalStatus `json:"atProvider,omitempty"`
}

// ResourceRecord Contains a DNS record value that you can use to validate ownership or control of a domain.
type ResourceRecord struct {
	// The name of the DNS record to create in your domain. This is supplied by ACM.
	Name *string `json:"name,omitempty"`

	// The type of DNS record. Currently this can be CNAME.
	// +kubebuilder:validation:Enum=CNAME
	Type *string `json:"type,omitempty"`

	// The value of the CNAME record to add to your DNS database.
	Value *string `json:"value,omitempty"`
}

// CertificateParameters defines the desired state of an AWS Certificate.
type CertificateParameters struct {

	// Region is the region you'd like your Certificate to be created in.
	Region string `json:"region"`

	// The Amazon Resource Name (ARN) of the private certificate authority (CA)that will be used to issue the certificate.
	// +optional
	// +crossplane:generate:reference:type=github.com/crossplane/provider-aws/apis/acmpca/v1beta1.CertificateAuthority
	CertificateAuthorityARN *string `json:"certificateAuthorityARN,omitempty"`

	// CertificateAuthorityARNRef references an AWS ACMPCA CertificateAuthority to retrieve its Arn
	// +optional
	CertificateAuthorityARNRef *xpv1.Reference `json:"certificateAuthorityARNRef,omitempty"`

	// CertificateAuthorityARNSelector selects a reference to an AWS ACMPCA CertificateAuthority to retrieve its Arn
	// +optional
	CertificateAuthorityARNSelector *xpv1.Selector `json:"certificateAuthorityARNSelector,omitempty"`

	// Fully qualified domain name (FQDN),that to secure with an ACM certificate.
	// +immutable
	DomainName string `json:"domainName"`

	// The domain name that you want ACM to use to send you emails so that you can
	// validate domain ownership.
	// +optional
	// +immutable
	DomainValidationOptions []*DomainValidationOption `json:"domainValidationOptions,omitempty"`

	// Currently, you can use this parameter to specify whether to add the certificate
	// to a certificate transparency log. Certificate transparency makes it possible to
	// detect SSL/TLS certificates that have been mistakenly or maliciously issued.
	// Certificates that have not been logged typically produce an error message in a
	// browser. For more information, see Opting Out of Certificate Transparency
	// Logging
	// (https://docs.aws.amazon.com/acm/latest/userguide/acm-bestpractices.html#best-practices-transparency).
	Options *CertificateOptions `json:"options,omitempty"`

	// Subject Alternative Name extension of the ACM certificate.
	// +optional
	// +immutable
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty"`

	// One or more resource tags to associate with the certificate.
	Tags []Tag `json:"tags"`

	// Method to validate certificate.
	// +optional
	// +kubebuilder:validation:Enum=DNS;EMAIL
	ValidationMethod string `json:"validationMethod,omitempty"`
}

// CertificateOptions contains options for your certificate. Currently, you can use
// this only to specify whether to opt in to or out of certificate transparency
// logging. Some browsers require that public certificates issued for your domain
// be recorded in a log. Certificates that are not logged typically generate a
// browser error. Transparency makes it possible for you to detect SSL/TLS
// certificates that have been mistakenly or maliciously issued for your domain.
// For general information, see Certificate Transparency Logging
// (https://docs.aws.amazon.com/acm/latest/userguide/acm-concepts.html#concept-transparency).
type CertificateOptions struct {
	// You can opt out of certificate transparency logging by specifying the DISABLED
	// option. Opt in by specifying ENABLED.
	// +kubebuilder:validation:Enum=ENABLED;DISABLED
	CertificateTransparencyLoggingPreference string `json:"certificateTransparencyLoggingPreference"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// Certificate is a managed resource that represents an AWS Certificate Manager.
// +kubebuilder:printcolumn:name="DOMAINNAME",type="string",JSONPath=".spec.forProvider.domainName"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.atProvider.status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateSpec   `json:"spec"`
	Status CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificate
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}
