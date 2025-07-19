package ginrestaurant

import (
	"shopfood/common"
	"shopfood/component/appctx"
	restaurantbiz "shopfood/module/restaurant/biz"
	restaurantmodel "shopfood/module/restaurant/model"
	restaurantstorage "shopfood/module/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.Status = []int{1}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewlistRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(201, common.NewSuccessResponse(result, pagingData, filter))
	}

}
