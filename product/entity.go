package product

import "time"

type Product struct {
	ID          int
	Title       string
	Price       int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
