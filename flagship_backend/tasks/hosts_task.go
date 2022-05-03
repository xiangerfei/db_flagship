package tasks

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
	"xiangerfer.com/db_flagship/common"
)



func Hosts_task(){
	cur_dir, _ := os.Getwd()
	log_file_abs := cur_dir + "/" + viper.GetString("server.tasks_log_file")
	log_file, _ := os.OpenFile(log_file_abs, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	logger := common.Logger{}

	logger.SetOutput(log_file)
	logger.SetPrefix("db_flagship:")
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