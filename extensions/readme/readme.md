# Provider Supabase

`provider-supabase` is a [Crossplane](https://crossplane.io/) provider built using [Upjet](https://github.com/crossplane/upjet) that exposes XRM-conformant managed resources for the [Supabase Management API](https://supabase.com/docs/reference/api/introduction).

This provider is distributed as a **provider family**. Install only the sub-provider you need; it automatically pulls in the shared `ProviderConfig` via `provider-family-supabase`.

| Package | Resources | CRDs |
|---------|-----------|------|
| `provider-family-supabase` | ProviderConfig (auto-installed) | 5 |
| `provider-supabase-project` | Projects, Settings, Branches, Edge Functions, Edge Function Secrets, API Keys | 12 |

## Install

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-supabase-project
spec:
  package: xpkg.upbound.io/wildbitca/provider-supabase-project:v0.1.0
```

## ProviderConfig

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: supabase-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "access_token": "YOUR_SUPABASE_ACCESS_TOKEN"
    }
---
apiVersion: upjet-supabase.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: supabase-creds
      namespace: crossplane-system
      key: credentials
```

## Report a Bug

Open an [issue](https://github.com/wildbitca/provider-upjet-supabase/issues).
