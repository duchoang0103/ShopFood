package ginrstlike

import (
	"net/http"
	"shopfood/common"
	"shopfood/component/appctx"
	retlikebiz "shopfood/module/restaurantlike/biz"
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

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := retlikebiz.NewUserDislikeRestaurantBiz(store)

		if err := biz.DislikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
