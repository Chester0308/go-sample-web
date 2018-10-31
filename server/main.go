package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sample-web/server/controller"
)

func main() {
	// default ミドルウェアと共に、gin router を作成
	router := gin.Default()

	// 各 method の指定方法
	//router.GET("/someGet", getting)
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	router.GET("/", controller.GetIndex)
	router.GET("/fullname", controller.GetFullName)
	router.POST("/message", controller.PostMessage)
	router.GET("/set-cookie", controller.SetCookie)
	router.GET("/path-param/:id/:name", controller.SetCookie)

	// group: auth
	// ミドルウェアに、BasicAuth を設定する
	auth := router.Group("/auth", gin.BasicAuth(gin.Accounts{
		"admin": "secret",
		"hoge": "1234",
	}))
	auth.GET("/admin", controller.BasicAuth)

	// group: sample
	//sample := router.Group("/sample")
	//{
	//	sample.GET("/", controller.GetIndex)
	//	sample.GET("/fullname", controller.GetFullName)
	//}

	// PORT 環境変数が設定されている場合は PORT
	// router.Run(":8080") のように、ハードコーディングできる
	// default は 8080 を listen
	router.Run()
}
