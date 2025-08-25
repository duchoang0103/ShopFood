package subscriber

import (
	"context"
	"shopfood/component/appctx"
)

func Setup(appCtx appctx.AppContext, ctx context.Context) {
	IncreaseLikeCountAfterUserLikeRestaurant(appCtx, ctx)
	DecreaseLikeCountAfterUserDisLikeRestaurant(appCtx, ctx)
	PushNotificationWhenUserLikeRestaurant(appCtx, ctx)
}
