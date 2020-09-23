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
}

type Product struct {
	Id int64 `json:"id"`
	Type string `json:"type"`
	Availability bool `json:"availability"`
}

type NewProduct struct {
	Type string `json:"type"`
	Availability bool `json:"availability"`
}