package user

type ResponseUser struct {
	ID         uint        `json:"id"`
	Email      string      `json:"email"`
	Role       string      `json:"role"`
	MerchantID interface{} `json:"id_merchant"`
	OutletID   interface{} `json:"id_outlet"`
}

type ResponseUsersView struct {
	User []*ResponseUserView `json:"user"`
}

type ResponseUserView struct {
	ID       uint                   `json:"id"`
	Email    string                 `json:"email"`
	Role     string                 `json:"role"`
	Merchant map[string]interface{} `json:"merchant"`
	Outlet   map[string]interface{} `json:"outlet"`
}
