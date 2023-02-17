package data

import "rest-api/entities"

var Products []entities.Product

func init() {
	Products = []entities.Product{
		{ID: 1, Name: "Tivi 1", Price: 300},
		{ID: 2, Name: "Tivi 2", Price: 400},
		{ID: 3, Name: "Tivi 3", Price: 500},
		{ID: 4, Name: "Tivi 4", Price: 600},
		{ID: 5, Name: "Tivi 5", Price: 700},
	}
}
