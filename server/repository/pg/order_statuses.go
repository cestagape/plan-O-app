package pg

import (
	"context"
	"plan-O/domain"
)

func (repo *PGRepo) GetOrderStatuses(ctx context.Context) ([]domain.OrderStatus, error) {
	var orderStatuses []domain.OrderStatus
	rows, err := repo.pool.Query(ctx, `SELECT id, name FROM order_statuses;`)
	if err != nil {
		return orderStatuses, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderStatus domain.OrderStatus
		err = rows.Scan(&orderStatus.ID, &orderStatus.Name)
		if err != nil {
			return orderStatuses, err
		}

		orderStatuses = append(orderStatuses, orderStatus)
	}
	return orderStatuses, nil
}

func (repo *PGRepo) GetOrderStatusByID(ctx context.Context, id int) (domain.OrderStatus, error) {
	var orderStatus domain.OrderStatus
	err := repo.pool.QueryRow(ctx, `SELECT id, name FROM order_statuses WHERE id = $1;`, id).Scan(&orderStatus.ID, &orderStatus.Name)
	return orderStatus, err
}
