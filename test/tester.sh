ID=661f70c9879917006e985908

FILE=data/child/jdk-${ID}.json
# RESULT=$(jq '[.[] | select(.type == "Test") | 
#         {"buildName": .buildName,
#         "buildDuration": .buildDuration,
#         "testSummary": .testSummary,
#         "platform": .buildParams | .[] | select(.name=="PLATFORM") | .value,
#         "version": .buildParams | .[] | select(.name=="JDK_VERSION") | .value,
#         "tests": .tests
#         }]' ${FILE})

RESULT=$(jq '.[] | select(.type == "Test") | {"buildName": .buildName}' ${FILE})
echo ${RESULT} > tmp.json