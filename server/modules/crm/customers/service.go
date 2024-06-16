package customers

import (
	"context"
	"net/http"
	"plan-O/domain"
	"plan-O/repository"

	"github.com/gorilla/mux"
)

type Service struct {
	r  *mux.Router
	db repository.CustomerRepo
}

func NewService(db repository.CustomerRepo) *Service {
	s := &Service{
		r:  mux.NewRouter(),
		db: db,
	}

	s.r.HandleFunc("/get/{id:[0-9]+}", s.getByID).Methods(http.MethodGet, http.MethodOptions)
	s.r.HandleFunc("/get", s.getAll).Methods(http.MethodGet, http.MethodOptions)
	s.r.HandleFunc("/update", s.update).Methods(http.MethodPatch, http.MethodOptions)
	s.r.HandleFunc("/create", s.create).Methods(http.MethodPost, http.MethodOptions)
	s.r.HandleFunc("/delete/{id:[0-9]+}", s.delete).Methods(http.MethodDelete, http.MethodOptions)

	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}

func (s *Service) CreateCustomer(ctx context.Context, customer domain.Customer) (int, error) {
	return s.db.CreateCustomer(ctx, customer)
}

func (s *Service) GetCustomers(ctx context.Context) ([]domain.Customer, error) {
	return s.db.GetCustomers(ctx)
}

func (s *Service) GetCustomer(ctx context.Context, id int) (domain.Customer, error) {
	return s.db.GetCustomerByID(ctx, id)
}

func (s *Service) UpdateCustomer(ctx context.Context, customer domain.Customer) error {
	return s.db.UpdateCustomer(ctx, customer)
}

func (s *Service) DeleteCustomer(ctx context.Context, id int) error {
	return s.db.DeleteCustomer(ctx, id)
}
