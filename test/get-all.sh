AFTER_DATE=$1
if [ -z ${AFTER_DATE} ]; then
    echo "Usage $0 after-date before-date"
    exit
fi

BEFORE_DATE=$2
if [ -z ${BEFORE_DATE} ]; then
    echo "Usage $0 after-date before-date"
    exit
fi

for VERSION in 8 11 17 21; do
    ./get-test-data.sh ${VERSION} ${AFTER_DATE} ${BEFORE_DATE}
done
