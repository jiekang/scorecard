#!/bin/bash

DATA_FILE=data/machine-data.json

if [ ! -f ${DATA_FILE} ]; then
    DATA_URL=https://ci.adoptium.net/computer/api/json
    echo "Aquiring machine data via ${DATA_URL}"
    curl -s -X 'GET' \
        ${DATA_URL} \
        -H 'accept: application/json' >${DATA_FILE}
fi

