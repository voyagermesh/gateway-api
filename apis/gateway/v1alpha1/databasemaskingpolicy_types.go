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
	ResourceCodeDatabaseMaskingPolicy     = "dbmask"
	ResourceKindDatabaseMaskingPolicy     = "DatabaseMaskingPolicy"
	ResourceSingularDatabaseMaskingPolicy = "databasemaskingpolicy"
	ResourcePluralDatabaseMaskingPolicy   = "databasemaskingpolicies"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=gateway-api
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:path=databasemaskingpolicies,singular=databasemaskingpolicy,shortName=dbmask,categories={policy,appscode}
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// DatabaseMaskingPolicy configures response-side data protection on the engine
// filter for the database route it targets: column masking, PII detection
// patterns, and the DLP hit cap.
//
// This is the response-side third of the DAM configuration. It is deliberately
// separate from DatabaseAccessPolicy (CEL rules, evaluated on the request leg by
// kubedb_dam_policy) and DatabaseAuditPolicy (audit sinks), because the three
// map one-to-one onto the three filters that consume them and have genuinely
// different lifecycles: a masking rule follows the database schema, a CEL rule
// follows compliance policy.
type DatabaseMaskingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of DatabaseMaskingPolicy.
	Spec DatabaseMaskingPolicySpec `json:"spec"`

	// Status defines the current state of DatabaseMaskingPolicy.
	Status DatabaseMaskingPolicyStatus `json:"status,omitempty"`
}

// DatabaseMaskingPolicySpec defines the desired state of DatabaseMaskingPolicy.
type DatabaseMaskingPolicySpec struct {
	// TargetRefs identifies the database routes this policy attaches to. Each
	// ref must name a route kind in the gateway.voyagermesh.com group, in the
	// same namespace as the policy.
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=16
	TargetRefs []gwv1.LocalPolicyTargetReferenceWithSectionName `json:"targetRefs"`

	// MaskingRules rewrite column values on their way back to the client.
	//
	// Columns are matched by NAME ONLY and case-insensitively; schema-qualified
	// matching is not supported, so a rule on `card_number` applies to that
	// column in every table the connection reads.
	//
	// +optional
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:items:XValidation:rule="self.type != 'PARTIAL' || has(self.showLast)",message="showLast is required when type is PARTIAL"
	MaskingRules []MaskingRule `json:"maskingRules,omitempty"`

	// PiiPatterns are RE2 regexes evaluated against response payloads to count
	// exposed sensitive values. Detection only -- patterns never block.
	//
	// Scanning happens BEFORE masking, so a column that is both masked and
	// matched by a pattern is still counted. That is deliberate: the counters
	// are compliance evidence of what was present, and whether a masked column
	// should also alert is a question for DatabaseAccessPolicy.
	//
	// +optional
	// +kubebuilder:validation:MaxItems=32
	PiiPatterns []PiiPattern `json:"piiPatterns,omitempty"`

	// DlpMaxHitsPerPattern caps how many times a single pattern is counted per
	// field, bounding the cost of a pathological payload. 0 means the filter
	// default (64).
	//
	// +optional
	// +kubebuilder:validation:Maximum=65535
	DlpMaxHitsPerPattern *uint32 `json:"dlpMaxHitsPerPattern,omitempty"`
}

// MaskingType is the transform applied to a matched column value.
// +kubebuilder:validation:Enum=FULL;PARTIAL;HASH
type MaskingType string

const (
	// MaskingTypeFull replaces every byte of the value with maskChar.
	// Length-preserving.
	MaskingTypeFull MaskingType = "FULL"
	// MaskingTypePartial replaces every byte except the last showLast.
	// Length-preserving. If showLast >= the value's length the rule is skipped
	// for that value rather than revealing more than the original.
	MaskingTypePartial MaskingType = "PARTIAL"
	// MaskingTypeHash replaces the value with the 64-character lower-case hex
	// SHA-256 of the original bytes -- a stable surrogate that supports joins
	// and dedup without exposing the value. Length-CHANGING.
	//
	// Engine support is not uniform: only Postgres implements HASH. On MySQL,
	// MSSQL and MongoDB it degrades to FULL, and Oracle has no masking field at
	// all. The controller reports that divergence on the policy's status rather
	// than dropping the rule silently.
	MaskingTypeHash MaskingType = "HASH"
)

// MaskingRule is one column-masking rule.
type MaskingRule struct {
	// ID is a stable identifier for this rule, unique within the policy. It
	// appears in audit records and in status conditions.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	ID string `json:"id"`

	// Columns this rule applies to, matched by name, case-insensitively.
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=64
	Columns []string `json:"columns"`

	// Type of transform to apply.
	Type MaskingType `json:"type"`

	// ShowLast is the number of trailing characters left visible by a PARTIAL
	// mask. Counted in BYTES, not codepoints, so a value that splits a
	// multi-byte character will not round-trip as valid UTF-8.
	//
	// +optional
	// +kubebuilder:validation:Maximum=64
	ShowLast *uint32 `json:"showLast,omitempty"`

	// MaskChar is the single-byte replacement character used by FULL and
	// PARTIAL. Defaults to "*"; ignored by HASH.
	//
	// +optional
	// +kubebuilder:validation:MaxLength=1
	MaskChar string `json:"maskChar,omitempty"`

	// ExemptRoles are identity roles that see the column unmasked. Matched
	// against request.identity.roles, which is populated from the client
	// certificate presented on the connection.
	//
	// A connection with no verified client certificate carries no roles, so no
	// rule is exempt and every rule applies -- the safe default.
	//
	// +optional
	// +kubebuilder:validation:MaxItems=32
	ExemptRoles []string `json:"exemptRoles,omitempty"`
}

// PiiPattern is a named RE2 regex used for response-side PII detection.
type PiiPattern struct {
	// Name labels matches of this pattern in metrics and audit records,
	// e.g. "ssn", "credit_card".
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=63
	Name string `json:"name"`

	// Regex is an RE2 expression, partial-matched against field values.
	//
	// +kubebuilder:validation:MinLength=1
	Regex string `json:"regex"`
}

// DatabaseMaskingPolicyStatus defines the observed state of DatabaseMaskingPolicy.
type DatabaseMaskingPolicyStatus struct {
	gwv1.PolicyStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// DatabaseMaskingPolicyList contains a list of DatabaseMaskingPolicy.
type DatabaseMaskingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseMaskingPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseMaskingPolicy{}, &DatabaseMaskingPolicyList{})
}

func (r *DatabaseMaskingPolicy) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(GroupVersion.WithResource(ResourcePluralDatabaseMaskingPolicy))
}
