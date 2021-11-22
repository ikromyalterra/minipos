package user

type RequestUser struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=4,max=16,alphanum"`
	Role       string `json:"role" validate:"oneof=admin merchant outlet"`
	MerchantID uint   `json:"id_merchant" validate:"required_if=Role merchant"`
	OutletID   uint   `json:"id_outlet" validate:"required_if=Role outlet"`
}

type RequestUserUpdate struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"omitempty,required,min=4,max=16,alphanum"`
	Role       string `json:"role" validate:"oneof=admin merchant outlet"`
	MerchantID uint   `json:"id_merchant" validate:"required_if=Role merchant"`
	OutletID   uint   `json:"id_outlet" validate:"required_if=Role outlet"`
}
