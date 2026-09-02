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
	ResourceCodeDatabaseAccessPolicy     = "dbaccess"
	ResourceKindDatabaseAccessPolicy     = "DatabaseAccessPolicy"
	ResourceSingularDatabaseAccessPolicy = "databaseaccesspolicy"
	ResourcePluralDatabaseAccessPolicy   = "databaseaccesspolicies"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=gateway-api
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:path=databaseaccesspolicies,singular=databaseaccesspolicy,shortName=dbaccess,categories={policy,appscode}
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// DatabaseAccessPolicy carries CEL rules evaluated on the REQUEST leg, before a
// statement reaches the database. This is where DENY lives.
//
// The request-side third of the DAM configuration. Masking and PII detection
// are response-side and belong to DatabaseMaskingPolicy; audit sinks belong to
// DatabaseAuditPolicy. The split follows the three filters that consume the
// config, and the three have genuinely different lifecycles.
type DatabaseAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of DatabaseAccessPolicy.
	Spec DatabaseAccessPolicySpec `json:"spec"`

	// Status defines the current state of DatabaseAccessPolicy.
	Status DatabaseAccessPolicyStatus `json:"status,omitempty"`
}

// DatabaseAccessPolicySpec defines the desired state of DatabaseAccessPolicy.
type DatabaseAccessPolicySpec struct {
	// TargetRefs identifies the database routes this policy attaches to.
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=16
	TargetRefs []gwv1.LocalPolicyTargetReferenceWithSectionName `json:"targetRefs"`

	// Rules form an ALLOWLIST. A statement matching NO rule is denied.
	//
	// Every rule is evaluated -- there is no short-circuit -- and the outcome
	// is decided by action rather than by position:
	//
	//	any DENY match  -> denied
	//	any other match -> permitted
	//	no match at all -> denied, implicitly
	//
	// Only DENY refuses. ALLOW permits silently, LOG permits and records,
	// ALERT permits and raises -- so a rule set of nothing but LOG rules is a
	// valid allowlist: it permits what it logs and the implicit rule refuses
	// the rest.
	//
	// Deciding by action means a carve-out cannot be defeated by where someone
	// put it in the list.
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:items:XValidation:rule="self.action != 'DENY' || has(self.severity)",message="severity is required for DENY rules"
	// +listType=atomic
	Rules []AccessRule `json:"rules"`

	// SimulationMode evaluates every rule and reports what would have happened
	// without enforcing DENY. Use it to introduce a rule against live traffic
	// before it can break anything.
	//
	// +optional
	SimulationMode bool `json:"simulationMode,omitempty"`

	// EnvFields are static values merged into `request.env` before evaluation,
	// for facts the data plane cannot know -- cluster name, deployment tier.
	//
	// +optional
	EnvFields map[string]string `json:"envFields,omitempty"`
}

// RuleAction is what happens when a rule matches.
//
// Only DENY refuses. Every other action permits.
//
// +kubebuilder:validation:Enum=LOG;ALERT;DENY;ALLOW
type RuleAction string

const (
	// RuleActionLog permits the statement and records the match.
	RuleActionLog RuleAction = "LOG"
	// RuleActionAlert permits the statement and raises severity on the audit
	// record.
	RuleActionAlert RuleAction = "ALERT"
	// RuleActionDeny closes the connection. Beats a matching ALLOW wherever
	// either sits in the list.
	RuleActionDeny RuleAction = "DENY"
	// RuleActionAllow permits the statement silently. LOG and ALERT permit
	// too; ALLOW is the one that records nothing, for traffic that is simply
	// expected.
	RuleActionAllow RuleAction = "ALLOW"
)

// RuleSeverity labels a match for triage.
// +kubebuilder:validation:Enum=INFO;LOW;MEDIUM;HIGH;CRITICAL
type RuleSeverity string

const (
	RuleSeverityInfo     RuleSeverity = "INFO"
	RuleSeverityLow      RuleSeverity = "LOW"
	RuleSeverityMedium   RuleSeverity = "MEDIUM"
	RuleSeverityHigh     RuleSeverity = "HIGH"
	RuleSeverityCritical RuleSeverity = "CRITICAL"
)

// AccessRule is one CEL rule evaluated per statement.
type AccessRule struct {
	// ID is stable and must be unique across every policy attached to the same
	// route. It appears in audit records and in status conditions.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	ID string `json:"id"`

	// Description is free text for humans; it does not affect evaluation.
	//
	// +optional
	Description string `json:"description,omitempty"`

	// When is a CEL expression returning bool, evaluated against the statement
	// envelope. Examples:
	//
	//   request.query.operation == "UPDATE"
	//   "dba" in request.identity.roles
	//   request.query.risk_score >= 80
	//
	// NOTE THE CASING, and note that this example previously read "update".
	// Postgres, MySQL and MSSQL all stamp query.operation in UPPER CASE;
	// MongoDB stamps the driver's own command names, which are lower and camel
	// case ("update", "findAndModify"). A rule with the wrong casing compiles,
	// loads, evaluates on every statement and never matches -- measured: a
	// lowercase rule let a real UPDATE straight through a DENY.
	//
	// A field path that does not exist evaluates false and the rule silently
	// never matches, so an expression referencing a misspelled field is
	// indistinguishable from one that simply never fires. Measured against the
	// filter rather than assumed: a DENY rule naming a misspelled field left
	// eval_errors, denies and rule_matches all at zero while evaluations
	// counted 1, and the statement went through. Nothing in the data plane
	// reports it, so the mistake has to be caught where the rule is written.
	//
	// +kubebuilder:validation:MinLength=1
	When string `json:"when"`

	// Action taken when When is true.
	Action RuleAction `json:"action"`

	// Severity attached to the match. Required for DENY.
	//
	// +optional
	Severity RuleSeverity `json:"severity,omitempty"`

	// Reason is a CEL expression returning a string, evaluated only on a match,
	// and recorded on the audit event. Note it is CEL, not literal text: a bare
	// word is an identifier and will fail to resolve, so a constant must be
	// quoted -- `"'dev may not write'"`.
	//
	// +optional
	Reason string `json:"reason,omitempty"`

	// Priority orders rules when several policies attach to one route: higher
	// runs first, ties broken by namespace/name. It is a control-plane concern
	// only -- the filter has no priority field, and position in the emitted
	// list IS evaluation order.
	//
	// +optional
	// +kubebuilder:default=0
	Priority int32 `json:"priority,omitempty"`
}

// DatabaseAccessPolicyStatus defines the observed state of DatabaseAccessPolicy.
type DatabaseAccessPolicyStatus struct {
	gwv1.PolicyStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// DatabaseAccessPolicyList contains a list of DatabaseAccessPolicy.
type DatabaseAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseAccessPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseAccessPolicy{}, &DatabaseAccessPolicyList{})
}

func (r *DatabaseAccessPolicy) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(GroupVersion.WithResource(ResourcePluralDatabaseAccessPolicy))
}
