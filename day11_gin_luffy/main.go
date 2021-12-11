package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main/logging"
	"net/http"
	"strconv"
)



// 用户 -> restful http接口 ——> 序列化 --> 入库
// 数据库 --> 反序列化 -> restful http接口  ——-> 用户

type (
	// 定义用户原始的数据库字段
	UserInfoModel struct {
		gorm.Model
		Name string `json:"name"`
		Sex string `json:"sex"`
		Phone int `json:"phone"`
		City string `json:"city"`
	}

	// 处理返回的字段
	transformedUserInfo struct {
		ID uint `json:"id"`
		Name string `json:"name"`
		Phone int `json:"phone"`
}
)

type userinfo struct {
	Name string `json:"name"`
	Sex string `json:"sex"`
	Phone string `json:"phone"`
	City string `json:"city"`

}

var db *gorm.DB
func init()  {

	var err error
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root112358@tcp(150.158.171.205:3306)/luffysex?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("failed to connect databases")
	}
	
	// Migrate the schema
	err1 := db.AutoMigrate(&UserInfoModel{})
	if err1 != nil {
		return 
	}

}



func createUser(c *gin.Context)  {
	var i userinfo
	if err := c.BindJSON(&i); err != nil{
		return
	}

	phone, _ := strconv.Atoi(i.Phone)
	u := UserInfoModel{
		Name: i.Name,
		Sex: i.Sex,
		Phone: phone,
		City: i.City,
	}

	log.Println(db)
	db.Create(&u)

	c.JSON(http.StatusCreated, i)

}

func fetchAllUsers(c *gin.Context){
	var users []UserInfoModel
	db.Find(&users)
	if len(users) <=0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found"})
	}
	c.JSON(http.StatusOK, &users)

	/*
	打印用户鉴权信息
	*/
	authKey, isAuth := c.Get(gin.AuthUserKey)
	fmt.Println(authKey)
	fmt.Println(isAuth)
	if !isAuth{
		return
	}
	user := c.MustGet(gin.AuthUserKey).(string)
	fmt.Println(user)
	/*
		使用zap日志打印
		2021-12-11T11:18:58.808+0800	debug	day11_gin_luffy/main.go:106	abc
	*/
	logging.DefaultLogger().Debug(user)

}


func fetchSingleUser(c *gin.Context)  {
	var user UserInfoModel
	userId := c.Param("id")
	db.First(&user, userId)
	if user.ID == 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user Found"})
		return
	}
	_user := transformedUserInfo{
		ID:   user.ID,
		Name: user.Name,
		Phone: user.Phone,

	}
	//c.JSON(http.StatusOK, &user)
	c.JSON(http.StatusOK, _user)
}
func updateSingleUser(c *gin.Context){
	var (
		user UserInfoModel
		putInfo userinfo
	)

	userID := c.Param("id")
	db.First(&user, userID)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found"})
	}
	c.BindJSON(&putInfo)
	db.Model(&user).Update("city", putInfo.City)
	phone, _ := strconv.Atoi(putInfo.Phone)
	db.Model(&user).Update("phone", phone)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message":"User updated successfully",
	})

}
func deleteSingleUser(c *gin.Context)  {
	var user UserInfoModel
	userId := c.Param("id")
	db.First(&user, userId)
	if user.ID == 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found"})
	}
	db.Delete(&UserInfoModel{}, userId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "delete user successfully"})

}

/*
分页
*/
type Pagination struct {
	Limit int `form:"limit, omitempty; query:limit"`
	Page int `form:"page,omitempty;query:page"`
	Total int64 `form:"total"`
	Results interface{} `form:"results"`
}

func (p Pagination) GetPage() int {
	if p.Page == 0{
		p.Page = 1
	}
	return p.Page
}
func (p Pagination) GetLimit() int {
	if p.Limit == 0{
		p.Limit = 10
	}
	return p.Limit
}

func (p Pagination) GetOffset() int {
	return (p.GetPage() -1 ) * p.GetLimit()
}

func paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB{
	var totalRows int64
	db.Model(value).Count(&totalRows)
	pagination.Total = totalRows
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
func fetchAllUsersByPaging(c *gin.Context)  {
	var users []UserInfoModel
	var pagination Pagination

	if err := c.ShouldBindQuery(&pagination); err != nil{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found"})
		return
	}
	db.Scopes(paginate(users, &pagination, db)).Find(&users)
	pagination.Results = users
	c.JSON(http.StatusOK, &pagination)
}

func main()  {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	/*
	添加跨域
	*/
	router.Use(cors.Default())
	/*
	使用用户鉴权

	请求方式：curl --location --request GET '127.0.0.1:8080/api/v2/users/' \
	--header 'Authorization: Basic YWJjOjEyMw=='
	*/
	v2 := router.Group("/api/v2/users", gin.BasicAuth(gin.Accounts{
		"abc":"123",
	}))
	{
		v2.POST("/", createUser)
		v2.GET("/", fetchAllUsers)
		v2.GET("/:id", fetchSingleUser)
		v2.PUT("/:id", updateSingleUser)
		v2.DELETE("/:id", deleteSingleUser)
		v2.GET("/paging", fetchAllUsersByPaging)
	}


	v1 := router.Group("/api/v1/users")
	{
		v1.POST("/", createUser)
		v1.GET("/", fetchAllUsers)
		v1.GET("/:id", fetchSingleUser)
		v1.PUT("/:id", updateSingleUser)
		v1.DELETE("/:id", deleteSingleUser)

		v1.GET("/paging", fetchAllUsersByPaging)
	}
	err := router.Run()
	if err != nil {
		return 
	}
}


/*
# 创建数据库,默认字符集是latin
create database if not exists luffysex
default charset utf8 collate utf8_general_ci;

# 启动gin失败
[GIN-debug] [ERROR] listen tcp :8080: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.

# 排查win10
- 查看端口号
netstat -aon|findstr "8080"
- 终止进程
taskkill -F -PID 21064


# 创建用户
curl --location --request POST '127.0.0.1:8080/api/v1/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"alex",
    "sex": "男",
    "phone":"119",
    "city":"beijing"
}'
*/