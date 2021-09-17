package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	*App		   `mapstructure:"app"`
	*LogConfig 	   `mapstructure:"log"`
	*MySQLConfig   `mapstructure:"mysql"`
	*RedisConfig   `mapstructure:"redis"`
}

type App struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	Port      int    `mapstructure:"port"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}

type LogConfig struct {
	Level		string	   `mapstructure:"level"`
	Filename 	string	   `mapstructure:"filename"`
	MaxSize		int		   `mapstructure:"max_size"`
	MaxAge		int		   `mapstructure:"max_age"`
	MaxBackups  int		   `mapstructure:"max_backups"`
}

type MySQLConfig struct{
	Host			string		`mapstructure:"host"`
	Driver			string		`mapstructure:"driver"`
	Port			int			`mapstructure:"port"`
	User			string		`mapstructure:"user"`
	Password		string		`mapstructure:"password"`
	DbName			string		`mapstructure:"db_name"`
	DbCharset		string		`mapstructure:"db_charset"`
	DbPrefix		string		`mapstructure:"db_prefix"`
	DbCollation		string		`mapstructure:"db_collation"`
	MaxIdleConns	string		`mapstructure:"max_idle_conns"`
	MaxOpenConns	string		`mapstructure:"max_open_conns"`
}

type RedisConfig struct {
	Host		string		`mapstructure:"host"`
	Port		int			`mapstructure:"port"`
	Password	string		`mapstructure:"password"`
	Db			int			`mapstructure:"db"`
	PoolSize	int			`mapstructure:"pool_size"`
}

func Init()(err error) {
	//viper.SetConfigFile("./conf/config.yaml")// 多文件指定
	viper.SetConfigName("config") //全局查找config的文件 不要有同名的配置文件
	//viper.SetConfigType("yaml") //专注远程读取类型 etcd
	viper.AddConfigPath(".") //指定配置文件的查找路径 （这里使用相对路径）
	viper.AddConfigPath("./conf")
	err = viper.ReadInConfig()
	if err != nil{
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n",err)
		return err
	}
	//将读取到的变量反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf);err != nil{
		fmt.Printf("viper.Unmarshal failed,err:%v\n",err)
	}else {
		fmt.Println("viper.Unmarshal success")
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config.yaml has been changed")
		if err := viper.Unmarshal(Conf);err != nil{
			fmt.Printf("viper.Unmarshal failed,err:%v\n",err)
		}
	})
	return nil
}
