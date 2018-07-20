package params

import (
	"apigin/utils"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

type Person struct {
	Username string `form:"username" json:"username" binding:"required"` //binding=required,无参数返回错误信息
	Email    string `form:"email"`
	Web      string `form:"web"`
	Phone    string `form:"phone"`
	Age      string `form:"age"`
	Time     string `form:"time"`
}

func (p *Person) CheckParams(r *http.Request) string {
	rules := govalidator.MapData{
		"username": []string{"required", "between:3,5"},
		"phone":    []string{"digits:11"},
		"email":    []string{"min:4", "max:20", "email"},
		"web":      []string{"url"},
		"age":      []string{"required", "numeric_between:18,65"},
		"time":     []string{"date"},
		"fulltime": []string{"date:dd-mm-yyyy"},
		"regex":    []string{"regex:^[a-zA-Z]+$"},
		"length":   []string{"len:4"},
		"ip":       []string{"ip"},
	}

	messages := govalidator.MapData{
		"username": []string{"required:用户名必填", "between:用户名范围3到5位之间"},
		"phone":    []string{"digits:电话需11位数据"},
		"email":    []string{"min:邮箱最小4位", "max:邮箱最大20位", "email:邮箱格式错误"},
		"web":      []string{"url:URL格式错误"},
		"age":      []string{"required:年龄必填", "numeric_between:年龄范围在18到65岁"},
		"time":     []string{"date:时间格式错误"},
	}

	opts := govalidator.Options{
		Request:         r,
		Rules:           rules,
		Messages:        messages,
		Data:            p,    //请求参数json,解析到结构体
		RequiredDefault: true, //true 默认验证参数都必填
	}

	return utils.CheckFirstErrMsg("form", opts)
}
