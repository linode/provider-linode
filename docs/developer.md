## Developing

Run code-generation pipeline - you need to do this everytime you change the linode-terraform-provider version, to update CRDs
```console
make generate
```

Run against a Kubernetes cluster:

```console
export KUBECONFIG=<kubeconfig>
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

### Add a new resource:

1. Find the resource (not data source) that was added in [Terraform Provider for Linode](https://github.com/linode/terraform-provider-linode). Newer resources are usually implemented through the `framework` plugin, and you can find them in `linode/framework_provider.go` under `Resources`. Older resources may be implemented with `SDKv2`.
2. In this repo, add the name of the provider to the appropriate list in `config/externalname.go`. If `framework` based, this will be in the `terraformPluginFrameworkExternalNameConfigs` list; make sure there is a 1:1 mapping with the list from `linode/framework_provider.go`.
3. Create a directory for the new resource group in `config/`.
4. Add a file `config.go` implementing the `Configure` method (see config/vpc/config.go for an example). The minimum change needed is overriding the `ShortGroup`.
5. Add the configure method to the loop at `config/provider.go:GetProvider()`  method.
6. Change the terraform provider version in the `Makefile` and the `go.mod` file. 
7. Run `make submodules` and `make generate`.
8. Check the new CRDs created in `examples-generated/`. If you need to make any changes, go back to step 4 and edit the config, then `make generate` again.
9. Run `make local-deploy`. This will create a Kind cluster set up with Crossplane and Provider Linode.
10. Create a ProviderConfig and a secret with your Linode API token
```
apiVersion: linode.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: linode-api-token
      namespace: default
      key: credentials
```
The Linode API token secret needs to be in the form `echo '{"token": "$LINODE_API_TOKEN", "api_version": "v4beta"}' | base64`. Take that value and input it here:
```
apiVersion: v1
data:
  credentials: PUT THE SECRET HERE
kind: Secret
metadata:
  name: linode-api-token
  namespace: default
type: Opaque
```
11. Create a manifest for your resource, or use the generated example depending on what you are testing. `k apply` the new resource, the ProviderConfig, and the linode API token secret. Test that it works!

### Upgrading to a new terraform provider:
1. Review the terraform-provider-linode release notes, and verify if there are is any migration from sdk->plugin framework. k ge
    a. If there's a migration, update the list in `config/externalname.go`
    b. If there's a new resource follow steps under Add a new resource

2. Change `TERRAFORM_PROVIDER_VERSION` in makefile
3. Change `github.com/linode/terraform-provider-linode/v2` in go.mod
4. Run `go mod tidy`
4. Run `make generate`
5. Test changes locally with `make local-deploy` - create a manifest for your resource and test that it works.

### Debugging using dlv:
1. use `make run` to run the provider locally
2. `dlv attach <pid>` to begin a debug session


### Creating and publishing xpkgs
1. Create an account on the upbound marketplace
2. Install UP from - https://docs.upbound.io/reference/cli/#homebrew 
3. login to the account using `up login` (if you're already logged in, logout first)
4. Create a repo if you don't have one `up repository create provider-linode`
5. run `make build`
6. run `make xpkg.build`
7. run `up xpkg push <username>/<repo_name>:<version> -f <path to built xpkg>`
```bash
   up xpkg push tchinmai7/provider-linode:v0.0.1 -f _output/xpkg/linux_amd64/provider-linode-v0.0.0-130.g7bacec4.xpkg
``` 
the version HAS to be a semver
8. modify your provider package to use this version instead.