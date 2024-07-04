go build ./cmd/compute
go build ./cmd/diff

./get-all.sh 2024-01-16 2024-01-30
./compute.sh January-24 2024-01

./get-all.sh 2024-04-17 2024-04-24
./compute.sh April-24 2024-04

./diff data/January-24.json data/April-24.json