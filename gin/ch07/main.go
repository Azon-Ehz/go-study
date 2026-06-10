package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	transEn "github.com/go-playground/validator/v10/translations/en"
	transZh "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

type LoginForm struct {
	Account  string `form:"account" json:"user" xml:"user" binding:"required,min=3,max=10"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type SignUpForm struct {
	Age        uint8  `form:"age" json:"user" xml:"user" binding:"required,gte=1,lte=130"`
	Name       string `form:"name" json:"name" xml:"name" binding:"required,min=3"`
	Email      string `form:"email" json:"email" xml:"email" binding:"required,email"`
	Password   string `form:"password" json:"password" xml:"password" binding:"required"`
	RePassword string `form:"repassword" json:"repassword" xml:"repassword" binding:"required,eqfield=Password"`
}

func InitTrans(locale string) (err error) {
	//修改gin框架中的validator引擎属性 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册获取 jsonTag 的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个是备用语言环境，后面的是应该支持的语言环境
		uni := ut.New(zhT, zhT, enT)
		if trans, ok = uni.GetTranslator(locale); ok {
			switch locale {
			case "en":
				transEn.RegisterDefaultTranslations(v, trans)
				break
			case "zh":
				transZh.RegisterDefaultTranslations(v, trans)
				break
			default:
				transEn.RegisterDefaultTranslations(v, trans)
			}
			return nil
		} else {
			return fmt.Errorf("translator not found %s", locale)
		}
	}
	return nil
}

func main() {
	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}
	router := gin.Default()
	router.POST("login", LoginJson)
	router.POST("register", RegisterJson)
	router.Run(":8383")
}

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func LoginJson(c *gin.Context) {
	var loginForm LoginForm
	if err := c.ShouldBind(&loginForm); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": errs.Error(),
			})
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account":  loginForm.Account,
		"password": loginForm.Password,
		"message":  "login success",
	})
}

func RegisterJson(c *gin.Context) {
	var registerForm SignUpForm
	if err := c.ShouldBind(&registerForm); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": errs.Error(),
			})
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account":  registerForm.Name,
		"password": registerForm.Password,
		"message":  "register success",
	})
}
