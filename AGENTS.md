# AGENTS.md

Instructions for AI coding agents working in this repository. Keep advice
agent-neutral; other harnesses read this file directly.

## What this repo is

The **VoyagerMesh/AppsCode** Gateway API extension types for KubeDB database
routing. It defines the database-protocol Route CRDs that upstream Gateway API
does not have, all in `apis/gateway/v1alpha1`:

```
MySQLRoute   PostgresRoute   MongoDBRoute   RedisRoute
KafkaRoute   OracleRoute     MSSQLServer
```

Each follows the same shape: `CommonRouteSpec` + `Hostnames` (SNI matching) +
`Rules []<Engine>RouteRule` with `BackendRefs`. `voyagermesh/gateway` consumes
these and translates them into Envoy listeners.

## Where new configuration goes

**Do not add filter configuration as bare fields on a Route spec.** `MySQLRouteSpec`
already carries `EnableDecoding bool` from before this rule existed; it is the
pattern to stop repeating, not a precedent to follow.

That shape does not scale. KubeDB DAM has roughly fifteen configurable fields
across seven engines, which would mean seven near-identical CRD surfaces to keep
in sync — and it has no home at all for the chain-level filters
(`kubedb_dam_policy`, `kubedb_dam_audit`) that are not protocol settings.

New configuration should instead attach as a **policy**, following the Gateway
API policy-attachment pattern that Envoy Gateway already implements
(`BackendTrafficPolicy`, `SecurityPolicy`, `ClientTrafficPolicy`,
`EnvoyExtensionPolicy` — all in the gateway fork's `api/v1alpha1/`). Embedding
`PolicyTargetReferences` gives you `targetRefs`, `targetSelectors` for label-based
fan-out, `MergeType` for gateway-wide defaults with per-database overrides, and a
status block reporting whether the policy actually attached — none of which a
bare route field can express.

## After changing types

Regenerate deepcopy and CRD manifests before committing, and keep the generated
output in the same commit as the type change.

## House rules

- Commit with `git commit -s`.
- Do not add Claude/AI authorship attribution to commits or PR bodies.
- Do not `git push`; the maintainer handles all pushes.
