# Provider Supabase

`provider-supabase` is a [Crossplane](https://crossplane.io/) provider built using [Upjet](https://github.com/crossplane/upjet) that exposes XRM-conformant managed resources for the [Supabase Management API](https://supabase.com/docs/reference/api/introduction), wrapping the [Supabase Terraform Provider](https://registry.terraform.io/providers/supabase/supabase/latest/docs).

## Getting Started

This provider uses the **family** distribution model: install only the sub-providers you need. The first sub-provider you install automatically pulls in `provider-family-supabase`, which manages the shared `ProviderConfig`.

> For provider families background, see [Scalable Provider Families](https://blog.crossplane.io/crd-scaling-provider-families/) and the [Upbound Provider Families docs](https://docs.upbound.io/manuals/packages/providers/provider-families/).

### Install

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-supabase-project
spec:
  package: xpkg.upbound.io/wildbitca/provider-supabase-project:v0.1.0
```

### Available sub-providers

| Sub-provider | Resources |
|-------------|-----------|
| `provider-supabase-project` | Projects, settings, branches, edge functions, edge function secrets, API keys |

### ProviderConfig

Create a Secret with Supabase credentials. The Supabase Terraform provider authenticates with a Management API **access token** (personal access token) and, optionally, a custom API `endpoint`:

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
```

Create a ProviderConfig:

```yaml
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

### Rate limiting

For large deployments, use a `DeploymentRuntimeConfig` to throttle:

```yaml
apiVersion: pkg.crossplane.io/v1beta1
kind: DeploymentRuntimeConfig
metadata:
  name: supabase-throttled
spec:
  deploymentTemplate:
    spec:
      selector: {}
      template:
        spec:
          containers:
            - name: package-runtime
              args:
                - --poll=30m
                - --max-reconcile-rate=1
                - --sync=4h
```

## Developing

```bash
make submodules
go install golang.org/x/tools/cmd/goimports@latest
export PATH="$(go env GOPATH)/bin:$PATH"
make generate
```

Run against a Kubernetes cluster:

```bash
make run
```

Build family sub-providers:

```bash
make build.family
make build.family FAMILY_SUBPACKAGES="config project"
```

## Supported resources

This provider exposes 6 managed resources from the [Supabase Terraform Provider v1.9.1](https://registry.terraform.io/providers/supabase/supabase/1.9.1/docs), organized in a single `project` API group:

- **Project**: `supabase_project`, `supabase_settings`, `supabase_branch`, `supabase_edge_function`, `supabase_edge_function_secrets`, `supabase_apikey`

Kinds (group `project.upjet-supabase.upbound.io` for cluster scope, `project.upjet-supabase.m.upbound.io` for namespaced):

| Kind | Terraform resource |
|------|--------------------|
| `Project` | `supabase_project` |
| `Settings` | `supabase_settings` |
| `Branch` | `supabase_branch` |
| `Function` | `supabase_edge_function` |
| `FunctionSecrets` | `supabase_edge_function_secrets` |
| `Apikey` | `supabase_apikey` |

## Report a Bug

Open an [issue](https://github.com/wildbitca/provider-upjet-supabase/issues).
