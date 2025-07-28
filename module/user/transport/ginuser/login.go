package ginuser

import (
	"net/http"
	"shopfood/common"

	"github.com/gin-gonic/gin"

	"shopfood/component/appctx"
	"shopfood/component/hasher"
	"shopfood/component/tokenprovider/jwt"
	userbiz "shopfood/module/user/biz"
	usermodel "shopfood/module/user/model"
	userstore "shopfood/module/user/store"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
