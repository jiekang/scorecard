BUILD_IDS=$(jq '.ids[]' data/builds.json)

function computeBuild() {
    ID=$1
}

function computeBuildData() {
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

function computeTestData() {
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

for ID in $(echo ${BUILD_IDS} | tr -d \"); do
    computeBuildData $ID
    computeTestData $ID
done
