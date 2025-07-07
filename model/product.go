package model

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Quantity string `json:"qty"`
}

type ResponseProduct struct {
	Result ResultProduct `json:"result,omitempty"`
	Error  string        `json:"error,omitempty"`
}

type ResultProduct struct {
	Sales    []SalesTransaction `json:"sales,omitempty"`
	Products []Product          `json:"products,omitempty"`
}

type SalesTransaction struct {
	BuyerId   string `json:"buyer_id"`
	ProductId string `json:"product_id"`
	Quantity  string `json:"qty"`
}
