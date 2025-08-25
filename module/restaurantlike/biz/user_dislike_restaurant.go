package biz

import (
	"context"
	"log"

	"shopfood/common"
	restaurantlikemodel "shopfood/module/restaurantlike/model"
	"shopfood/pubsub"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, data *restaurantlikemodel.Like) error
}

// type DecLikedCountResStore interface {
// 	DecreaseLikeCount(ctx context.Context, id int) error
// }

type userDislikeRestaurantBiz struct {
	store UserDislikeRestaurantStore
	// decStore DecLikedCountResStore
	ps pubsub.Pubsub
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore,
	// decStore DecLikedCountResStore,
	ps pubsub.Pubsub,
) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{
		store: store,
		// decStore: decStore,
		ps: ps,
	}
}

func (biz *userDislikeRestaurantBiz) DislikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := biz.store.Delete(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	//Send message
	if err := biz.ps.Publish(ctx, common.TopicUserDisLikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	// // side effect ----- KHong can nua vi pubsub
	// job := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.decStore.DecreaseLikeCount(ctx, data.RestaurantId)
	// })

	// if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
	// 	log.Println(err) //Khong bat error o day tranh th api update bi block
	// }

	return nil
}
