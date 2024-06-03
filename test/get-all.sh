DATE=$1
if [ -z ${DATE} ]; then
    echo "Usage $0 date"
    exit
fi

for VERSION in 8 11 17 21
do 
    ./get-data.sh ${VERSION} $DATE
done