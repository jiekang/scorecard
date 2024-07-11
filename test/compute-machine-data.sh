#!/bin/bash

DATA_FILE=data/machine-data.json

RESULT_FILE=data/machine-results.json

RESULT=$(jq '[.computer | .[] |
    {
    "displayName": .displayName,
    "testLabel": .assignedLabels | .[] | select(.name=="ci.role.test") | .name,
    "arch": .assignedLabels | .[] | select(.name | startswith("hw.arch")) | .name,
    "os": .assignedLabels | .[] | select(.name | startswith("sw.os")) | .name,
    "offline": .offline
    }]
    ' ${DATA_FILE})

echo $RESULT > ${RESULT_FILE}

# linux x86, aarch64, ppc64, s390x, aarch32
# mac x86, aarch64
# windows x86, aarch64
