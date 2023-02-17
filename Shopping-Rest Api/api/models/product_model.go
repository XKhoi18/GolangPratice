package models

import (
	"Shop-Api/config"
	"Shop-Api/entities"
	"fmt"
)

// type ProductModel struct {
// 	Db *sql.DB
// }

type ProductModel struct {
}

// func (productModel ProductModel) FindAll() ([]entities.Product, error) {
// 	rows, err := productModel.Db.Query("select * from product order by id desc")
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		var products []entities.Product
// 		for rows.Next() {
// 			var id int64
// 			var name string
// 			var price float64
// 			var quantity int64
// 			var photo string
// 			err2 := rows.Scan(&id, &name, &price, &quantity, &photo)
// 			if err2 != nil {
// 				return nil, err2
// 			} else {
// 				product := entities.Product{
// 					Id:       id,
// 					Name:     name,
// 					Price:    price,
// 					Quantity: quantity,
// 					Photo:    photo,
// 				}
// 				products = append(products, product)
// 			}
// 		}
// 		return products, nil
// 	}
// }

func (*ProductModel) FindAll() ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select id,name,price,quantity,photo from product order by id desc")
		if err2 != nil {
			return nil, err2
		} else {
			var products []entities.Product
			var product entities.Product
			for rows.Next() {
				err3 := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Photo)
				if err3 != nil {
					fmt.Println("sql error")
					return nil, err3
				} else {
					products = append(products, product)
				}
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
		rows, err2 := db.Query("select id,name,price,quantity,photo from product where id = ? ", id)
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

func (*ProductModel) Search(keyword string) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select id,name,price,quantity,photo from product where name like ? order by id desc", "%"+keyword+"%")
		if err2 != nil {
			return nil, err2
		} else {
			var products []entities.Product
			var product entities.Product
			for rows.Next() {
				err3 := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Photo)
				if err3 != nil {
					fmt.Println("sql error")
					return nil, err3
				} else {
					products = append(products, product)
				}
			}
			return products, nil
		}
	}
}

func (*ProductModel) SearchPrices(min, max float64) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select id,name,price,quantity,photo from product where price >= ? and price <= ?", min, max)
		if err2 != nil {
			return nil, err2
		} else {
			var products []entities.Product
			var product entities.Product
			for rows.Next() {
				err3 := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Photo)
				if err3 != nil {
					fmt.Println("sql error")
					return nil, err3
				} else {
					products = append(products, product)
				}
			}
			return products, nil
		}
	}
}

func (*ProductModel) Create(product *entities.Product) (err error) {
	db, err := config.GetDB()
	if err != nil {
		return err
	} else {
		result, err2 := db.Exec("insert into product (name,price,quantity,photo) values (?,?,?,?)",
			product.Name, product.Price, product.Quantity, product.Photo)
		if err2 != nil {
			return err2
		} else {
			product.Id, _ = result.LastInsertId()
			return err2
		}
	}
}

func (*ProductModel) Update(product *entities.Product) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	} else {
		//var result sql.Result
		var err2 error
		if product.Photo == "" {
			_, err2 = db.Exec("update product set name = ?, price = ?, quantity = ? where id=?",
				product.Name, product.Price, product.Quantity, product.Id)
		} else {
			_, err2 = db.Exec("update product set name = ?, price = ?, quantity = ?, photo = ? where id=?",
				product.Name, product.Price, product.Quantity, product.Photo, product.Id)
		}
		if err2 != nil {
			return err2
		} else {
			return err2
		}
	}
}

func (*ProductModel) Delete(id int64) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		result, err2 := db.Exec("delete from product where id = ?", id)
		if err2 != nil {
			return 0, err2
		} else {
			return result.RowsAffected()
		}
	}
}
