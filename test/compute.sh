BUILD_IDS=$(jq '.ids[]' data/builds.json)

LEVELS="dev sanity extended special"
GROUPS="build functional openjdk system external perf jck"
LINUX_SYSTEMS="x86-64_linux aarch64_linux s390x_linux ppc64le_linux"

function computeBuild() {
    ID=$1
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
        "tests": .tests
        }]' ${FILE})
    echo ${RESULT} >data/child/test-${ID}-compute.json
}

TOTAL_EXECUTION_TIME=0

PLATFORM_EXECUTION_TIME=0

VERSION_EXECUTION_TIME=0

for ID in $(echo ${BUILD_IDS} | tr -d \"); do
    computeTestData $ID
done
