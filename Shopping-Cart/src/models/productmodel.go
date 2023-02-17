package models

import (
	"database/sql"
	"shopping-cart/config"
	"shopping-cart/entities"
)

type ProductModel struct {
}

func (*ProductModel) FindAll() ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from product order by id desc")
		if err2 != nil {
			return nil, err2
		} else {
			var products []entities.Product
			var product entities.Product
			for rows.Next() {
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Photo)
				products = append(products, product)
			}
			return products, nil
		}
	}
}

func (*ProductModel) Find(id int64) (entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return entities.Product{}, err
	} else {
		rows, err2 := db.Query("select * from product where id = ? ", id)
		if err2 != nil {
			return entities.Product{}, err2
		} else {
			var product entities.Product
			for rows.Next() {
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Photo)
			}
			return product, nil
		}
	}
}

func (*ProductModel) Create(product *entities.Product) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("insert into product (name,price,quantity,photo) values (?,?,?,?)",
			product.Name, product.Price, product.Quantity, product.Photo)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*ProductModel) Update(product entities.Product) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		var result sql.Result
		var err2 error
		if product.Photo == "" {
			result, err2 = db.Exec("update product set name = ?, price = ?, quantity = ? where id=?",
				product.Name, product.Price, product.Quantity, product.Id)
		} else {
			result, err2 = db.Exec("update product set name = ?, price = ?, quantity = ?, photo = ? where id=?",
				product.Name, product.Price, product.Quantity, product.Photo, product.Id)
		}
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*ProductModel) Delete(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("delete from product where id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}
