package ginrstlike

import (
	"net/http"
	"shopfood/common"
	"shopfood/component/appctx"

	"github.com/gin-gonic/gin"

	restaurantstorage "shopfood/module/restaurant/storage"
	rstlikebiz "shopfood/module/restaurantlike/biz"
	restaurantlikemodel "shopfood/module/restaurantlike/model"
	restaurantlikestorage "shopfood/module/restaurantlike/store"
)

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			common.ErrInvalidRequest(err)
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		incStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewUserLikeRestaurantBiz(store, incStore)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, common.ErrInternal(err))
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
