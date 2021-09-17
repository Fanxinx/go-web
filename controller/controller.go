package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	validator2 "webapp/utils/validator"
)

type Con struct {
	c gin.Context
	req map[string]interface{}
	validate interface{}
	validateCheck bool
	logic func() LogicRet
}

type LogicRet struct {
	result interface{}
	msg string
	err error
}

func NewCon(c *gin.Context) *Con {
	return &Con{c:*c,validateCheck: false}
}

//func NewLazyCon(i ...interface{}) *Con{
//	var con *Con
//	con.validateCheck = false
//	for _, value := range i {
//		if v,ok:= value.(gin.Context);ok{
//			con.c = v
//		}
//		if v,ok := value.(interface{});ok{
//			con.SetValidate(v)
//		}
//		if v,ok := value.(func() LogicRet);ok{
//			con.SetLogic(v)
//		}
//	}
//	con.CheckValidate().BindRequest()
//	return con
//}

func (con *Con) BindRequest(){
	if con.validateCheck == false {
		return
	}
	result := con.logic()
	if result.err == nil {
		con.c.JSON(http.StatusOK,gin.H{
			"code":http.StatusOK,
			"msg":result.msg,
			"result":result.result,
		})
		return
	}else {
		con.c.JSON(http.StatusBadRequest,gin.H{
			"code":http.StatusBadRequest,
			"msg":result.msg,
			"err":result.err.Error(),
		})
		return
	}
}

func (con *Con) SetValidate(s interface{}) *Con{
	con.validate = s
	return con
}

func (con *Con) CheckValidate() (*Con){
	reqParams := con.validate
	if err := con.c.ShouldBindJSON(reqParams);err != nil {
		errs, ok := err.(validator.ValidationErrors)
		zap.L().Error("validate error with invalid param",zap.Error(err))
		if !ok {
			con.c.JSON(http.StatusBadRequest,gin.H{
				"code":http.StatusBadRequest,
				"msg":"JSON请求格式错误",
				"err":err.Error(),
			})
			return con
		}else {
			con.c.JSON(http.StatusBadRequest,gin.H{
				"code":http.StatusBadRequest,
				"msg":validator2.RemoveTopStruct(errs.Translate(validator2.GetTrans())),
			})
			return con
		}
	}
	con.req = validator2.ConvertToMap(reqParams)
	con.validateCheck = true
	return con
}

func (con *Con) SetLogic(logic func()(LogicRet))*Con{
	con.logic = logic
	return con
}

func (con *Con) GetAllRequestParam() map[string]interface{}{
	return con.req
}
