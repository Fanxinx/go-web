package mysql

import (
	"errors"
	"fmt"
	"webapp/models"
	"webapp/utils/helper"
)



func CheckUserExist(username string) bool{
	//sqlStr := `select count(user_id) from user where username = ?`
	//var count int
	var user models.User
	db.Where("username = ?",username).First(&user)
	if user.ID == 0{
		return false
	}
	fmt.Println("check user exist",username)
	return true
}

func InsertUser(user *models.User) (bool,error){
	//执行sql语句入库
	result := db.Create(&user)
	if result.Error != nil{
		return false,result.Error
	}else {
		return true,nil
	}
}


func Login(username string,password string) (models.User,error){
	var user models.User
	db.Where("username = ? and password = ?",username,helper.StrMd5Encode(password)).First(&user)
	if user.ID != 0 {
		return user,nil
	}else {
		return user,errors.New("用户账户密码错误")
	}
}
