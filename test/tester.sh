BUILD_IDS=""

for VERSION in 8 11 17 21
do 
    BUILDS_FILE=data/builds-${VERSION}.json
    if [ -f ${BUILDS_FILE} ]; then
        BUILD_IDS=${BUILD_IDS}' '$(jq '.ids[]' ${BUILDS_FILE})
    fi
done

RESULT=$(echo ${BUILD_IDS} | jq -n '{ids: [inputs]}')
echo $RESULT > data/builds.json