package data

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/sjw/user/model"
)

type userRepo struct {
	data *Data
}

func (ur *userRepo) Register(ctx context.Context, user *model.User) error {

	result := ur.data.db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	//write back to redis
	ur.AddToCache(user)

	return nil
}

func (ur *userRepo) ChangeUserName(ctx context.Context, userId int64, newName string) error {
	return nil
}

func (ur *userRepo) AddToCache(user *model.User) {

	ur.RemoveFromCache(user.User_Id)

	userInfo, _ := json.Marshal(user)

	err := ur.data.rdb.Set(strconv.Itoa(int(user.User_Id)), string(userInfo), 1000).Err()

	if err != nil {
		panic(err)
	}
}

func (ur *userRepo) RemoveFromCache(uid int64) {
	ur.data.rdb.Del(strconv.Itoa(int(uid))).Result()
}
