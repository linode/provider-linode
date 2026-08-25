#!/usr/bin/env bash
set -euo pipefail

show_help() {
  cat <<'EOF'
Run a provider-linode smoke test against the current Kubernetes cluster.

Creates a Placement Group and Linode Instance through Crossplane, waits for Ready, then deletes them.

Environment:
  LINODE_API_TOKEN              Linode API token. Prompted when unset.
  SMOKE_TEST_NAME               Kubernetes resource and Linode label prefix.
  SMOKE_TEST_REGION             Linode region. Default: us-central.
  SMOKE_TEST_TYPE               Linode type. Default: g6-nanode-1.
  SMOKE_TEST_IMAGE              Linode image. Default: linode/ubuntu22.04.
  SMOKE_TEST_TIMEOUT            Ready/delete timeout. Default: 20m.
  SMOKE_TEST_KEEP_RESOURCES     Set to true to skip cleanup.
  KUBECTL                       kubectl path. Default: kubectl.
  CROSSPLANE_NAMESPACE          Secret namespace. Default: upbound-system.
EOF
}

case "${1:-}" in
-h | --help)
  show_help
  exit 0
  ;;
"")
  ;;
*)
  echo "unknown argument: $1" >&2
  show_help >&2
  exit 2
  ;;
esac

KUBECTL=${KUBECTL:-kubectl}
CROSSPLANE_NAMESPACE=${CROSSPLANE_NAMESPACE:-upbound-system}
SMOKE_TEST_NAME=${SMOKE_TEST_NAME:-provider-linode-smoke-$(date +%s)}
SMOKE_TEST_REGION=${SMOKE_TEST_REGION:-us-central}
SMOKE_TEST_TYPE=${SMOKE_TEST_TYPE:-g6-nanode-1}
SMOKE_TEST_IMAGE=${SMOKE_TEST_IMAGE:-linode/ubuntu22.04}
SMOKE_TEST_TIMEOUT=${SMOKE_TEST_TIMEOUT:-20m}
SMOKE_TEST_KEEP_RESOURCES=${SMOKE_TEST_KEEP_RESOURCES:-false}

TOKEN_SECRET_NAME=${SMOKE_TEST_NAME}-token
ROOT_PASSWORD_SECRET_NAME=${SMOKE_TEST_NAME}-root-password
ROOT_PASSWORD="Crossplane1!$(od -An -N12 -tx1 /dev/urandom | tr -d ' \n')"

if [[ -z "${LINODE_API_TOKEN:-}" ]]; then
  if [[ ! -t 0 ]]; then
    echo "LINODE_API_TOKEN is required when stdin is not interactive" >&2
    exit 1
  fi
  read -r -s -p "Linode API token: " LINODE_API_TOKEN
  printf '\n'
fi

if [[ -z "${LINODE_API_TOKEN}" ]]; then
  echo "Linode API token cannot be empty" >&2
  exit 1
fi

cleanup() {
  if [[ "${SMOKE_TEST_KEEP_RESOURCES}" == "true" || "${SMOKE_TEST_KEEP_RESOURCES}" == "1" ]]; then
    echo "Keeping smoke test resources. Delete ${SMOKE_TEST_NAME} manually to avoid ongoing charges."
    return
  fi

  local cleanup_status=0
  echo "Cleaning up smoke test resources..."
  "${KUBECTL}" delete instances.instance.linode.upbound.io "${SMOKE_TEST_NAME}" --ignore-not-found --wait=true --timeout="${SMOKE_TEST_TIMEOUT}" || cleanup_status=$?
  "${KUBECTL}" delete placementgroups.placementgroup.linode.upbound.io "${SMOKE_TEST_NAME}" --ignore-not-found --wait=true --timeout="${SMOKE_TEST_TIMEOUT}" || cleanup_status=$?
  "${KUBECTL}" delete providerconfigs.linode.upbound.io "${SMOKE_TEST_NAME}" --ignore-not-found || cleanup_status=$?
  "${KUBECTL}" -n "${CROSSPLANE_NAMESPACE}" delete secret "${TOKEN_SECRET_NAME}" "${ROOT_PASSWORD_SECRET_NAME}" --ignore-not-found || cleanup_status=$?
  return "${cleanup_status}"
}

finish() {
  local test_status=$?
  local cleanup_status=0
  cleanup || cleanup_status=$?
  if [[ "${test_status}" -ne 0 ]]; then
    exit "${test_status}"
  fi
  exit "${cleanup_status}"
}

trap finish EXIT

echo "Creating Linode provider credentials secret..."
"${KUBECTL}" -n "${CROSSPLANE_NAMESPACE}" create secret generic "${TOKEN_SECRET_NAME}" \
  --from-file=credentials=/dev/stdin \
  --dry-run=client \
  -o yaml <<<"{\"token\":\"${LINODE_API_TOKEN}\"}" | "${KUBECTL}" apply -f -

echo "Creating root password secret..."
printf '%s' "${ROOT_PASSWORD}" | "${KUBECTL}" -n "${CROSSPLANE_NAMESPACE}" create secret generic "${ROOT_PASSWORD_SECRET_NAME}" \
  --from-file=password=/dev/stdin \
  --dry-run=client \
  -o yaml | "${KUBECTL}" apply -f -

echo "Creating ProviderConfig and Instance manifests..."
cat <<EOF | "${KUBECTL}" apply -f -
apiVersion: linode.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: ${SMOKE_TEST_NAME}
spec:
  credentials:
    source: Secret
    secretRef:
      name: ${TOKEN_SECRET_NAME}
      namespace: ${CROSSPLANE_NAMESPACE}
      key: credentials
---
apiVersion: placementgroup.linode.upbound.io/v1alpha1
kind: PlacementGroup
metadata:
  name: ${SMOKE_TEST_NAME}
spec:
  providerConfigRef:
    name: ${SMOKE_TEST_NAME}
  forProvider:
    label: ${SMOKE_TEST_NAME}
    placementGroupType: "anti_affinity:local"
    region: ${SMOKE_TEST_REGION}
    placementGroupPolicy: "flexible"
---
apiVersion: instance.linode.upbound.io/v1alpha1
kind: Instance
metadata:
  name: ${SMOKE_TEST_NAME}
  labels:
    testing.upbound.io/example-name: ${SMOKE_TEST_NAME}
spec:
  providerConfigRef:
    name: ${SMOKE_TEST_NAME}
  forProvider:
    image: ${SMOKE_TEST_IMAGE}
    label: ${SMOKE_TEST_NAME}
    region: ${SMOKE_TEST_REGION}
    rootPassSecretRef:
      key: password
      name: ${ROOT_PASSWORD_SECRET_NAME}
      namespace: ${CROSSPLANE_NAMESPACE}
    swapSize: 256
    type: ${SMOKE_TEST_TYPE}
    placementGroup:
      - idRef:
          name: ${SMOKE_TEST_NAME}
EOF

echo "Waiting for Instance ${SMOKE_TEST_NAME} to become Ready..."
"${KUBECTL}" wait instances.instance.linode.upbound.io "${SMOKE_TEST_NAME}" \
  --for=condition=Ready \
  --timeout="${SMOKE_TEST_TIMEOUT}"

echo "Smoke test passed: Instance ${SMOKE_TEST_NAME} is Ready."
