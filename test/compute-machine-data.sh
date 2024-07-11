#!/bin/bash

DATA_FILE=data/machine-data.json

RESULT_FILE=data/machine-results.json

RESULT=$(jq '[.computer | .[] | select(.offline == false) |
    {
    "displayName": .displayName,
    "testLabel": .assignedLabels | .[] | select(.name=="ci.role.test") | .name,
    "arch": .assignedLabels | .[] | select(.name | startswith("hw.arch")) | .name,
    "os": .assignedLabels | .[] | select(.name | startswith("sw.os")) | .name,
    }]
    ' ${DATA_FILE})

echo $RESULT > ${RESULT_FILE}