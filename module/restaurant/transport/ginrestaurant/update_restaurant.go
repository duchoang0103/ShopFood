package ginrestaurant

import (
	"shopfood/common"
	"shopfood/component/appctx"
	restaurantbiz "shopfood/module/restaurant/biz"
	restaurantmodel "shopfood/module/restaurant/model"
	restaurantstorage "shopfood/module/restaurant/storage"

	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var newData restaurantmodel.RestaurantUpdate

		if err := c.ShouldBindJSON(&newData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), &newData, id); err != nil {
			panic(err)
		}

		c.JSON(201, common.OnlySuccessResponse(true))
	}

}
