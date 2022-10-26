package main

import (
	"fmt"
	"gin-gorm/migrations"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PASS"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_NAME"))

	fmt.Println("dsn:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := InitializeDB()
	errMigration := migrations.InitializeMigration(db)
	if errMigration != nil {
		panic(errMigration.Error())
	}

	r := gin.Default()

	r.Use(requestid.New())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "id:"+requestid.Get(c))
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
