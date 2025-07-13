package restaurantstorage

import (
	"context"
	"shopfood/common"
	restaurantmodel "shopfood/module/restaurant/model"
)

func (s *sqlStore) Update(context context.Context, newData *restaurantmodel.RestaurantUpdate, id int) error {

	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(newData).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
