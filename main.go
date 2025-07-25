package main

import (
	"log"
	"os"
	"shopfood/component/appctx"
	"shopfood/component/uploadprovider"
	"shopfood/middleware"
	"shopfood/module/restaurant/transport/ginrestaurant"
	"shopfood/module/upload/transport/ginupload"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id   int     `json:"id" gorm:"column:id;"`
	Name string  `json:"name" gorm:"column:name;"`
	Addr *string `json:"Addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("MYSQL_STRING")
	if dsn == "" {
		log.Fatal("MYSQL_STRING environment variable is not set")
	}

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	s3Provioder := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appContext := appctx.NewAppContext(db, s3Provioder, secretKey)

	r := gin.Default()

	r.Use(middleware.Recover(appContext))

	v1 := r.Group("/v1")

	// API /upload file
	v1.POST("/upload", ginupload.Upload(appContext))

	restaurants := v1.Group("/restaurants")

	// API /restaurants
	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
	restaurants.GET("/:id", ginrestaurant.DetailRestaurant(appContext))
	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appContext))

	r.Run()

}
