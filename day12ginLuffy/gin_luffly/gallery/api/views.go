package api

import (
	"gin_luffly/async"
	usersDB "gin_luffly/gallery/models"
	"gin_luffly/settings"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	usersDB usersDB.UserDB
}

func (h *Handler) createUser(c *gin.Context) {
	var i userinfo
	if err := c.BindJSON(&i); err != nil {
		return
	}
	phone, _ := strconv.Atoi(i.Phone)
	u := usersDB.UserInfoModel{
		Name:  i.Name,
		Sex:   i.Sex,
		Phone: phone,
	}
	h.usersDB.Db.Create(&u)
	c.JSON(http.StatusCreated, i)
}

func (h *Handler) fetchAllUsers(c *gin.Context) {
	var users []usersDB.UserInfoModel
	h.usersDB.Db.Find(&users)
	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
	}
	c.JSON(http.StatusOK, &users)
}

func (h *Handler) fetchSingleUsers(c *gin.Context) {
	var users usersDB.UserInfoModel
	userID := c.Param("id")
	h.usersDB.Db.First(&users, userID)
	if users.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}
	_user := usersDB.TransformedUserInfo{ID: users.ID, Name: users.Name}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _user})
}

func (h *Handler) updateUser(c *gin.Context) {
	var (
		users   usersDB.UserInfoModel
		putInfo userinfo
	)
	userID := c.Param("id")
	h.usersDB.Db.First(&users, userID)
	if users.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}
	c.BindJSON(&putInfo)
	h.usersDB.Db.Model(&users).Update("city", putInfo.City)
	phone, _ := strconv.Atoi(putInfo.Phone)
	h.usersDB.Db.Model(&users).Update("phone", phone)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User updated successfully"})
}

type userinfo struct {
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Phone string `json:"phone"`
	City  string `json:"city"`
}

func (h *Handler) fetchAllUserByPaging(c *gin.Context) {
	var users []usersDB.UserInfoModel
	var pagination Pagination
	//user := c.MustGet(gin.AuthUserKey).(string)
	//logging.DefaultLogger().Debug(user)

	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found"})
		return
	}
	h.usersDB.Db.Scopes(paginate(users, &pagination, h.usersDB.Db)).Find(&users)
	pagination.Results = users
	c.JSON(http.StatusOK, &pagination)

}

func (h *Handler) longTask(c *gin.Context) {
	// time.Sleep(time.Millisecond * 1000000)
	async.Send()
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK, "message": "ok"})
}

func RouteV1(cfg *settings.Config, h *Handler, r *gin.Engine) {
	v1 := r.Group("/api/v1/users")
	{
		v1.POST("/", h.createUser)
		v1.GET("/", h.fetchAllUsers)
		v1.GET("/:id", h.fetchSingleUsers)
		v1.PUT("/:id", h.updateUser)
		v1.GET("paging/", h.fetchAllUserByPaging)

		v1.GET("/longtask/", h.longTask)
	}
}

func NewHandler(usersDB *usersDB.UserDB) *Handler {
	return &Handler{usersDB: *usersDB}
}
