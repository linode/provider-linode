# Provider Linode

**###############################################################################################################**

**NOTE: THIS IS A PRE-RELEASE REPO AND SHOULD NOT BE DEPLOYED TO PRODUCTION AT THIS TIME UNTIL OFFICIALLY RELEASED**

**###############################################################################################################**



`provider-linode` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Linode API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/linode/provider-linode):
```
up ctp provider install linode/provider-linode:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-linode
spec:
  package: linode/provider-linode:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/linode/provider-linode).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
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

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/linode/provider-linode/issues).
