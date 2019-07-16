package main

import (
	"fmt"
	"os"
	"log"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {

					file, header , err := c.Request.FormFile("upload")       // 拿到上传的文件的信息 form-data格式
					if err != nil {
						log.Fatal(err)
					}
					filename := header.Filename
					fmt.Println(filename)
					fmt.Println(header.Filename)
					out, err := os.Create("./tmp/"+filename+".png")
					if err != nil {
							log.Fatal(err)
					}
					defer out.Close()
					_, err = io.Copy(out, file)                 // 拷贝上传的文件信息到新建的out文件中
					if err != nil {
							log.Fatal(err)
					}
	})
	router.Run(":8080")
}
