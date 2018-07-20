package ctrls

import (
	"apigin/apis/models"
	"apigin/apis/params"
	"apigin/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GET,测试数据库
func TestDb(c *gin.Context) {
	mid, _ := strconv.Atoi(c.Param("id"))
	mem, err := models.OneMember(mid)
	if err != nil {
		utils.EchoMessage(c, http.StatusExpectationFailed, http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   mem,
		})
	}
}

//GET,测试参数验证过滤
func TestValidator(c *gin.Context) {
	person := new(params.Person)
	// log.Printf("query:%s,urlquery:%s,formvalue:%s,default:%s", c.Query("username"), c.Request.URL.Query().Get("username"), c.Request.FormValue("username"), c.DefaultQuery("username", "anonymous"))
	if errmsg := person.CheckParams(c.Request); errmsg != "" {
		utils.EchoMessage(c, http.StatusExpectationFailed, http.StatusBadRequest, errmsg)
	}

	if c.ShouldBindQuery(person) == nil {
		//解析get数据到结构体
		log.Println(person.Username)
		log.Println(person.Email)
	}

	c.String(200, "Success")
}

//POST,测试表单文件上传
func TestFile(c *gin.Context) {
	m := new(models.Member)
	// m.LoginName = c.PostForm("login_name")
	// m.Password = c.DefaultPostForm("password", "password")
	m.LoginName = c.Request.FormValue("login_name")
	m.Password = c.Request.FormValue("password")

}
