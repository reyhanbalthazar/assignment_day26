package main

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}

type Inventory struct {
	ID        uint   `gorm:"primaryKey"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Location  string `json:"location"`
}

type Order struct {
	ID        uint   `gorm:"primaryKey"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	OrderDate string `json:"order_date"`
}
