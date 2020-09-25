package entities


type User struct {
	Id          int64  `json:"id"`
	UserName   string `json:"user_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
	CreatedAt string `json:"created_at"`
}

type NewUser struct {
	UserName   string `json:"user_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
	CreatedAt string `json:"created_at"`
}

type Product struct {
	ProductID int64 `json:"product_id"`
	ProductName string `json:"product_name"`
	ProductType string `json:"type"`
	Availability bool `json:"availability"`
}

type NewProduct struct {
	ProductName string `json:"product_name"`
	ProductType string `json:"type"`
	Availability bool `json:"availability"`
}