package main

import (
	"VinTravelAI/handler"
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.POST("/image/upload", handler.UploadImage)

  router.Run()
}
