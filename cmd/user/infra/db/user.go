package db

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID        int64  `gorm:"column:user_id"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

func (u User) TableName() string {
	return "t_user"
}

func CreateUser(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Create(user).Error
}

func MGetUser(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("user_id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryUser(ctx context.Context, username string) ([]*User, error) {
	res := make([]*User, 0)
	err := DB.WithContext(ctx).Where("username = ?", username).Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return res, nil
}
