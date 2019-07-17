package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// 处理静态文件请求
func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.Run(":8081")
}