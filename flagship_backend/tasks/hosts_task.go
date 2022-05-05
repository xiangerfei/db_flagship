package tasks

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"time"
	"xiangerfer.com/db_flagship/common"

)





func HostsTask(){

	common.InitConfig()
	cur_dir, _ := os.Getwd()
	log_file_name := viper.GetString("server.tasks_log_file")
	fmt.Println(log_file_name)
	log_file_split_interval := viper.GetString("server.tasks_log_split_interval")
	fmt.Println(log_file_split_interval)
	var date_f string
	if strings.ToUpper(log_file_split_interval) == "HOUR"{
		date_f = time.Now().Format(common.DATE_FORMATE_D)
	}else{
		date_f = time.Now().Format(common.DATE_FORMATE_D)
	}

	// 日志文件处理
	var log_file_abs_low string
	if strings.HasSuffix(log_file_name, ".log") || strings.HasSuffix(log_file_name, ".txt"){
		// 判断日志文件的格式 filename or filename.log 如果是.txt 或者.log是的话截开里面的串。
		fmt.Println("filename:",  log_file_name)
		file_name_prefix := strings.Split(log_file_name, ".")[0]
		file_name_suffix := strings.Split(log_file_name, ".")[1]
		fmt.Println(log_file_name)
		log_file_abs_low =  cur_dir + "/" + file_name_prefix + "-" + date_f + "." + file_name_suffix
	}else{
		fmt.Println("filename:",  log_file_name)
		// 如果不是.log或者.txt结尾
		log_file_abs_low = cur_dir + "/" + log_file_name + "-" + date_f + ".log"
	}

	log_file, _ := os.OpenFile(log_file_abs_low, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer log_file.Close()
	logger := common.Logger{}

	logger.SetOutput(log_file)

	logger.SetFlags(log.LstdFlags | log.Lshortfile | log.Ltime)


	tc := time.NewTicker(3 * time.Second)
	defer tc.Stop()
	i := 0
	for {
		<- tc.C
		logger.Info("普通信息 %d", i)
		logger.Warning("警告信息 %d", i)
		logger.Error("错误信息 %d", i)
		i++
	}

}

