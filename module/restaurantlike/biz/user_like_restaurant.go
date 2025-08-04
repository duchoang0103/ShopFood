package biz

import (
	"context"
	"shopfood/common"
	restaurantlikemodel "shopfood/module/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncLikedCountResStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncLikedCountResStore
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	incStore IncLikedCountResStore,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:    store,
		incStore: incStore,
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

	go func() {
		defer common.AppRecover()

		biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId) //Khong bat error o day tranh th api update bi block
	}()

	return nil
}
