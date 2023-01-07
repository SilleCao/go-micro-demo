# GO-MICRO-DEMO

## Technical Selection
* gin
* grom
* mysql
* gorm-gen
* yaml
* JWT
* swagger: gin-swagger


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
