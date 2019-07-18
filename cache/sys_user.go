package cache

import (
	"order_food/models"
	"order_food/store"
	"strconv"
	"encoding/json"
	"time"
)

//存取用户信息的缓存

func GetUserById(id int) (m *models.OrderUser, err error) {
	if store.Re == nil && store.Rc.IsExist(store.CacheKeyUserInfo+ "_"+strconv.Itoa(id)) {
		if data ,err1 := store.Rc.RedisBytes(store.CacheKeyUserInfo+ "_"+strconv.Itoa(id)) ; err1 == nil{
			err = json.Unmarshal(data, &m)
			if m != nil {
				return
			}
		}
	}
	m, err = models.GetUserById(id)
	if err != nil {
		return
	}
	if data ,err := json.Marshal(m); err == nil{
		store.Rc.Put(store.CacheKeyUserInfo+ "_"+strconv.Itoa(id), data , 1*time.Hour)
	}
	return
}