# e.g. 17
VERSION=$1
if [ -z $VERSION ]; then
    echo "Usage $0 version-number"
    exit
fi

# e.g. 2024-04-01
DATE=$2
if [ -z $DATE] ]; then
    echo "Usage $0 version-number date"
    exit
fi

TIMESTAMP=`date -u -d "${DATE}" "+%s%3N"`

mkdir -p data

PIPELINE_NAME=release-openjdk${VERSION}-pipeline
PIPELINE_INFO=data/${PIPELINE_NAME}.json

echo "Processing releases for JDK ${VERSION}"

RELEASE_INFO=data/releases${VERSION}.json

RELEASES_URL="https://api.adoptium.net/v3/assets/feature_releases/${VERSION}/ga?heap_size=normal&image_type=jdk&jvm_impl=hotspot&page=0&page_size=10&project=jdk&sort_method=DEFAULT&sort_order=DESC&vendor=eclipse"

if [ ! -f ./${RELEASES_INFO} ]; then
    curl -s -X 'GET' \
        ${RELEASES_URL} \
        -H 'accept: application/json' >${RELEASE_INFO}
fi
echo "Processing pipelines for JDK ${VERSION}"

PIPELINE_URL="https://trss.adoptium.net/api/getBuildHistory?buildName=${PIPELINE_NAME}&status=Done&url=https://ci.adoptium.net/job/build-scripts&limit=120"

if [ ! -f ./${PIPELINE_INFO} ]; then
    curl -s -X 'GET' \
        ${PIPELINE_URL} \
        -H 'accept: application/json' >${PIPELINE_INFO}
fi

BUILD_IDS=$(jq '{ids: [.[] | select(.timestamp > '${TIMESTAMP}') | ._id]}' ${PIPELINE_INFO})
echo ${BUILD_IDS} > data/builds.json

BUILD_IDS=$(jq '.ids[]' data/builds.json)
mkdir -p data/totals/
mkdir -p data/child/

echo "Processing builds for JDK ${VERSION}"

for ID in $(echo ${BUILD_IDS} | tr -d \"); do
    echo "Acquiring totals for ${ID}"
    TOTALS_URL="https://trss.adoptium.net/api/getTotals?id=${ID}"
    curl -s -X 'GET' \
        ${TOTALS_URL} \
        -H 'accept: application/json' >data/totals/${ID}.json

    echo "Acquiring jdk jobs for ${ID}"
    CHILD_JDK_URL="https://trss.adoptium.net/api/getAllChildBuilds?buildNameRegex=%5E(jdk%5B0-9%5D%7B1%2C2%7D%7CBuild_)&parentId=${ID}"
    curl -s -X 'GET' \
        ${CHILD_JDK_URL} \
        -H 'accept: application/json' >data/child/jdk-${ID}.json

    echo "Acquiring test jobs for ${ID}"
    CHILD_TEST_URL="https://trss.adoptium.net/api/getAllChildBuilds?buildNameRegex=%5ETest_openjdk.*&parentId=${ID}"
    curl -s -X 'GET' \
        ${CHILD_TEST_URL} \
        -H 'accept: application/json' >data/child/test-${ID}.json
done

LEVELS="dev sanity extended special"
GROUPS="build functional openjdk system external perf jck"
