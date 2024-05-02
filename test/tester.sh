INTERMEDIATE=$(cat ./totals.json)


RESULT=$(jq '
        {"blob": .}
        ' totals.json)

echo $RESULT
echo $RESULT >tmp.json
