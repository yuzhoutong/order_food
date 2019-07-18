package service

import (
	"encoding/json"
	"fmt"
	"order_food/models"
	"order_food/store"
)

func AutoSaveToDate() {

	defer func() {

		if err := recover(); err != nil {
			fmt.Println("AutoInsertlogToDB", err)
		}
	}()

	for {
		store.Rc.Brpop(store.CacheKeySystemLogs, func(bytes []byte) {
			var log models.SysLog
			if err := json.Unmarshal(bytes, &log); err != nil {
				fmt.Println("json unmarshal err", err)

			}
			if _, err := models.AddLogs(&log); err != nil {
				fmt.Println("insert to database err", err.Error(), log)
			}

		})
	}

}
