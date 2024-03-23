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

Add a new resource:

1. Find out if the resource is implemented using the `frameworkprovider` or the terraform v2 provider in the linode-terraform-provider - newer resources are usually frameworkprovider based
2. Add the name of the provider to the appropriate list in `config/externalname.go`
3. Create a directory for the resource group in config/
4. Add a file `config.go` implementing the `Configure` method (see config/vpc/vpc.go for an example)
5. Add the configure method to the loop at `config/provider.go:GetProvider()`  method
6. Run `make generate`, `make local-deploy` - create a manifest for your resource and test that it works.


Debugging using dlv:
1. use `make run` to run the provider locally
2. `dlv attach <pid>` to begin a debug session
