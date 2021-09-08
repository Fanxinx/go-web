package mysql

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)
//
//const (
//	DB_DRIVER	 = "mysql"
//	DB_HOST	     = "127.0.0.1"
//	DB_PORT	     = "3306"
//	DB_DATABASE  = "opgg"
//	DB_USERNAME  = "root"
//	DB_PASSWORD  = "asd456789"
//	DB_CHARSET   = "utf8mb4"
//	DB_COLLATION = "utf8mb4_unicode_ci"
//	DB_PREFIX    = "opgg_"
//	SET_MAX_IDLE_CONNS = 20
//	SET_MAX_OPEN_CONNS = 100
//)
var db *gorm.DB
func Init()(err error){
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local&collation=%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.db_name"),
		viper.GetString("mysql.db_charset"),
		viper.GetString("mysql.db_collation"))

	db,err = gorm.Open(viper.GetString("mysql.driver"),args)

	if err != nil{
		zap.L().Error("fail connect db",zap.Error(err))
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return viper.GetString("mysql.db_prefix") + defaultTableName
	}

	db.SingularTable(true)
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	db.DB().SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	//db.DB().SetConnMaxLifetime(time.Hour)
	//设置表面前缀
	//gorm.DefaultTableNameHandler = initTablePrefix

	//DB = db

	//sqlDB := db.DB()
	//// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	//sqlDB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	//// SetMaxOpenConns 设置打开数据库连接的最大数量。
	//sqlDB.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	//// SetConnMaxLifetime 设置了连接可复用的最大时间。
	//sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
	//return db
}

func Close(){
	_ = db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// deleteCallback will set `DeletedOn` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}


