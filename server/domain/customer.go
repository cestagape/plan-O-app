package domain

type Customer struct {
	ID          int
	Name        string
	Type        CustomerType
	Addr        string
	Email       string
	Phone       string
	ManagerName string
	Notes       any
}

type CustomerType struct {
	ID   int
	Name string
}
