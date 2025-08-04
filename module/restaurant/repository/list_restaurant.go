package repository

import (
	"context"
	"shopfood/common"
	restaurantmodel "shopfood/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type LikeRestaurantStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}
type listRestaurantRepo struct {
	store ListRestaurantStore
}

func NewlistRestaurantRepo(store ListRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store}
}

func (biz *listRestaurantRepo) ListRestaurantRepo(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging, morekeys ...string) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.store.ListDataWithCondition(context, filter, paging, morekeys...)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	// ids := make([]int, len(result))

	// for i := range result {
	// 	ids[i] = result[i].Id
	// }

	// likeMap, err := biz.likeStore.GetRestaurantLikes(context, ids)

	// if err != nil {
	// 	log.Println("err: ", err)
	// 	return result, nil
	// }

	// for i, item := range result {
	// 	result[i].LikedCount = likeMap[item.Id]
	// }

	return result, nil
}
