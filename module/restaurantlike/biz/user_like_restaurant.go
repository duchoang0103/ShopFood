package biz

import (
	"context"
	"log"
	"shopfood/common"
	restaurantlikemodel "shopfood/module/restaurantlike/model"
	"shopfood/pubsub"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

// type IncLikedCountResStore interface {
// 	IncreaseLikeCount(ctx context.Context, id int) error
// }

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	// incStore IncLikedCountResStore
	ps pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	// incStore IncLikedCountResStore,
	ps pubsub.Pubsub,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
		// incStore: incStore,
		ps: ps,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	//Send message
	if err := biz.ps.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	// //side effect ----- KHong can nua vi pubsub
	// job := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	// })

	// if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
	// 	log.Println(err) //Khong bat error o day tranh th api update bi block
	// }

	return nil
}
