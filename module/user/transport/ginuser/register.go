package ginuser

import (
	"net/http"

	"shopfood/common"
	"shopfood/component/appctx"
	"shopfood/component/hasher"
	userbiz "shopfood/module/user/biz"
	usermodel "shopfood/module/user/model"
	userstore "shopfood/module/user/store"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
