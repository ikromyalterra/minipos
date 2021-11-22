package productoutlet

type RequestProductOutlet struct {
	ProductID uint `json:"id_product" validate:"required,number,gt=0"`
	OutletID  uint `json:"id_outlet" validate:"required,number,gt=0"`
	Price     uint `json:"price"`
}

type RequestProductOutletUpdate struct {
	ProductID uint `json:"id_product"`
	OutletID  uint `json:"id_outlet"`
	Price     uint `json:"price"`
}
