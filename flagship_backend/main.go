package main

//import "xiangerfer.com/db_flagship/model"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"time"
	"xiangerfer.com/db_flagship/common"
	"xiangerfer.com/db_flagship/cron"
	"xiangerfer.com/db_flagship/router"
)


func main(){

	// 热加载
	// go get -v -u github.com/pilu/fresh
	// 读取配置
	InitConfig()
	db := common.InitDB()
	go func(){
		for{
			if err := db.DB().Ping(); err != nil{
				fmt.Println("有问题")
			}else{
				fmt.Println("链接没问题")
			}
			time.Sleep(time.Second * 600)
		}

	}()

	defer db.Close()

	// 设置模式
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r = router.CollectRoute(r)
	listenAddr := viper.GetString("server.listenAddr")
	listenPort := viper.GetString("server.listenPort")

	// 周期任务
	// go tasks.HostsTask()
	// 定时任务
	go cron.CronTask()

	// Logging to a file.

	//gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)



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