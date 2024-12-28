package main

import (
	AuthController "changpab/jwt-api/controller/auth"
	UserController "changpab/jwt-api/controller/user"
	"changpab/jwt-api/middleware"
	_ "changpab/jwt-api/middleware"
	"changpab/jwt-api/orm"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Lastname string
	Email    string
}

type Image struct {
	gorm.Model
	Post_ID int
	Img_url string
}

type Post struct {
	gorm.Model
	Post_ID string
	User_ID string
	Detail  string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	orm.InitDB()

	r := gin.Default()

	// ให้เซิร์ฟเวอร์สามารถเข้าถึงไฟล์ในโฟลเดอร์ ./uploads ได้
	r.Static("/get_image", "./uploads/user_profile") // เพิ่มบรรทัดนี้

	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)

	authorized := r.Group("/users", middleware.JWTAuthen())
	authorized.GET("/readall", UserController.PGAll)
	authorized.GET("/profile", UserController.Profile)
	authorized.GET("/upload_imagepost", UserController.Uploadimage)
	authorized.GET("/upload_imageprofile", UserController.Uploadimage_profile)
	authorized.GET("/profile_imge", UserController.Profile_Img)

	device_host := os.Getenv("DEVICE_HOST")

	r.Run(device_host + ":8080")
}
