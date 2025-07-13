package main

import (
	"log"
	"os"
	"shopfood/component/appctx"
	"shopfood/middleware"
	"shopfood/module/restaurant/transport/ginrestaurant"

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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	appContext := appctx.NewAppContext(db)

	r := gin.Default()

	r.Use(middleware.Recover(appContext))

	v1 := r.Group("/v1")

	restaurants := v1.Group("/restaurants")

	// API /restaurants
	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
	restaurants.GET("/:id", ginrestaurant.DetailRestaurant(appContext))
	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appContext))

	r.Run()

}
