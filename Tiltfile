allow_k8s_contexts(k8s_context())

templated_yaml = local("find package/crds -name '*.yaml' | xargs cat")

k8s_yaml(templated_yaml)

local_resource(
    "crossplane-provider-linode",
    serve_cmd="make run",
    deps=[
        "apis/",
        "go.mod",
        "go.sum",
        "cmd/",
        "config/",
        "internal"
    ],
    ignore=[
        "build/",
        "_output/",
    ],
)
