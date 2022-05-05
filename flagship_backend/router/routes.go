package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
	"xiangerfer.com/db_flagship/controller"
	"xiangerfer.com/db_flagship/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine{
	/* 用户信息 */
	r.POST("/api/v1/auth/register", controller.Register)
	r.POST("/api/v1/auth/login", controller.Login)
	r.GET("/api/v1/auth/info", middleware.AuthMiddleware(), controller.Info)

	/* 主机API */
	r.GET("/api/v1/auth/hosts", controller.HostInfo)
	//r.POST("/api/v1/auth/login", controller.Login)
	//r.GET("/api/v1/auth/info", middleware.AuthMiddleware(), controller.Info)


	/* 测试api参数
	1. :name 完全匹配
	2. :action 剩下的全匹配
	3. *other 全匹配
	*/
	r.GET("/api/v1/user/:name/:action/*other", controller.ApiParam)

	/*
		测试url参数
		请求： 127.0.0.1:8081/api/v1/user/url?username=yixiang&password=helloworld
	*/
	r.GET("/api/v1/user/url", controller.UrlQuery)


	/*
		测试表单参数
	*/
	r.POST("/api/v1/user/formpost", controller.FormPostParam)


	/*
		 测试提交文件
	*/
	r.POST("/api/v1/user/formpostfile", controller.FormPostFile)


	/*
		测试404页面
	*/
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"msg": "not found",
		})
	})

	/*
		测试xml 数据
	*/
	r.GET("/api/v1/someXML", func(c *gin.Context) {
		/*
		<map>
		    <message>abc</message>
		</map>
		*/
		c.XML(200, gin.H{"message": "abc"})
	})

	// yaml
	r.GET("/api/v1/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "zhangsan", "age": 19})
		//c.ProtoBuf()
	})


	// html 渲染
	cur_dir, _ := os.Getwd()
	// 加载所有的html文件
	//r.LoadHTMLGlob(cur_dir + "/static/*")
	r.LoadHTMLFiles(cur_dir + "/static/index.html")
	r.GET("/api/v1/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
	})


	// 重定向
	r.GET("/api/v1/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// 同步，异步
	// 1.异步， 这个可以用在sql执行。先返回，然后执行进度写入到数据库里面。
	r.GET("/api/v1/long_async", func(c *gin.Context) {
		// 需要搞一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})
	// 2.同步
	r.GET("/api/v1/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
	})

	// 局部中间件
	//局部中间键使用
	r.GET("/api/v1/mid3", MiddleWare3(), func(c *gin.Context) {
		// 取值
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		// 页面接收
		c.JSON(200, gin.H{"request中间件3": req})
	})


	// 中间件，全局定义，在上面的不会收到这个影响，下面全部会经过这个。可以通过中间件判断是否认证，没有则重定向到登陆
	r.Use(MiddleWare())
	r.GET("/api/v1/mid", func(c *gin.Context) {
		// 取值
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		// 页面接收
		c.JSON(200, gin.H{"request": req})
	})

	// next()方法，可以把要执行的函数放中间执行。
	r.Use(MiddleWare2())
	r.GET("/api/v1/mid2", func(c *gin.Context) {
		//
		fmt.Println("函数执行")
		// 页面接收
		c.JSON(200, gin.H{"request": "mid2"})
	})



	return r
}

// 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}


func MiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件2开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件2执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

// 局部中间件
func MiddleWare3() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件3开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件3执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}