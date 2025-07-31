package ginrestaurant

import (
	"shopfood/common"
	"shopfood/component/appctx"
	restaurantbiz "shopfood/module/restaurant/biz"
	restaurantmodel "shopfood/module/restaurant/model"
	restaurantstorage "shopfood/module/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var newRestaurant restaurantmodel.RestaurantCreate

		if err := c.ShouldBindJSON(&newRestaurant); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		newRestaurant.OwnerId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &newRestaurant); err != nil {
			panic(err)
		}

		newRestaurant.Mask(false)

		c.JSON(201, common.SimpleSuccessResponse(newRestaurant.FakeId.String()))
	}

}
