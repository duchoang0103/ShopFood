package main

import (
	"context"
	"log"
	"os"
	"shopfood/component/appctx"
	"shopfood/component/uploadprovider"
	"shopfood/middleware"
	"shopfood/pubsub/localpb"
	"shopfood/subscriber"

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
	ps := localpb.NewPubSub()
	appContext := appctx.NewAppContext(db, s3Provioder, secretKey, ps)

	// setup subscribers
	subscriber.Setup(appContext, context.Background())

	r := gin.Default()

	r.Use(middleware.Recover(appContext))

	v1 := r.Group("/v1")

	SetupRouter(appContext, v1)
	SetupAdminRouter(appContext, v1)

	r.Run()

}
