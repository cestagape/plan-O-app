package repository

import (
	"context"
	"plan-O/domain"
)

type Repo interface {
	UserRepo
	UnitRepo
	TaskRepo
	SupplierRepo
	RoleRepo
	ProductTypeRepo
	PaymentTypeRepo
	OrdersRepo
	ItemsRepo
	CustomerRepo
}

type CRMRepo interface {
	ItemsRepo
	SupplierRepo
	OrdersRepo
	CustomerRepo
}

type UserRepo interface {
	GetUserByID(ctx context.Context, id int) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (int, error)
	UpdateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, id int) error
}

type UnitRepo interface {
	GetUnitByID(ctx context.Context, id int) (domain.Unit, error)
	GetUnits(ctx context.Context) ([]domain.Unit, error)
}

type TaskRepo interface {
	GetTaskByID(ctx context.Context, id int) (domain.Task, error)
	GetTasks(ctx context.Context) ([]domain.Task, error)
	CreateTask(ctx context.Context, task domain.Task) (int, error)
	UpdateTask(ctx context.Context, task domain.Task) error
	DeleteTask(ctx context.Context, id int) error
	GetTasksByEmployeeID(ctx context.Context, id int) ([]domain.Task, error)
}

type SupplierRepo interface {
	GetSupplierByID(ctx context.Context, id int) (domain.Supplier, error)
	GetSuppliers(ctx context.Context) ([]domain.Supplier, error)
	CreateSupplier(ctx context.Context, supplier domain.Supplier) (int, error)
	UpdateSupplier(ctx context.Context, supplier domain.Supplier) error
	DeleteSupplier(ctx context.Context, id int) error
}

type RoleRepo interface {
	GetRoleByID(ctx context.Context, id int) (domain.Role, error)
	GetRoles(ctx context.Context) ([]domain.Role, error)
}

type ProductTypeRepo interface {
	GetProductTypeByID(ctx context.Context, id int) (domain.ProductType, error)
	GetProductTypes(ctx context.Context) ([]domain.ProductType, error)
}

type PaymentTypeRepo interface {
	GetPaymentTypeByID(ctx context.Context, id int) (domain.PaymentType, error)
	GetPaymentTypes(ctx context.Context) ([]domain.PaymentType, error)
}

type OrdersRepo interface {
	GetOrderByID(ctx context.Context, id int) (domain.Order, error)
	GetOrders(ctx context.Context) ([]domain.Order, error)
	CreateOrder(ctx context.Context, order domain.Order) (int, error)
	UpdateOrder(ctx context.Context, order domain.Order) error
	DeleteOrder(ctx context.Context, id int) error
}

type ItemsRepo interface {
	GetItemByID(ctx context.Context, id int) (domain.Item, error)
	GetItems(ctx context.Context) ([]domain.Item, error)
	CreateItem(ctx context.Context, item domain.Item) (int, error)
	UpdateItem(ctx context.Context, item domain.Item) error
	DeleteItem(ctx context.Context, id int) error
}

type CustomerRepo interface {
	GetCustomerByID(ctx context.Context, id int) (domain.Customer, error)
	GetCustomers(ctx context.Context) ([]domain.Customer, error)
	CreateCustomer(ctx context.Context, customer domain.Customer) (int, error)
	UpdateCustomer(ctx context.Context, customer domain.Customer) error
	DeleteCustomer(ctx context.Context, id int) error
}
