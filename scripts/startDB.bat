

docker run -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin --name mongo -v mongo:/data/db  -d mongo:4.2.1