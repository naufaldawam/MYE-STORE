package delivery

type InsertFormat struct {
	IdProduct int `json:"product_id" form:"product_id"`
	Stock     int `json:"stock" form:"stock"`
}
