package restaurantbiz

import (
	"context"
	"shopfood/common"
	restaurantmodel "shopfood/module/restaurant/model"
)

type deleteRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	Delete(context context.Context, id int) error
}
type deleteRestaurantBiz struct {
	store     deleteRestaurantStore
	requester common.Requester
}

func NewDeleteRestaurantBiz(store deleteRestaurantStore, requester common.Requester) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store, requester: requester}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {

	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	if oldData.OwnerId != biz.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
