/*
Copyright 2023.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/client-go/apiextensions"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
	crds "voyagermesh.dev/gateway-api/config/crd/bases"
)

const (
	ResourceCodeDatabaseAuditPolicy     = "dbaudit"
	ResourceKindDatabaseAuditPolicy     = "DatabaseAuditPolicy"
	ResourceSingularDatabaseAuditPolicy = "databaseauditpolicy"
	ResourcePluralDatabaseAuditPolicy   = "databaseauditpolicies"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=gateway-api
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:path=databaseauditpolicies,singular=databaseauditpolicy,shortName=dbaudit,categories={policy,appscode}
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// DatabaseAuditPolicy configures what the audit filter records for a database
// route and where it writes it.
//
// The third of the DAM split. Masking is response-side, CEL rules are
// request-side, and this is the evidence surface -- the record that says what
// happened, which is what a compliance claim is ultimately made of.
type DatabaseAuditPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of DatabaseAuditPolicy.
	Spec DatabaseAuditPolicySpec `json:"spec"`

	// Status defines the current state of DatabaseAuditPolicy.
	Status DatabaseAuditPolicyStatus `json:"status,omitempty"`
}

// DatabaseAuditPolicySpec defines the desired state of DatabaseAuditPolicy.
type DatabaseAuditPolicySpec struct {
	// TargetRefs identifies the database routes this policy attaches to.
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=16
	TargetRefs []gwv1.LocalPolicyTargetReferenceWithSectionName `json:"targetRefs"`

	// Sinks are where audit records are written. At least one is required --
	// an audit policy with nowhere to write is configuration that looks like
	// coverage and produces nothing.
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=8
	Sinks []AuditSink `json:"sinks"`

	// MinRiskScore drops records below this score. 0 records everything.
	//
	// Raising it is a deliberate trade: it reduces volume by discarding the
	// low-risk statements that establish what normal looks like, which is what
	// an investigation needs in order to say a given statement was abnormal.
	//
	// +optional
	// +kubebuilder:validation:Maximum=100
	MinRiskScore *uint32 `json:"minRiskScore,omitempty"`

	// OnlyOnAlert records only statements that raised an alert.
	//
	// +optional
	OnlyOnAlert bool `json:"onlyOnAlert,omitempty"`

	// OnlyOnBlock records only statements that were blocked.
	//
	// Note what this excludes: a statement that was allowed through is not
	// recorded, so the log can no longer answer "who read this table". Use it
	// for a narrow enforcement feed, not as the audit trail.
	//
	// +optional
	OnlyOnBlock bool `json:"onlyOnBlock,omitempty"`
}

// AuditSinkType selects the sink implementation.
// +kubebuilder:validation:Enum=File;OTLP
type AuditSinkType string

const (
	// AuditSinkTypeFile writes one JSON record per line to a path inside the
	// proxy container.
	AuditSinkTypeFile AuditSinkType = "File"

	// AuditSinkTypeOTLP streams records to an OpenTelemetry collector over
	// gRPC, from the proxy itself. Nothing is written to disk and nothing
	// tails a file, so records reach a collector as they are produced.
	//
	// The record carries the same fields the file sink writes -- the
	// statement envelope as the log body, and ts, ns, conn_id and event_kind
	// as attributes. event_kind is not optional: the audit filter samples a
	// shared, mutable metadata namespace on every read and write pass, so a
	// record emitted when the statement has just changed still carries the
	// previous statement's response counters. A consumer that ignores
	// event_kind will report exposures that never happened.
	AuditSinkTypeOTLP AuditSinkType = "OTLP"
)

// AuditSink is one destination for audit records.
// +kubebuilder:validation:XValidation:rule="self.type != 'File' || has(self.file)",message="file is required when type is File"
// +kubebuilder:validation:XValidation:rule="self.type != 'OTLP' || has(self.otlp)",message="otlp is required when type is OTLP"
type AuditSink struct {
	// Type of sink.
	Type AuditSinkType `json:"type"`

	// OTLP configures an OpenTelemetry sink.
	//
	// +optional
	OTLP *OTLPAuditSink `json:"otlp,omitempty"`

	// File configures a file sink.
	//
	// +optional
	File *FileAuditSink `json:"file,omitempty"`
}

// OTLPAuditSink streams audit records to an OpenTelemetry collector.
type OTLPAuditSink struct {
	// Host of the collector's OTLP gRPC endpoint. A Kubernetes service name
	// is the usual answer; it is resolved by the proxy, so it must be
	// reachable from the proxy's own namespace.
	//
	// +kubebuilder:validation:MinLength=1
	Host string `json:"host"`

	// Port of the collector's OTLP gRPC endpoint.
	//
	// +kubebuilder:default=4317
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +optional
	Port int32 `json:"port,omitempty"`

	// LogName identifies this stream to the collector, so records from two
	// databases can be told apart without parsing them.
	//
	// +kubebuilder:default=kubedb-dam-audit
	// +optional
	LogName string `json:"logName,omitempty"`
}

// FileAuditSink writes audit records to a path in the proxy container.
type FileAuditSink struct {
	// Path written inside the proxy container. Use /dev/stdout to have records
	// collected by whatever already reads the container's logs, which is
	// usually what you want -- a path on the container filesystem does not
	// survive a restart and nothing ships it anywhere.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:default="/dev/stdout"
	Path string `json:"path"`
}

// DatabaseAuditPolicyStatus defines the observed state of DatabaseAuditPolicy.
type DatabaseAuditPolicyStatus struct {
	gwv1.PolicyStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// DatabaseAuditPolicyList contains a list of DatabaseAuditPolicy.
type DatabaseAuditPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseAuditPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseAuditPolicy{}, &DatabaseAuditPolicyList{})
}

func (r *DatabaseAuditPolicy) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(GroupVersion.WithResource(ResourcePluralDatabaseAuditPolicy))
}
