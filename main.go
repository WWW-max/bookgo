package main

import (
	"bookgo/database"
	"bookgo/routes"
	"log"
)


func main() {
	// 初始化数据库
	database.InitDatabase()
	defer database.DB.Close()

	// 设置路由
	router := routes.SetupRouter()

	// 启动服务器
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}