package restaurantstorage

import (
	"context"
	"shopfood/common"
	restaurantmodel "shopfood/module/restaurant/model"

	"gorm.io/gorm"
)

func (s *sqlStore) Update(context context.Context, newData *restaurantmodel.RestaurantUpdate, id int) error {

	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(newData).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
