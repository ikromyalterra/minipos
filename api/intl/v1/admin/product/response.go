package product

type ResponseProduct struct {
	ID         uint   `json:"id"`
	SKU        string `json:"sku"`
	Image      string `json:"image"`
	Name       string `json:"name"`
	Price      uint   `json:"price"`
	MerchantID uint   `json:"id_merchant"`
}

type ResponseProductsView struct {
	Product []ResponseProductView `json:"product"`
}

type ResponseProductView struct {
	ID       uint                   `json:"id"`
	SKU      string                 `json:"sku"`
	Image    string                 `json:"image"`
	Name     string                 `json:"name"`
	Price    uint                   `json:"price"`
	Merchant map[string]interface{} `json:"merchant"`
}
