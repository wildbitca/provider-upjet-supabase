# Changelog

## Unreleased

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
