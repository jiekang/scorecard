#!/bin/bash

# e.g. 17
VERSION=$1
if [ -z $VERSION ]; then
    echo "Usage $0 version-number after-date before-date"
    exit
fi

# e.g. 2024-04-01
AFTER_DATE=$2
if [ -z $AFTER_DATE ]; then
    echo "Usage $0 version-number after-date before-date"
    exit
fi

BEFORE_DATE=$3
if [ -z $BEFORE_DATE ]; then
    echo "Usage $0 version-number after-date before-date"
    exit
fi

AFTER_TIMESTAMP=$(date -u -d "${AFTER_DATE}" "+%s%3N")
BEFORE_TIMESTAMP=$(date -u -d "${BEFORE_DATE}" "+%s%3N")

mkdir -p data

PIPELINE_NAME=release-openjdk${VERSION}-pipeline
PIPELINE_INFO=data/${PIPELINE_NAME}.json

RELEASE_INFO=data/releases${VERSION}.json

if [ ! -f ${RELEASES_INFO} ]; then
    RELEASES_URL="https://api.adoptium.net/v3/assets/feature_releases/${VERSION}/ga?heap_size=normal&image_type=jdk&jvm_impl=hotspot&page=0&page_size=10&project=jdk&sort_method=DEFAULT&sort_order=DESC&vendor=eclipse"
    echo "Processing releases for JDK ${VERSION} via ${RELEASES_URL}"
    curl -s -X 'GET' \
        ${RELEASES_URL} \
        -H 'accept: application/json' >${RELEASE_INFO}
fi

if [ ! -f ${PIPELINE_INFO} ]; then
    PIPELINE_URL="https://trss.adoptium.net/api/getBuildHistory?buildName=${PIPELINE_NAME}&status=Done&url=https://ci.adoptium.net/job/build-scripts&limit=120"
    echo "Processing pipelines for JDK ${VERSION} via ${PIPELINE_URL}"
    curl -s -X 'GET' \
        ${PIPELINE_URL} \
        -H 'accept: application/json' >${PIPELINE_INFO}
fi

BUILD_IDS=$(jq '{ids: [.[] | select(.timestamp > '${AFTER_TIMESTAMP}' and .timestamp < '${BEFORE_TIMESTAMP}') | ._id]}' ${PIPELINE_INFO})
echo ${BUILD_IDS} >data/builds-${VERSION}.json

BUILD_IDS=$(jq '.ids[]' data/builds-${VERSION}.json)
mkdir -p data/child/

echo "Processing builds for JDK ${VERSION}"

for ID in $(echo ${BUILD_IDS} | tr -d \"); do
    if [ ! -f data/child/jdk-${ID}.json ]; then
        CHILD_JDK_URL="https://trss.adoptium.net/api/getAllChildBuilds?buildNameRegex=%5E(jdk%5B0-9%5D%7B1%2C2%7D%7CBuild_)&parentId=${ID}"
        echo "Acquiring jdk jobs for ${ID} via ${CHILD_JDK_URL}"
        curl -s -X 'GET' \
            ${CHILD_JDK_URL} \
            -H 'accept: application/json' >data/child/jdk-${ID}.json

        CHILD_TEST_URL="https://trss.adoptium.net/api/getAllChildBuilds?buildNameRegex=%5ETest_openjdk.*&parentId=${ID}"
        echo "Acquiring test jobs for ${ID} via ${CHILD_TEST_URL}"
        curl -s -X 'GET' \
            ${CHILD_TEST_URL} \
            -H 'accept: application/json' >data/child/test-${ID}.json
    fi
done
