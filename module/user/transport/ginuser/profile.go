package ginuser

import (
	"net/http"
	"shopfood/common"
	"shopfood/component/appctx"

	"github.com/gin-gonic/gin"
)

func Profile(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		// u là con trỏ User mà đã implement Requester interface => có thể dùng nó để gọi các hàm trong Requester interface
		u := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
