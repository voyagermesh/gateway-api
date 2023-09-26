package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
	_ "sigs.k8s.io/gateway-api/apis/v1alpha2"
	"sigs.k8s.io/gateway-api/apis/v1beta1"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=gateway-api
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BackendTLSPolicy provides a way to publish TLS configuration
// that enables a gateway client to connect to a backend pod.
type BackendTLSPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of BackendTLSPolicy.
	Spec BackendTLSPolicySpec `json:"spec"`

	// Status defines the current state of BackendTLSPolicy.
	Status BackendTLSPolicyStatus `json:"status,omitempty"`
}

// BackendTLSPolicySpec defines the desired state of
// BackendTLSPolicy.
// Note: there is no Override or Default policy configuration.
//
// Support: Core
type BackendTLSPolicySpec struct {
	// TargetRef identifies an API object to apply policy to.
	// Services are the only valid API target references.
	// Note that this config applies to the entire referenced resource
	// by default, but this default may change in the future to provide
	// a more granular application of the policy.
	TargetRef v1alpha2.PolicyTargetReference `json:"targetRef"`

	// TLS contains backend TLS policy configuration.
	TLS *BackendTLSPolicyConfig `json:"tls"`
}

// BackendTLSPolicyConfig contains backend TLS policy configuration.
type BackendTLSPolicyConfig struct {
	// CertRefs contains one or more references to
	// Kubernetes objects that contain PEM-encoded TLS certificates,
	// which are used to establish a TLS handshake between the gateway
	// and backend pod.
	//
	// If CertRefs is empty or unspecified, then StandardCerts must
	// be specified.  Only one of CertRefs or StandardCerts may be
	// specified, not both.
	//
	// If CertRefs is empty or unspecified, then system trusted
	// certificates should be used. If there are none, or the
	// implementation doesn't define system trusted certificates,
	// then a TLS connection must fail.
	//
	// References to a resource in a different namespace are
	// invalid.
	//
	// A single CertRef to a Kubernetes ConfigMap kind has "Core"
	// support.  Implementations MAY choose to support attaching
	// multiple certificates to a backend, but this behavior is
	// implementation-specific.  Also implementation-specific is
	// a CertRef of other object kinds, e.g. Secret.
	//
	// Support: Core - An optional single reference to a Kubernetes
	// ConfigMap.
	//
	// Support: Implementation-specific (No reference, more than one
	// reference, or resource types other than ConfigMaps.
	// Service mesh may ignore.)
	//
	// +kubebuilder:validation:MaxItems=8
	// +optional
	CertRefs []ConfigMapObjectReference `json:"certRefs,omitempty"`

	// StandardCerts specifies whether system CA certificates may
	// be used in the TLS handshake between the gateway and
	// backend pod.
	//
	// If StandardCerts is unspecified or set to "", then CertRefs must
	// be specified with at least one entry for a valid configuration.
	// If StandardCerts is unspecified or set to "", then CertRefs must
	// be specified.  Only one of CertRefs or StandardCerts may be
	// specified, not both.
	//
	// StandardCerts must be set to "System" when CertRefs is unspecified.
	//
	// If StandardCerts is set to "System", then the system trusted
	// certificates should be used. If there are none, or the
	// implementation doesn't define system trusted certificates,
	// then a TLS connection must fail.
	//
	// Support: Core - An optional value to specify whether to use
	// system certificates or not.
	//
	// Support: Implementation-specific (In the absence of support
	// for usable system certs, may be ignored. Service mesh may ignore.)
	//
	// +optional
	StandardCerts *StandardCertType `json:"standardCerts,omitempty"`

	// Hostname is the Server Name Indication that the Gateway uses to
	// connect to the backend.  It represents the fully qualified domain
	// name of a network host, as defined by RFC1123 - except that numeric
	// IP addresses are not allowed. Each label of the FQDN must consist
	// of lower case alphanumeric characters or '-', and must start and
	// end with an alphanumeric character.  No other punctuation is allowed.
	// Wildcard domain names are specifically disallowed.
	//
	// It specifies the hostname that may authenticate, and must be in the
	// certificate served by the matching backend.
	//
	// Support: Core - A required value used by the Gateway to connect to
	// the backend when a BackendTLSPolicy is specified.
	Hostname v1beta1.PreciseHostname `json:"hostname"`
}

// StandardCertType is the type of CA certificate that will be used when
// the TLS.certRefs is unspecified.
// +kubebuilder:validation:Enum=System
type StandardCertType string

const (
	StandardCertSystem StandardCertType = "System"
)

// ConfigMapObjectReference identifies an API object including its namespace,
// defaulting to ConfigMap.
//
// The API object must be valid in the cluster; the Group and Kind must
// be registered in the cluster for this reference to be valid.
//
// References to objects with invalid Group and Kind are not valid, and must
// be rejected by the implementation, with appropriate Conditions set
// on the containing object.
type ConfigMapObjectReference struct {
	// Group is the group of the referent.  For example, "gateway.networking.k8s.io".
	// When unspecified or empty string, core API group is inferred.
	//
	// +optional
	// +kubebuilder:default=""
	Group *v1beta1.Group `json:"group"`

	// Kind is the kind of the referent.  For example, "ConfigMap".
	//
	// +optional
	// +kubebuilder:default=ConfigMap
	Kind *v1beta1.Kind `json:"kind"`

	// Name is the metadata.name of the referenced config map.
	// +kubebuilder:validation:Required
	Name v1beta1.ObjectName `json:"name"`

	// Namespace is the namespace of the referenced object. When unspecified, the local
	// namespace is inferred.
	//
	// Note that when a namespace different than the local namespace is specified,
	// a ReferenceGrant object is required in the referent namespace to allow that
	// namespace's owner to accept the reference. See the ReferenceGrant
	// documentation for details.
	//
	// Support: Core
	//
	// +optional
	Namespace *v1beta1.Namespace `json:"namespace,omitempty"`
}

// BackendTLSPolicyStatus defines the observed state of BackendTLSPolicy.
type BackendTLSPolicyStatus struct {
	// Conditions describe the current conditions of the BackendTLSPolicy.
	//
	// Implementations should prefer to express BackendTLSPolicy
	// conditions using the `BackendTLSPolicyConditionType` and
	// `BackendTLSPolicyConditionReason` constants so that
	// operators and tools can converge on a common vocabulary to
	// describe BackendTLSPolicy state.
	// Known condition types are:
	//
	// * “Accepted”
	//
	// +optional
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	// +kubebuilder:default={{type: "Accepted", status: "Unknown", reason:"Pending", message:"Waiting for validation", lastTransitionTime: "1970-01-01T00:00:00Z"},}
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// BackendTLSPolicy contains a list of BackendTLSPolicy
type BackendTLSPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendTLSPolicy `json:"items"`
}

// BackendTLSPolicyConditionType is the type of a condition used
// as a signal by BackendTLSPolicy.  This type should be used with
// the BackendTLSPolicyStatus.Conditions field.
type BackendTLSPolicyConditionType string

//	BackendTLSPolicyConditionReason is a reason that explains why a
//
// particular BackendTLSPolicyConditionType was generated.
type BackendTLSPolicyConditionReason string

const (
	// BackendTLSPolicyConditionAccepted This condition indicates that the BackendTLSPolicy has been
	// accepted as valid.
	// Possible reason for this condition to be True is:
	//
	// * “Accepted”
	//
	// Possible reasons for this condition to be False are:
	//
	// * “Invalid”
	// * “Pending”
	BackendTLSPolicyConditionAccepted BackendTLSPolicyConditionType = "Accepted"

	// BackendTLSPolicyReasonAccepted This reason is used with the “Accepted” condition when the condition is true.
	BackendTLSPolicyReasonAccepted BackendTLSPolicyConditionReason = "Valid"

	// BackendTLSPolicyReasonInvalid This reason is used with the “Accepted” condition when the BackendTLSPolicy is invalid,
	// e.g. use of a CertRef that crosses namespace boundaries.
	BackendTLSPolicyReasonInvalid BackendTLSPolicyConditionReason = "Invalid"

	// BackendTLSPolicyReasonPending This reason is used with the “Accepted” condition when the BackendTLSPolicy is pending validation.
	BackendTLSPolicyReasonPending BackendTLSPolicyConditionReason = "Pending"
)

func init() {
	SchemeBuilder.Register(&BackendTLSPolicy{}, &BackendTLSPolicyList{})
}
