[![CI](https://github.com/linode/provider-linode/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/linode/provider-linode/actions/workflows/ci.yml)

# Provider Linode

**###############################################################################################################**

**NOTE: THE LINODE CROSSPLANE PROVIDER IS CURRENTLY CONSIDERED BETA. PLEASE TRY IT AND LET US KNOW HOW IT WORKS FOR YOU!**

**###############################################################################################################**

- Crossplane Website: <https://www.crossplane.io>
- Linode Website: <https://www.linode.com>
- Documentation: <https://marketplace.upbound.io/providers/linode/provider-linode>

## Abount Crossplane

<img src="https://github.com/crossplane/artwork/blob/420102818bc649730cc97de5b4ed8178e9333eb5/logo/icon.svg" height="200px" width="300px">

Crossplane is an open source Kubernetes add-on that transforms your cluster into a universal control plane. Crossplane enables platform teams to assemble infrastructure from multiple vendors, and expose higher level self-service APIs for application teams to consume, without having to write any code.

Crossplane extends your Kubernetes cluster to support orchestrating any infrastructure or managed service. Compose Crossplane’s granular resources into higher level abstractions that can be versioned, managed, deployed and consumed using your favorite tools and existing processes. Install Crossplane into any Kubernetes cluster to get started.

Crossplane is a Cloud Native Compute Foundation project.

## About Linode / Akamai Cloud Computing

<img src="https://www.linode.com/wp-content/themes/linode-website-theme/images/linode-akamai-logo.svg?ver=1663187393" height="200px" width="300px">

Akamai Cloud Computing based on Linode accelerates innovation with scalable, simple, affordable, and accessible Linux cloud solutions and services. Our products, services, and people give developers and enterprises the flexibility, support, and trust they need to build, deploy, secure, and scale applications more easily and cost-effectively from cloud to edge on the world’s most distributed network.

## Maintainers

This provider plugin is maintained by Linode.

## Requirements

- [Crossplane](https://docs.crossplane.io/v1.10/getting-started/install-configure/) 0.10.0+
- [Go](https://golang.org/doc/install) 1.18.0 or higher (to build the provider plugin)

## Using the provider

`provider-linode` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Linode API.

## Getting Started

Install the provider by following the [getting started guide](docs/install.md)

## Developing

Build and develop the provider locally by following the [developer guide](docs/developer.md)

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/linode/provider-linode/issues).
