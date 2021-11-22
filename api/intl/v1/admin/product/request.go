package product

type RequestProduct struct {
	SKU        string `json:"sku" validate:"required,alphanum"`
	Image      string `json:"image"`
	Name       string `json:"name" validate:"required"`
	MerchantID uint   `json:"id_merchant" validate:"required,number,gt=0"`
	Price      uint   `json:"price"`
}
