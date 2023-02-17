package entities

type Cart struct {
	Item        Item  `json: "item"`
	Customer_id int64 `json: "customer_id"`
}
