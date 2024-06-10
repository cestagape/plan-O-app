package orders

import (
	"context"
	"net/http"
	"plan-O/domain"
	"plan-O/repository"

	"github.com/gorilla/mux"
)

type Service struct {
	r  *mux.Router
	db repository.OrdersRepo
}

func NewService(db repository.OrdersRepo) *Service {
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

func (s *Service) CreateOrder(ctx context.Context, order domain.Order) (int, error) {
	return s.db.CreateOrder(ctx, order)
}

func (s *Service) GetOrders(ctx context.Context) ([]domain.Order, error) {
	return s.db.GetOrders(ctx)
}

func (s *Service) GetOrder(ctx context.Context, id int) (domain.Order, error) {
	return s.db.GetOrderByID(ctx, id)
}

func (s *Service) UpdateOrder(ctx context.Context, order domain.Order) error {
	return s.db.UpdateOrder(ctx, order)
}

func (s *Service) DeleteOrder(ctx context.Context, id int) error {
	return s.db.DeleteOrder(ctx, id)
}
