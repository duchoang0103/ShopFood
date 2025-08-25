package subscriber

import (
	"context"
	"log"
	"shopfood/common"
	"shopfood/component/appctx"
	restaurantstorage "shopfood/module/restaurant/storage"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	// GetUserId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscriber(ctx, common.TopicUserLikeRestaurant)

	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId)
			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}

func PushNotificationWhenUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscriber(ctx, common.TopicUserLikeRestaurant)

	// store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId)
			// _ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
			log.Println("Push Notification When User Like Restaurant", likeData)
		}
	}()
}
