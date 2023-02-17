package models

import (
	"shopping-cart/config"
	"shopping-cart/entities"
	"time"
	//"reflect"
)

type CartModel struct {
}

func (*CartModel) Order(cart []entities.Item, username string) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		now := time.Now().UTC()
		_, err2 := db.Exec("insert into orders (orderdate,orderstatus,staffid,customername) values (?,?,?,?)",
			now, "Processing", 1, username)
		if err2 != nil {
			return false
		}
		id := 1
		row, err3 := db.Query("select max(id) as id from orders")
		if err3 != nil {
			return false
		} else {
			for row.Next() {
				row.Scan(&id)
			}
		}
		var rowsAffected int64
		for _, item := range cart {
			results, errs := db.Exec("insert into orderdetail (orderid,productid,price,quantity) values (?,?,?,?)",
				id, item.Product.Id, item.Product.Price, item.Quantity)
			if errs != nil {
				return false
			} else {
				rowsAffected, _ = results.RowsAffected()
			}
		}
		return rowsAffected > 0
	}
}
