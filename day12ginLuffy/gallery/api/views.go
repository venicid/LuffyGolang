package api

import (
	usersDB "day12ginLuffy/gallery/models"
	"day12ginLuffy/logging"
	"day12ginLuffy/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userinfo struct {
	Name string `json:"name"`
	Sex string `json:"sex"`
	Phone string `json:"phone"`
	City string `json:"city"`

}

type Handler struct {
	userDB usersDB.UserDB
}


/*
增删改查
*/
func (h *Handler) createUser(c *gin.Context)  {
	var i userinfo
	if err := c.BindJSON(&i); err != nil{
		return
	}

	phone, _ := strconv.Atoi(i.Phone)
	u := usersDB.UserInfoModel{
		Name: i.Name,
		Sex: i.Sex,
		Phone: phone,
		City: i.City,
	}

	h.userDB.Db.Create(&u)
	c.JSON(http.StatusCreated, i)
}

func (h *Handler) fetchAllUsers(c *gin.Context){
	var users []usersDB.UserInfoModel
	h.userDB.Db.Find(&users)
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


func (h *Handler) fetchSingleUser(c *gin.Context)  {
	var user usersDB.UserInfoModel
	userId := c.Param("id")
	h.userDB.Db.First(&user, userId)
	if user.ID == 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user Found"})
		return
	}
	_user := usersDB.TransformedUserInfo{
		ID:   user.ID,
		Name: user.Name,
		Phone: user.Phone,

	}
	//c.JSON(http.StatusOK, &user)
	c.JSON(http.StatusOK, _user)
}
func (h *Handler) updateSingleUser(c *gin.Context){
	var (
		user usersDB.UserInfoModel
		putInfo userinfo
	)

	userID := c.Param("id")
	h.userDB.Db.First(&user, userID)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found"})
	}
	c.BindJSON(&putInfo)
	h.userDB.Db.Model(&user).Update("city", putInfo.City)
	phone, _ := strconv.Atoi(putInfo.Phone)
	h.userDB.Db.Model(&user).Update("phone", phone)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message":"User updated successfully",
	})

}
func (h *Handler) deleteSingleUser(c *gin.Context)  {
	var user usersDB.UserInfoModel
	userId := c.Param("id")
	h.userDB.Db.First(&user, userId)
	if user.ID == 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found"})
	}
	h.userDB.Db.Delete(&usersDB.UserInfoModel{}, userId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "delete user successfully"})

}

func (h *Handler) fetchAllUsersByPaging(c *gin.Context)  {
	var users []usersDB.UserInfoModel
	var pagination Pagination

	if err := c.ShouldBindQuery(&pagination); err != nil{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found"})
		return
	}
	h.userDB.Db.Scopes(paginate(users, &pagination, h.userDB.Db)).Find(&users)
	pagination.Results = users
	c.JSON(http.StatusOK, &pagination)
}



/*
router路由
*/
func RouteV1(cfg *settings.Config, h *Handler, r *gin.Engine)  {
	v1 := r.Group("/api/v1/users")
	{
		v1.POST("/", h.createUser)
		v1.GET("/", h.fetchAllUsers)
		v1.GET("/:id", h.fetchSingleUser)
		v1.PUT("/:id", h.updateSingleUser)
		v1.DELETE("/:id", h.deleteSingleUser)

		v1.GET("/paging", h.fetchAllUsersByPaging)
	}
}

func RouteV2(cfg *settings.Config, h *Handler, r *gin.Engine)  {
	v2 := r.Group("/api/v2/users", gin.BasicAuth(gin.Accounts{
		"abc":"123",
	}))
	{
		v2.POST("/", h.createUser)
	}
}

// 为视图创造构造函数，方便其他地方调用
func NewHandler(usersDB *usersDB.UserDB) *Handler  {
	return &Handler{userDB: *usersDB}
}