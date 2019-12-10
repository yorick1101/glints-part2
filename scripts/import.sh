
network=glints-part2_default
db_user=admin
db_password=admin
db_host=mongo
db_port=27017
db_name=glints

docker run --network=$network -e ENV_DB_USER=$db_user -e ENV_DB_PASSWORD=$db_password -e ENV_DB_HOST=$db_host -e ENV_DB_PORT=$db_port -e ENV_DB_NAME=$db_name -v $1:/app/companies.json glints-part2-etl:latest /app/etl.exe /app/companies.json
