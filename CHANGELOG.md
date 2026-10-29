# Changelog

## 1.0.0 - 2026-10-29

First stable release. All Wildbit providers move to 1.x together.

### Breaking changes

- Upjet moved from the `wildbitca/upjet` fork to upstream `crossplane/upjet` v2.5.1 and crossplane-runtime to v2.4.0. Go consumers importing `apis/*/v1beta1` must switch to the `ProviderConfig` / `ProviderConfigUsage` types backed by `github.com/crossplane/crossplane/apis/v2/core/v2`.
- All CRDs were regenerated with controller-tools v0.20.0 (new `status.lastHandledReconcileAt`); managed-resource schemas for the cluster and namespaced scopes changed accordingly.
- Requires a Crossplane v2 control plane.

### Migration

- Upgrade Crossplane to v2 first, then update the `Provider` packages (family and every installed sub-provider) to `v1.0.0` together so the shared `ProviderConfig` CRDs stay consistent.
- Re-apply any manifests that pin CRD fields from the previous version; existing managed resources keep their external names and need no manual changes.
- Review `legacy_api_keys_enabled` on `Project`: it is now flagged deprecated by the Terraform provider.

### Changed

- Upjet moved from the `wildbitca/upjet` fork (v2.2.2) to upstream `crossplane/upjet` v2.5.1; the `replace` directive is gone.
- crossplane-runtime v2.1.0 -> v2.4.0; common APIs now come from `github.com/crossplane/crossplane/apis/v2/core/v2` (`ProviderConfig` and `ProviderConfigUsage` types in `apis/*/v1beta1`).
- Go 1.26.8, controller-tools v0.20.0, Kubernetes libraries v0.35, controller-runtime v0.23.
- OpenTofu 1.11.5 -> 1.13.2 (Makefile, publish workflow, schema regeneration).
- All CRDs regenerated: new `status.lastHandledReconcileAt` field; `controller-gen` annotation bumped to v0.20.0.
- `config/schema.json` regenerated with OpenTofu 1.13.2; `supabase_project.legacy_api_keys_enabled` now carries a deprecation message.

### Added

- Conversion webhook registration (`SetupWebhookWithManager`) for `ProviderConfig` and generated resources, wired in the provider entrypoint when `--certs-dir` is set.
- Unit test covering provider construction for both scopes.

### Notes

- The Supabase Terraform provider stays at v1.11.0 (already the latest release): no new, removed or renamed resources or data sources.
- The `patches/apply.sh` vendor patch (empty Terraform ID guard) is kept; upstream v2.5.1 still lacks the guard.
