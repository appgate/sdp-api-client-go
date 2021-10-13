#!/usr/bin/env bash

set -eou pipefail


supportedVersions=(16)

for version in "${supportedVersions[@]}"; do
    git clone git@github.com:appgate/sdp-api-specification.git --single-branch --branch "version-$version" "spec/spec/v$version"
    cd spec
    apigentools generate
    cd ..
    mkdir -p "api/v$version/openapi"
    mv spec/generated/sdp-api-client-go/openapi/*.go "api/v$version/openapi"
done

