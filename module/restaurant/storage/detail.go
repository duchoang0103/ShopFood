package restaurantstorage

import (
	"context"
	"shopfood/common"
	restaurantmodel "shopfood/module/restaurant/model"
)

func (s *sqlStore) Detail(context context.Context, id int) (*restaurantmodel.Restaurant, error) {

	var result restaurantmodel.Restaurant

	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		First(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
