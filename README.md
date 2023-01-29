# RUYA-API

ruya-api is backend service built in golang

## Technical Selection
* gin
* grom
* mysql
* gorm-gen
* yaml
* JWT
* swagger: gin-swagger
* copier: jinzhu/copier
* Oso: https://github.com/osohq/oso
* Redis: use Redis to cache data

## Features
* Trace all request with trace id
* 

## How to Run
 ### default login user
 username: admin
 password: 111000

 ### Swagger-UI
Access swagger portal: http://localhost:7081/swagger/index.html
* Update swagger
  When you add/update the swagger configuration, need to run below command to regenerate swagger document 
  ```
  swag init
  ```
  how to update the swagger configuration, can refer to https://pkg.go.dev/github.com/swaggo/gin-swagger

### Docker build
```
docker build --pull --rm -f "Dockerfile" -t gomicrodemo:latest "."
```
