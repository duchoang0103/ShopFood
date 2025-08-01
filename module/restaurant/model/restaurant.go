package restaurantmodel

import (
	"errors"
	"shopfood/common"
	"strings"
)

type RestaurantType string

const TypeNormal RestaurantType = "Normal"
const TypePremium RestaurantType = "Premium"

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	Addr            string             `json:"Addr" gorm:"column:addr;"`
	Typer           RestaurantType     `json:"type" gorm:"column:type;"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
	OwnerId         int                `json:"-" gorm:"column:owner_id;"`
	User            *common.SimpleUser `json:"user" gorm:"foreignKey:OwnerId;references:Id;preload:false;"`
	LikedCount      int                `json:"liked_count" gorm:"-"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) Mask(isAdminOwner bool) {
	r.GenUID(common.DbTypeRestaurant)

	if u := r.User; u != nil {
		u.Mask(isAdminOwner)
	}
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            *string        `json:"Addr" gorm:"column:addr;"`
	OwnerId         int            `json:"-" gorm:"column:owner_id;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (data *RestaurantCreate) Mask(isAdminOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"Addr" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

var (
	ErrNameIsEmpty = errors.New("Name can't is empty")
)
