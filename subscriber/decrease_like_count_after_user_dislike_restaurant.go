package subscriber

import (
	"context"
	"shopfood/component/appctx"
	restaurantstorage "shopfood/module/restaurant/storage"
	"shopfood/pubsub"
)

// func DecreaseLikeCountAfterUserDisLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
// 	c, _ := appCtx.GetPubSub().Subscriber(ctx, common.TopicUserDisLikeRestaurant)

// 	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

// 	go func() {
// 		defer common.AppRecover()
// 		for {
// 			msg := <-c
// 			likeData := msg.Data().(HasRestaurantId)
// 			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
// 		}
// 	}()
// }

func DecreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user dislike restaurant",
		Hld: func(ctx context.Context, msg *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := msg.Data().(HasRestaurantId)

			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
