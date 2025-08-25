package subscriber

import (
	"context"
	"log"
	"shopfood/component/appctx"
	restaurantstorage "shopfood/module/restaurant/storage"
	"shopfood/pubsub"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	// GetUserId() int
}

// func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
// 	c, _ := appCtx.GetPubSub().Subscriber(ctx, common.TopicUserLikeRestaurant)

// 	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

// 	go func() {
// 		defer common.AppRecover()
// 		for {
// 			msg := <-c
// 			likeData := msg.Data().(HasRestaurantId)
// 			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
// 		}
// 	}()
// }

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user like restaurant",
		Hld: func(ctx context.Context, msg *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := msg.Data().(HasRestaurantId)

			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func PushNotificationWhenUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Push notification when user like restaurant",
		Hld: func(ctx context.Context, msg *pubsub.Message) error {
			likeData := msg.Data().(HasRestaurantId)
			log.Println("Push Notification When User Like Restaurant Id:", likeData.GetRestaurantId())

			return nil
		},
	}
}
