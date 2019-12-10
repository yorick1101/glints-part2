base=$(dirname "$0")

docker build -t glints-part2-etl:latest -f $base/../api/Dockerfile-etl $base/../api