#!/usr/bin/env bash

set -eou pipefail

if ! command -v apigentools &> /dev/null; then
    echo "apigentools could not be found"
    echo "Get it - https://github.com/DataDog/apigentools#get-it"
    exit 1
fi


# Starting from v16 we will use apigentools to generate the sdk with openapi-generator,
# older versions are not supported by apigentools.
supportedVersions=(16 17)



for version in "${supportedVersions[@]}"; do
    # remove any existing spec files and download a fresh copy again
    rm -rf "spec/spec/v$version"
done

for version in "${supportedVersions[@]}"; do
    echo "Cloning $version to spec/spec/v$version"
    git clone git@github.com:appgate/sdp-api-specification.git --depth 1 --single-branch --branch "version-$version" "spec/spec/v$version"
done

for version in "${supportedVersions[@]}"; do
    # apply patches if there are any
    if [ -d "$PWD/spec/spec-patches/v$version" ]; then
        echo "Apply custom patches for openapi v$version"
        find "$PWD/spec/spec-patches/v$version" -name "*.patch" -print0 | while read -r -d $'\0' patch; do
            echo "Applying $patch"
            patch --fuzz 0 --no-backup-if-mismatch -p1 -i "$patch" -d "$PWD/spec/spec/v$version"
        done
    fi

done

cd spec
apigentools validate
apigentools generate

for version in "${supportedVersions[@]}"; do
    mkdir -p "../api/v${version}/openapi/"
    find "generated/sdp-api-client-go/openapi_v${version}" -name '*.go' -exec cp {} "../api/v$version/openapi/" \;
done
