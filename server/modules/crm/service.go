package crm

import (
	"net/http"
	"plan-O/modules/crm/customers"
	"plan-O/modules/crm/items"
	"plan-O/modules/crm/orders"
	"plan-O/repository"

	"github.com/gorilla/mux"
)

type Service struct {
	r               *mux.Router
	itemService     *items.Service
	ordersService   *orders.Service
	supplierService *items.Service
	customerService *customers.Service
}

func NewService(db repository.CRMRepo) *Service {
	s := &Service{
		r:               mux.NewRouter(),
		itemService:     items.NewService(db),
		ordersService:   orders.NewService(db),
		supplierService: items.NewService(db),
		customerService: customers.NewService(db),
	}

	s.r.PathPrefix("/items").Handler(http.StripPrefix("/items", s.itemService))
	s.r.PathPrefix("/orders").Handler(http.StripPrefix("/orders", s.ordersService))
	s.r.PathPrefix("/suppliers").Handler(http.StripPrefix("/suppliers", s.supplierService))
	s.r.PathPrefix("/customers").Handler(http.StripPrefix("/customers", s.customerService))

	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}

func (s *Service) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, SessionID")
		next.ServeHTTP(w, r)
	})
}
