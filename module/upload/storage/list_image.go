package storage

import (
	"context"
	"shopfood/common"
)

func (store *sqlStore) ListImages(
	context context.Context,
	ids []int,
	moreKeys ...string,
) ([]common.Image, error) {
	db := store.db
	var result []common.Image

	db = db.Table(common.Image{}.TableName())

	if err := db.Where("id in (?)", ids).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
