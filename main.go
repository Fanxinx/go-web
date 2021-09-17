package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webapp/dao/mysql"
	"webapp/dao/redis"
	"webapp/logger"
	"webapp/routes"
	"webapp/settings"
	"webapp/utils/snowflake"
	"webapp/utils/validator"
)

// Go web 脚手架模板
func main() {
	//1.加载配置文件
	if err := settings.Init();err != nil{
		fmt.Printf("init setttings failed,err:%v\n",err)
		return
	}
	//2.初始化日志
	if err := logger.Init(viper.GetString("app.mode"));err != nil{
		fmt.Printf("init logger failed,err:%v\n",err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success ...")
	//3.初始化Mysql
	mysql.Init()
	defer mysql.Close()
	//4.初始化Redis
	if err := redis.Init();err != nil{
		fmt.Printf("init redis failed,err:%v\n",err)
		return
	}
	defer redis.Close()
	//5.注册路由
	r := routes.Setup()
	//6.初始化验证翻译器
	validator.Init("zh")
	//7.初始化ID
	snowflake.Init(viper.GetString("app.start_time"),viper.GetInt64("app.machine_id"))
	//8.优雅关机重启
	srv := &http.Server{
		Addr:fmt.Sprintf(":%d",viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed{
			zap.L().Info("listen: " + err.Error())
			//log.Fatal("listen: %s\n",err)
		}
	}()

	quit := make(chan os.Signal,1)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	zap.L().Info("shutdown server ...")
	//log.Println("shutdown server ...")
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx);err != nil{
		zap.L().Fatal("server shutdown: ",zap.Error(err))
	}
	zap.L().Info("server exiting")
}