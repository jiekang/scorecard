VERSION=$1
if [ -z $VERSION ]; then
    echo "Usage $0 version-number"
    exit
fi

mkdir -p data

PIPELINE_NAME=release-openjdk${VERSION}-pipeline
PIPELINE_INFO=data/${PIPELINE_NAME}.json

echo "Processing releases for JDK ${VERSION}"

RELEASE_INFO=data/releases${VERSION}.json

RELEASES_URL="https://api.adoptium.net/v3/assets/feature_releases/${VERSION}/ga?heap_size=normal&image_type=jdk&jvm_impl=hotspot&page=0&page_size=10&project=jdk&sort_method=DEFAULT&sort_order=DESC&vendor=eclipse"

if [ ! -f ./${RELEASES_INFP} ]; then
    curl -s -X 'GET' \
        ${RELEASES_URL} \
        -H 'accept: application/json' >${RELEASE_INFO}
fi
echo "Processing pipelines for JDK ${VERSION}"

PIPELINE_URL="https://trss.adoptium.net/api/getBuildHistory?buildName=${PIPELINE_NAME}&url=https://ci.adoptium.net/job/build-scripts&limit=120"

if [ ! -f ./${PIPELINE_INFO} ]; then
    curl -s -X 'GET' \
        ${PIPELINE_URL} \
        -H 'accept: application/json' >${PIPELINE_INFO}
fi

DATA=$(jq '.[] | select(.status == "Done") | ._id' ${PIPELINE_INFO})
mkdir -p data/totals/

for ID in $(echo $DATA | tr -d \"); do
    TOTALS_URL=https://trss.adoptium.net/api/getTotals?id=${ID}
    curl -s -X 'GET' \
        ${TOTALS_URL} \
        -H 'accept: application/json' >data/totals/${ID}.json
done

LEVELS="dev sanity extended special"
GROUPS="build functional openjdk system external perf jck"
