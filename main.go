package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gogin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var DB *gorm.DB

func main() {
	Run()
	ConnectDB()
}

func ConnectDB() {
	dsn := "test:@tcp(localhost:3306)/userInformation?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.UserInformation{})

}
func Run() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//로그인
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/login")
	})
	r.GET("/login", roadlogin)
	r.POST("/login/success", successLogin)

	//회원가입
	r.GET("/register", roadregister)
	r.POST("/register/success", successRegister)

	//서버 시작
	r.Run(":3000")
}

// 회원가입
func roadregister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func successRegister(c *gin.Context) {
	password := c.PostForm("password")
	userId := c.PostForm("userid")
	DB.Create(&models.UserInformation{
		UserID:    userId,
		Password:  password,
		CreatedAt: time.Now(),
	})
}

// 로그인
func roadlogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func successLogin(c *gin.Context) func(c *gin.Context) {
	userid := c.PostForm("userid")
	password := c.PostForm("password")
	DB.Select("userid").Where()
}
