package domain

type Supplier struct {
	ID          int
	CompanyName string
	ManagerName string
	Email       string
	Phone       string
	ProductType ProductType
	PaymentType PaymentType
	Website     string
	Comments    string
}
