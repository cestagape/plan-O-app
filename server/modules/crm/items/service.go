package items

import (
	"context"
	"net/http"
	"plan-O/domain"
	"plan-O/repository"

	"github.com/gorilla/mux"
)

type Service struct {
	r  *mux.Router
	db repository.ItemsRepo
}

func NewService(db repository.ItemsRepo) *Service {
	s := &Service{
		r:  mux.NewRouter(),
		db: db,
	}
	s.r.HandleFunc("/get/{id:[0-9]+}", s.getByID).Methods(http.MethodGet, http.MethodOptions)
	s.r.HandleFunc("/get", s.getAll).Methods(http.MethodGet, http.MethodOptions)
	s.r.HandleFunc("/create", s.create).Methods(http.MethodPost, http.MethodOptions)
	s.r.HandleFunc("/update", s.update).Methods(http.MethodPatch, http.MethodOptions)
	s.r.HandleFunc("/delete/{id:[0-9]+}", s.delete).Methods(http.MethodDelete, http.MethodOptions)

	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}

func (s *Service) CreateItem(ctx context.Context, item domain.Item) (int, error) {
	return s.db.CreateItem(ctx, item)
}

func (s *Service) GetItems(ctx context.Context) ([]domain.Item, error) {
	return s.db.GetItems(ctx)
}

func (s *Service) GetItemByID(ctx context.Context, id int) (domain.Item, error) {
	return s.db.GetItemByID(ctx, id)
}

func (s *Service) UpdateItem(ctx context.Context, item domain.Item) error {
	return s.db.UpdateItem(ctx, item)
}

func (s *Service) DeleteItem(ctx context.Context, id int) error {
	return s.db.DeleteItem(ctx, id)
}
