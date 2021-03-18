package handler

import (
	"VinTravelAI/middleware"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func UploadImage(c *gin.Context) {
  session, err := middleware.Init()
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot create new session")
    fmt.Println(err)
    return
  }

  err = godotenv.Load(".env")
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot load env file")
    fmt.Println(err)
    return
  }

  hours, minutes, second := time.Now().Clock()
  newFileName := strconv.Itoa(hours) + strconv.Itoa(minutes) + strconv.Itoa(second) + ".png"

  fileImage, _, err := c.Request.FormFile("image")
  if err != nil {
    c.JSON(http.StatusBadRequest, "Cannot get image")
    fmt.Println(err)
    return
  }
  defer fileImage.Close()
  newFile, err := os.Create("./AI Model/" + newFileName)
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot create new file to save on server")
    fmt.Println(err)
    return
  }
  defer newFile.Close()
  _, err = io.Copy(newFile, fileImage)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err.Error)
    fmt.Println(err)
    return
  }

  url, err := middleware.Upload(session, "./AI Model/" + newFileName)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err.Error())
    fmt.Println(err)
    return
  }
  c.JSON(http.StatusOK, url)
  fmt.Println("New file url: ", url)
}
