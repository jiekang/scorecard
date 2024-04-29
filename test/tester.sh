ID=661f70c9879917006e985908

FILE=data/child/test-${ID}.json

RESULT=$(jq '[.[] | 
    {"buildName": .buildName,
    "buildDuration": .buildDuration,
    "testSummary": .testSummary,
    "platform": .buildParams | .[] | select(.name=="PLATFORM") | .value,
    "version": .buildParams | .[] | select(.name=="JDK_VERSION") | .value
    }]' ${FILE})
echo $RESULT
echo $RESULT >tmp.json
