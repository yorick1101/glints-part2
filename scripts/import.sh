

docker run --network=glints-part2_default -e ENV_DB_USER=admin -e ENV_DB_PASSWORD=admin -e ENV_DB_HOST=mongo -e ENV_DB_PORT=27017 -e ENV_DB_NAME=glints -v $1:/app/companies.json -it etl:latest /app/etl.exe /app/companies.json
