VERSION=$1
if [ -z $VERSION ]; then
    echo "Usage $0 version-number"
    exit
fi

mkdir -p data

PIPELINE_NAME=release-openjdk${VERSION}-pipeline
PIPELINE_INFO=data/${PIPELINE_NAME}.json

echo "Processing ${PIPELINE_NAME}"

URL="https://trss.adoptium.net/api/getBuildHistory?buildName=${PIPELINE_NAME}&url=https://ci.adoptium.net/job/build-scripts&limit=120"

curl -s -X 'GET' \
    ${URL} \
    -H 'accept: application/json' >${PIPELINE_INFO}

jq '.[] | select((.status == "Done") and (.startBy | . and contains ("upstream")))' ${PIPELINE_INFO}
