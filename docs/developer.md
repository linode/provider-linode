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

1. Find out if the resource is implemented using the `frameworkprovider` or the terraform v2 provider in the linode-terraform-provider - newer resources are usually frameworkprovider based
2. Add the name of the provider to the appropriate list in `config/externalname.go`
3. Create a directory for the resource group in config/
4. Add a file `config.go` implementing the `Configure` method (see config/vpc/vpc.go for an example)
5. Add the configure method to the loop at `config/provider.go:GetProvider()`  method
6. Run `make generate`, `make local-deploy` - create a manifest for your resource and test that it works.

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
