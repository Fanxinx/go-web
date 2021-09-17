package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"webapp/logic"
	validator2 "webapp/utils/validator"
)

func SignUpHandler(c *gin.Context){
	type validate struct {
		UserName   string `json:"username" binding:"required"`
		Password   string `json:"password" binding:"required"`
		RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	}
	reqParams := new(validate)
	if err := c.ShouldBindJSON(reqParams);err != nil{
		errs,ok := err.(validator.ValidationErrors)
		if !ok{

			return
		}


		c.JSON(http.StatusBadRequest,gin.H{
			"msg":validator2.RemoveTopStruct(errs.Translate(validator2.GetTrans())),
		})

		return
	}

	var s = struct{user string}{
		user : "dsadasda",
	}
	fmt.Println("type",reflect.TypeOf(s))


	m := validator2.ConvertToMap(reqParams)
	err := logic.SignUp(m)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"name":reqParams.UserName,
		"paswd":reqParams.Password,
		"msg":"注册成功",
	})
	//1.参数校验  validate
	//2.执行对应的业务，或异常返回 logic  or  (catch err)
	//3.业务数据返回	response c.Json(http.statusOK,"ok")
}

func LoginHandler(c *gin.Context)  {
	//定义validate
	type validate struct {
		UserName   string `json:"username" binding:"required"`
		Password   string `json:"password" binding:"required"`
	}
	//验证validate
	controller := NewCon(c).SetValidate(new(validate)).CheckValidate()
	//定义logic流程
	var logic = func()(LogicRet){
		//通过返回的error来控制response的请求头是200 还是400
		//result msg err
		var res LogicRet
		res.result,res.err = logic.Login(controller.GetAllRequestParam())
		if res.err == nil{
			res.msg = "登录成功"
		}else {
			res.msg = "登录失败"
		}
		return res
	}
	//执行请求
	controller.SetLogic(logic).BindRequest()
}

func Index(c *gin.Context){
	//定义验证器类
	type validate struct {
		Id 	int64  `json:"id" binding:"required"`
	}
	//定义逻辑执行流程
	var logic = func() LogicRet{
		//定义返回结果
		var logicRet LogicRet
		//封装返回结果
		logicRet.result = gin.H{
			"userId":   c.MustGet("userId"),
			"userName": c.MustGet("userName"),
		}
		return logicRet
	}
	////请求封装执行
	{
		NewCon(c).
			SetLogic(logic).
			SetValidate(new(validate)).CheckValidate().
			BindRequest()
	}
}