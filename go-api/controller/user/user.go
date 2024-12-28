package user

import (
	"changpab/jwt-api/orm"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func PGAll(c *gin.Context) {
	device_host := os.Getenv("DEVICE_HOST")
	var imageHost = "http://" + device_host + ":8080/get_image/"
	var users []orm.User
	// orm.Db.Where("Role = ?", "PG").Find(&users)
	orm.Db.Find(&users, "Role = ?", "PG")
	for i := range users {
		users[i].Img_profile = imageHost + users[i].Img_profile
		// fmt.Println("%s", users[i].Img_profile)
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "userId": users})
}

func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user orm.User
	orm.Db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "userId": user})
}

func Profile_Img(c *gin.Context) {
	device_host := os.Getenv("DEVICE_HOST")
	var imageHost = "http://" + device_host + ":8080/get_image/"
	userId := c.MustGet("userId").(float64)
	var user orm.User
	orm.Db.First(&user, userId)
	fmt.Println(user.Img_profile)
	fmt.Println(userId)
	imageURL := fmt.Sprintf(imageHost+"%s", user.Img_profile)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "image_url": imageURL})
}

type ImageBody struct {
	User_ID int    `json:"user_id" binding:"required"`
	Img_url string `json:"img_url" binding:"required"`
}

func Uploadimage(c *gin.Context) {
	userId := c.MustGet("userId").(float64)

	// รับไฟล์จาก form-data
	file, _ := c.FormFile("file") // 'file' คือ key ที่ส่งมาจาก client

	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No file found"})
		return
	}

	// ตั้งชื่อไฟล์ใหม่ หรือจัดการชื่อไฟล์ตามที่ต้องการ
	fileName := fmt.Sprintf("%d_%s", int(userId), file.Filename)
	filePath := fmt.Sprintf("./uploads/%s", fileName)

	fmt.Println(fileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to save file"})
		return
	}

	image := orm.Image{
		Post_ID: int(userId),
		Img_url: fileName,
	}
	orm.Db.Create(&image)
	if image.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Create Success", "userId": image.ID})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Create Failed"})
	}

	// ส่ง response
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Profile image uploaded successfully",
		"file":    fileName,
	})
}

func Uploadimage_profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)

	// รับไฟล์จาก form-data
	file, _ := c.FormFile("file") // 'file' คือ key ที่ส่งมาจาก client

	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No file found"})
		return
	}

	// ตั้งชื่อไฟล์ใหม่ หรือจัดการชื่อไฟล์ตามที่ต้องการ
	fileName := fmt.Sprintf("%d_%s", int(userId), file.Filename)
	filePath := fmt.Sprintf("./uploads/user_profile/%s", fileName)

	fmt.Println(fileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to save file"})
		return
	}

	var user orm.User
	user.ID = uint(userId)
	user.Img_profile = fileName
	if err := orm.Db.Model(&user).Select("Img_profile").Updates(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update user profile"})
		return
	}

	// ส่ง response
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Profile image uploaded successfully",
		"file":    fileName,
	})
}
