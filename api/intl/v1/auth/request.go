package auth

type RequestAuth struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=16,alphanum"`
}
