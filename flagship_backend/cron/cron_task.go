/* 测试一下定时任务 */
package cron

import (
	"fmt"
	"github.com/robfig/cron"
)


func CronTask(){
	c := cron.New()
	c.AddFunc("0 20 17 * * *", func() {
		fmt.Println("17:20 hello world")
	})

	c.Start()
}

