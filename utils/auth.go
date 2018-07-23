package utils

import (
	"apigin/config"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

//返回接口参数第一个错误信息
func CheckFirstErrMsg(typ string, opts govalidator.Options) string {

	v := govalidator.New(opts)
	//log.Println(v.Validate())

	if typ == "json" {
		e := v.ValidateJSON()
		if e != nil {
			for _, arr := range e {
				return arr[0]
			}
		}
	} else {
		e := v.Validate()
		if e != nil {
			for _, arr := range e {
				return arr[0]
			}
		}
	}

	return ""
}

//退出登录清除session会话
func Clear(c *gin.Context, id string) bool {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	config.RedisHandle.Del(fmt.Sprintf("user:%s", id))
	return true
}

//登录成功设置session会话
func PreSession(c *gin.Context, id string) bool {
	userAccessToken := RandStr(32)
	session := sessions.Default(c)
	session.Set("user", id)
	session.Set("security", userAccessToken)
	session.Save()
	config.RedisHandle.Set(fmt.Sprintf("user:%s", id), userAccessToken, 86400000000)
	return true
}

//验证登录信息
func Checklogic(c *gin.Context) bool {
	session := sessions.Default(c)
	session.Save()

	if u := session.Get("user"); u != nil {

		if s := session.Get("security"); s != nil {
			//用户ID找到redis对应key匹配
			if s.(string) != config.RedisHandle.Get(fmt.Sprintf("user:%s", u.(string))).Val() {
				return false
			}
		} else {
			return false
		}

		//设置全局变量
		c.Set("user", session.Get("user"))
	} else {
		return false
	}

	return true
}

//输出信息
func Result(c *gin.Context, code, status int, msg string) {
	c.JSON(code, gin.H{
		"status":  status,
		"message": msg,
	})
	c.Abort()
}
