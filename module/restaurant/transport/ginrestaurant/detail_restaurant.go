package ginrestaurant

import (
	"shopfood/common"
	"shopfood/component/appctx"
	restaurantbiz "shopfood/module/restaurant/biz"
	restaurantstorage "shopfood/module/restaurant/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DetailRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDetailRestaurantBiz(store)

		result, err := biz.DetailRestaurant(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(201, common.OnlySuccessResponse(result))
	}

}
