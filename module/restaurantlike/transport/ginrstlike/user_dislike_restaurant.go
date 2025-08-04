package ginrstlike

import (
	"net/http"
	"shopfood/common"
	"shopfood/component/appctx"
	restaurantstorage "shopfood/module/restaurant/storage"
	retlikebiz "shopfood/module/restaurantlike/biz"
	restaurantlikemodel "shopfood/module/restaurantlike/model"
	restaurantlikestorage "shopfood/module/restaurantlike/store"

	"github.com/gin-gonic/gin"
)

// DELETE /v1/restaurants/:id/unlike
func UserDislikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		deStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := retlikebiz.NewUserDislikeRestaurantBiz(store, deStore)

		if err := biz.DislikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
