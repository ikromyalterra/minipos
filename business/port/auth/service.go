package auth

type (
	// Product data
	AuthUser struct {
		UserID   uint
		Email    string
		Role     string
		Password string
		Token    string
	}

	// Service is inbount port
	Service interface {
		// Verify token
		Verify(tokenString string) (interface{}, error)

		// Login user
		Login(authUser *AuthUser) error

		// Logout user
		Logout(tokenID uint) error
	}
)
