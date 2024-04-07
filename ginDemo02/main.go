package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// go run github.com/pilu/fresh

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "值:%v", "首页")
	})

	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好Gin",
		})
	})

	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "你好Gin--22",
		})
	})

	r.GET("/json3", func(c *gin.Context) {

		a := &Article{
			Title:   "我是一个标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSON(200, a)
	})

	//响应Jsonp请求
	// http://localhost:8080/jsonp?callback=xxxx
	// xxxx({"title":"我是一个标题-jsonp","desc":"描述","content":"测试内容"});
	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article{
			Title:   "我是一个标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSONP(200, a)
	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "你好Gin--我是一个xml",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台的数据",
			//"msg":     "你好Gin--我是一个xml",
		})
	})

	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是后台的数据goods",
			//"msg":     "你好Gin--我是一个xml",
			"price": 20,
		})
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
