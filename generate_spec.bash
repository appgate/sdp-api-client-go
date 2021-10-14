#!/usr/bin/env bash

set -eou pipefail

if ! command -v apigentools &> /dev/null; then
    echo "apigentools could not be found"
    echo "Get it - https://github.com/DataDog/apigentools#get-it"
    exit 1
fi


# Starting from v16 we will use apigentools to generate the sdk with openapi-generator,
# older versions are not supported by apigentools.
supportedVersions=(16)



for version in "${supportedVersions[@]}"; do
    # remove any existing spec files and download a fresh copy again
    rm -rf "spec/spec/v$version"
done

for version in "${supportedVersions[@]}"; do
    git clone git@github.com:appgate/sdp-api-specification.git --depth 1 --single-branch --branch "version-$version" "spec/spec/v$version"
    cd spec
    apigentools generate
    cd ..
    mkdir -p "api/v$version/openapi"
    mv spec/generated/sdp-api-client-go/openapi/*.go "api/v$version/openapi"
done

