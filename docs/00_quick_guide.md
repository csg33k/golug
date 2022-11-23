## Initial steps

Install docker, docker-compose and if you're using defaults, just copy env.template to .env

## Bootstrap the database

### Bring up the database.  

```sh 
docker-compose up -d db 
```

### Apply SQL to database

```sh 
dbmate up 
```

sample output:
Applying: 20221110141229_linux_user.sql
Writing: ./db/schema.sql

## Build app
### Lazy pattern

```sh 
make build
```
### Detailed Steps

#### Generate swagger

```sh 
go generate
```

sample output:
2022/11/23 17:36:12 Generate swagger docs....
2022/11/23 17:36:12 Generate general API Info, search dir:./
2022/11/23 17:36:14 Generating models.LinuxUser
2022/11/23 17:36:14 create docs.go at  api/docs.go
2022/11/23 17:36:14 create swagger.json at  api/swagger.json
2022/11/23 17:36:14 create swagger.yaml at  api/swagger.yaml

#### Build binary 
go build -o www_svc

## Run Application 

update conf/golog.yaml as appropriate then simply run: ./www_svc

All endpoints should be documented under http://localhost:3000/swaggerui/

