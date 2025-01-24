package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(User) error
	GetUserByID(id int) (*User, error)
}

type ProductStore interface {
	GetProducts() ([]Product, error)
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=150"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Product struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}
