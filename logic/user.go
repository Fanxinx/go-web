package logic

import (
	"errors"
	"github.com/gin-gonic/gin"
	"webapp/dao/mysql"
	"webapp/models"
	"webapp/utils/helper"
	"webapp/utils/jwt"
	"webapp/utils/snowflake"
)

func SignUp(req gin.H) error{
	//1.查询数据库 和 业务条件逻辑
	isExist := mysql.CheckUserExist(req["username"].(string))
	if isExist {
		return errors.New("用户已存在")
	}
	userId := snowflake.GetID()
	user := &models.User{
		UserID: userId,
		Username: req["username"].(string),
		Password: helper.StrMd5Encode(req["password"].(string)),
	}

	if isInsertOk,err := mysql.InsertUser(user);isInsertOk {
		return nil
	}else {
		return err
	}
	// 构造一个user实例
	return nil
}

func Login(req gin.H) (string,error){
	user,err := mysql.Login(req["username"].(string),req["password"].(string))
	if err != nil{
		return "",err
	}
	token,err := jwt.GenToken(user.UserID,user.Username)
	if err != nil{
		return "",err
	}
	return token,err
	//fmt.Println("logic login",req)
}
