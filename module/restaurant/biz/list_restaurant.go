package restaurantbiz

import (
	"context"
	"shopfood/common"
	restaurantmodel "shopfood/module/restaurant/model"
)

type ListRestaurantRepo interface {
	ListRestaurantRepo(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging, morekeys ...string) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewlistRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo: repo}
}

func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging, morekeys ...string) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.repo.ListRestaurantRepo(context, filter, paging, morekeys...)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
