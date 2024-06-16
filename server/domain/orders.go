package domain

type Order struct {
	ID           int
	CreatedAt    int64
	Design       string
	Quantnity    int
	Status       OrderStatus
	Price        int
	Addr         string
	Employee     User
	Deadline     int64
	PaymentType  PaymentType
	DeliveryInfo any
	Additional   any
	Items        []Item
}
