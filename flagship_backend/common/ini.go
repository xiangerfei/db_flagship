package common

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

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