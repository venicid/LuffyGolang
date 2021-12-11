## Learning Gin By Luffycity
----- Boyle GU

### Go Gin Web框架课程内容
1. 用Gin + gorm V2 实现一套标准的restful接口，(包括分页、用户鉴权等)
2. 高性能动态注入组织大型Gin项目最佳时间
3. 基于Gin的仿celery异步接口调用实现
4. fasthttp介绍

### Restful接口
开发场景 /api/v1/users/
1. 创建用户 POST /api/v1/users/
2. 用户列表 GET /api/v1/users/
3. 用户详情 GET /api/v1/users/:id
4. 更新用户 PUT /api/v1/users/:id
5. 删除用户 DELETE /api/v1/users/:id

数据库：Mysql5.7


### 跨域&用户鉴权&日志
跨域 cors

用户鉴权 gin.BasicAuth

日志 zap

