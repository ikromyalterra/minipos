package productoutlet

type ResponseProductOutlet struct {
	ID        uint `json:"id"`
	ProductID uint `json:"id_product"`
	OutletID  uint `json:"id_outlet"`
	Price     uint `json:"price"`
}

type ResponseProductOutletsView struct {
	ProductOutlet []ResponseProductOutletView `json:"product_outlet"`
}

type ResponseProductOutletView struct {
	ID      uint                   `json:"id"`
	Product map[string]interface{} `json:"product"`
	Outlet  map[string]interface{} `json:"outlet"`
	Price   uint                   `json:"price"`
}
