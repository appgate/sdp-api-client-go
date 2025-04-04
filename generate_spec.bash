#!/usr/bin/env bash

set -eou pipefail

if ! command -v apigentools &> /dev/null; then
    echo "apigentools could not be found"
    echo "Get it - https://github.com/DataDog/apigentools#get-it"
    exit 1
fi


go_version_minimum="1.22.2"

go_version() {
    go version | sed -nE -e 's/[^0-9.]+([0-9.]+).+/\1/p'
}

version_lt() {
    # Return true if $1 is a lower version than than $2,
    local ver1=$1
    local ver2=$2
    # Reverse sort the versions, if the 1st item != ver1 then ver1 < ver2
    if  [[ $(echo -e -n "$ver1\n$ver2\n" | sort -rV | head -n1) != "$ver1" ]]; then
        return 0
    else
        return 1
    fi
}

if version_lt "$(go_version)" "$go_version_minimum"; then
    echo "The generator requires Go >= $go_version_minimum, Go $(go_version) found." >&2
    exit 1
fi


# Starting from v16 we will use apigentools to generate the sdk with openapi-generator,
# older versions are not supported by apigentools.
supportedVersions=(18 19 20 21 22)



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
        find "$PWD/spec/spec-patches/v$version" -name "*.patch" -print0 | sort -z | while read -r -d $'\0' patch; do
            echo "Applying $patch"
            patch --fuzz 0 --no-backup-if-mismatch -p1 -i "$patch" -d "$PWD/spec/spec/v$version"
        done
    fi

done

for version in "${supportedVersions[@]}"; do
    pushd spec
    if [[ $version == 16 ]]; then
        # shellcheck disable=SC2016
        sed 's/${IMAGE}/openapitools\/openapi-generator-cli:v5.2.1/g' config/config_template.yaml | tee config/config.yaml 1> /dev/null
    else
        # shellcheck disable=SC2016
        sed 's/${IMAGE}/openapitools\/openapi-generator-cli:v6.0.0/g' config/config_template.yaml | tee config/config.yaml 1> /dev/null
    fi
    if [[ $version -lt 19 ]]; then
        cp template-patches/go/identity_providers_legacy.patch template-patches/go/identity_providers.patch
    else
        cp template-patches/go/identity_providers_19_and_above.patch template-patches/go/identity_providers.patch
    fi

    apigentools --api-versions "v${version}" validate
    apigentools --api-versions "v${version}" generate
    mkdir -p "../api/v${version}/openapi/"
    find "generated/sdp-api-client-go/openapi_v${version}" -name '*.go' -exec cp {} "../api/v$version/openapi/" \;
    popd
done

# use custom go run to generate arbitrary go code for the current version, such as identity provicers custom API services and structs
for version in "${supportedVersions[@]}"; do
    go run go-generators/client/main.go  -v -version "$version"
    go test -v -count 1 "./api/v${version}/..."
done


