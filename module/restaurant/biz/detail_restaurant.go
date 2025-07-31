package restaurantbiz

import (
	"context"
	"shopfood/common"
	restaurantmodel "shopfood/module/restaurant/model"
)

type detailRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}
type detailRestaurantBiz struct {
	store detailRestaurantStore
}

func NewDetailRestaurantBiz(store detailRestaurantStore) *detailRestaurantBiz {
	return &detailRestaurantBiz{store: store}
}

func (biz *detailRestaurantBiz) DetailRestaurant(context context.Context, id int, morekeys ...string) (*restaurantmodel.Restaurant, error) {

	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id}, morekeys...)
	if err != nil {
		return nil, common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	return oldData, nil
}
