package biz

import (
	"context"

	"shopfood/common"
	restaurantlikemodel "shopfood/module/restaurantlike/model"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, data *restaurantlikemodel.Like) error
}

type DecLikedCountResStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userDislikeRestaurantBiz struct {
	store    UserDislikeRestaurantStore
	decStore DecLikedCountResStore
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore, decStore DecLikedCountResStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{
		store:    store,
		decStore: decStore,
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

	go func() {
		defer common.AppRecover()

		biz.decStore.DecreaseLikeCount(ctx, data.RestaurantId) //Khong bat error o day tranh th api update bi block
	}()

	return nil
}
