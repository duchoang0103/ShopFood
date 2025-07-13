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

	Detail(context context.Context, id int) (*restaurantmodel.Restaurant, error)
}
type detailRestaurantBiz struct {
	store detailRestaurantStore
}

func NewDetailRestaurantBiz(store detailRestaurantStore) *detailRestaurantBiz {
	return &detailRestaurantBiz{store: store}
}

func (biz *detailRestaurantBiz) DetailRestaurant(context context.Context, id int) (*restaurantmodel.Restaurant, error) {

	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	result, err := biz.store.Detail(context, id)
	if err != nil {
		return nil, common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}
	return result, nil

}
