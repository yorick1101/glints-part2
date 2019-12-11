# glints-part2

### Prerequisites
- Docker, Docker Compose are installed 
- Make sure port 80 is available
- The import data script can only run in a linux environment

### API Server
- API server is composite of three components,
  - MongoDB as database
  - Web server which is implemented in golang as service provider
  - Nginx as reversed proxy and listens on port 80 for outside world
 
- These three components are linked to each other inside a docker-compose network. 
- The data of database is persisted in a docker volume so that the data won't be disppeared once the container is restarted. 
 
#### Build and Start
- After checkout the project, cd into glints-part2
- ```docker-compose build```
- ```docker-compose up -d``` to start in the background or ```docker-compose up ``` if you want to check the log directly on the terminal
- All APIs are running on url http://localhost/api/xxxx
#### Stop
- ```docker-compose stop```

#### Rate-limiting and Response Time Monitoring
- The setting is in nginx/nginx.conf, the current value is 10r/s per remote address
- The access log contains the request_time and upstream_response_time as the last two column and the unit is second
- As the example below, the request time is 0.037 sec and upstream response time is 0.030 sec
```
ginx_1  | 172.19.0.1 - - [10/Dec/2019:11:02:18 +0000] "GET /api/company?id=52cdef7c4bab8bd675297d8a HTTP/1.1" 200 5175 "http://localhost:8080/" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36" 0.037 0.030
```
### Import Data to Database
- The API server must be started so that database is alive
- Goto the glints-part2/scripts
- Make the script executable
```
chmod +x build-etl.sh
chmod +x import.sh
```
#### Build
- Build by executing the script build-etl.sh, an image glints-part2-etl:latest should be built
```
./build-etl.sh
```
#### Run
- The script works by joining the network of API server, please help make sure the <i>network</i> in the script is set to correct value
 ```
 network=glints-part2_default
 ```
- Execute import.sh with path to the data file as first argument, please be noted that full path is required here.
```
./import.sh $(pwd)/companies.json
```
- Log if succeed, one database and two collections should be created
    ```
    time="2019-12-11T02:09:27Z" level=info msg="db connected"
    time="2019-12-11T02:09:27Z" level=info msg="db switch database glints"
    time="2019-12-11T02:09:27Z" level=info msg=/app/companies.json
    time="2019-12-11T02:09:30Z" level=info msg="inserted company:18801"
    ```
- Port 27017 is exposed to host, so that it is reacheable by mongo client on the host

### API Document
- API document is written in openapi 3 format.
- There are 2 way to check the document
  - Use cloud swagger editor
    - Goto https://editor.swagger.io/
    - Select File-> Import file -> select the file glints-part2/doc/swagger.yaml
  - Use the downloaded swagger UI
    - Install http-server via npm
    ```
    npm install http-server -g
    ```
    - Goto folder glints-part2/doc
    - Start the web server
    ```
    http-server -cors
    ```
    - The document can be viewed on http://127.0.0.1:8080/
    - Use "Try it out" button to send request to local api server.


