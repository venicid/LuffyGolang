## Learning Gin By Luffly
---- Boyle Gu

Go Gin Web框架课程内容：
1. 用Gin + gorm v2 实现一套标准的restful接口（包括分页、用户鉴权等）

2. 高性能动态注入组织大型Gin项目最佳实践

3. 基于Gin的仿celery异步接口调用实现

4. fasthttp介绍


### RestFul接口，开发场景
/api/v1/users/

1。 创建用户
   POST   /api/v1/users/  
2. 用户列表
   GET    /api/v1/users
3. 用户详情
   GET    /api/v1/users/:id
4. 更新用户
   PUT/PATCH    /api/v1/users/:id

### 数据库： Mysql 5.7

### 项目组织
gin_luffly
  gallery--
         | -- api
              |
              |-views.go
              |-pagination.go
           -- models.go
  settings 
        |__ db.go
        |__ config.go
        |__ defaults.go
  main.go
    |
    |----  1. 加载项目配置文件
           2. 数据库连接
           3. appname -> views models
           4. 路由配置  （整个项目的入口）

### 启动Gin & 异步队列

1. 启动server：go run main.go server

2. 启动异步队列: go run main.go worker
