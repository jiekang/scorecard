#!/bin/bash

function pruneBuilds() {
    BUILD_IDS=$(jq '.ids[]' data/builds.json)
    for ID in $(echo ${BUILD_IDS} | tr -d \"); do
        pruneBuildData $ID
        pruneTestData $ID
    done
}

function pruneBuildData() {
    ID=$1
    FILE=data/child/jdk-${ID}.json
    RESULT=$(jq '[.[] | select(.type == "Test") | 
        {"buildName": .buildName,
        "buildDuration": .buildDuration,
        "testSummary": .testSummary,
        "platform": .buildParams | .[] | select(.name=="PLATFORM") | .value,
        "version": .buildParams | .[] | select(.name=="JDK_VERSION") | .value,
        "topLevelTarget": .buildParams | .[] | select(.name=="TARGET") | .value,
        "testTargets": .tests
        }]' ${FILE})
    echo ${RESULT} >data/child/jdk-${ID}-compute.json
}

function pruneTestData() {
    ID=$1
    FILE=data/child/test-${ID}.json
    RESULT=$(jq '[.[] | 
        {"buildName": .buildName,
        "buildDuration": .buildDuration,
        "testSummary": .testSummary,
        "platform": .buildParams | .[] | select(.name=="PLATFORM") | .value,
        "version": .buildParams | .[] | select(.name=="JDK_VERSION") | .value,
        "topLevelTarget": .buildParams | .[] | select(.name=="TARGET") | .value,
        "testTargets": .tests
        }]' ${FILE})
    echo ${RESULT} >data/child/test-${ID}-compute.json
}

BUILD_IDS=""

for VERSION in 8 11 17 21; do
    BUILDS_FILE=data/builds-${VERSION}.json
    if [ -f ${BUILDS_FILE} ]; then
        BUILD_IDS=${BUILD_IDS}' '$(jq '.ids[]' ${BUILDS_FILE})
    fi
done
RESULT=$(echo ${BUILD_IDS} | jq -n '{ids: [inputs]}')
echo $RESULT >data/builds.json

pruneBuilds
./compute $1 $2 $3
