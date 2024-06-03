package suppliers

import (
	"context"
	"net/http"
	"plan-O/domain"
	"plan-O/repository"

	"github.com/gorilla/mux"
)

type Service struct {
	r  *mux.Router
	db repository.SupplierRepo
}

func NewService(db repository.SupplierRepo) *Service {
	s := &Service{
		r:  mux.NewRouter(),
		db: db,
	}

	s.r.HandleFunc("/get/{id:[0-9]+}", s.getByID).Methods(http.MethodGet, http.MethodOptions)
	s.r.HandleFunc("/get", s.getAll).Methods(http.MethodGet, http.MethodOptions)
	s.r.HandleFunc("/create", s.create).Methods(http.MethodPost, http.MethodOptions)
	s.r.HandleFunc("/update/", s.update).Methods(http.MethodPatch, http.MethodOptions)
	s.r.HandleFunc("/delete/{id:[0-9]+}", s.delete).Methods(http.MethodDelete, http.MethodOptions)

	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}

func (s *Service) GetSuppliers(ctx context.Context) ([]domain.Supplier, error) {
	return s.db.GetSuppliers(ctx)
}

func (s *Service) GetSupplierByID(ctx context.Context, id int) (domain.Supplier, error) {
	return s.db.GetSupplierByID(ctx, id)
}

func (s *Service) CreateSupplier(ctx context.Context, supplier domain.Supplier) (int, error) {
	return s.db.CreateSupplier(ctx, supplier)
}

func (s *Service) UpdateSupplier(ctx context.Context, supplier domain.Supplier) error {
	return s.db.UpdateSupplier(ctx, supplier)
}

func (s *Service) DeleteSupplier(ctx context.Context, id int) error {
	return s.db.DeleteSupplier(ctx, id)
}
