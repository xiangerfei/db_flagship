package main

//import "xiangerfer.com/db_flagship/model"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"os"
	"xiangerfer.com/db_flagship/common"
	"xiangerfer.com/db_flagship/tasks"
)


func main(){

	// 读取配置
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	listenAddr := viper.GetString("server.listenAddr")
	listenPort := viper.GetString("server.listenPort")

	// 任务
	go tasks.Hosts_task()

	if listenPort != ""{
		panic(r.Run(fmt.Sprintf("%s:%s", listenAddr, listenPort)))
	}
	panic(r.Run(fmt.Sprintf("%s:%s", listenAddr, "8081")))
}

func InitConfig()  {
	wordDir, _ := os.Getwd()
	fmt.Printf("workdir: %s", wordDir)
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(wordDir + "/config")
	fmt.Println(viper.ConfigFileUsed())
	err := viper.ReadInConfig()

	if err != nil{
		log.Printf("read config failed ,err : %v", err)
		panic("read config failed, err: " + err.Error())
	}


} 
