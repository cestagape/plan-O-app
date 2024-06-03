package pg

import (
	"context"
	"plan-O/domain"
)

func (repo *PGRepo) GetOrders(ctx context.Context) ([]domain.Order, error) {
	var orders []domain.Order
	rows, err := repo.pool.Query(ctx, `
		SELECT
			id,
			created_at,
			deadline,
			st.id,
			st.name,
			design,
			quantnity,
			price,
			addr,
			employee_id,
			deadline,
			pt.id,
			pt.name,
			delivery_info,
			additional
		INNER JOIN order_statuses st ON st.id = orders.status_id
		INNER JOIN payment_types pt ON pt.id = orders.payment_type_id
		FROM orders
	`)
	if err != nil {
		return orders, err
	}
	defer rows.Close()

	for rows.Next() {
		var order domain.Order
		err = rows.Scan(
			&order.ID,
			&order.CreatedAt,
			&order.Deadline,
			&order.Status.ID,
			&order.Status.Name,
			&order.Design,
			&order.Quantnity,
			&order.Price,
			&order.Addr,
			&order.Employee.ID,
			&order.Deadline,
			&order.PaymentType.ID,
			&order.PaymentType.Name,
			&order.DeliveryInfo,
			&order.Additional,
		)
		if err != nil {
			return orders, err
		}

		orders = append(orders, order)
	}

	for i := range orders {
		items, err := repo.GetItemsByOrderID(ctx, orders[i].ID)
		if err != nil {
			return orders, err
		}

		user, err := repo.GetUserByID(ctx, orders[i].Employee.ID)
		if err != nil {
			return orders, err
		}
		orders[i].Items = items
		orders[i].Employee = user
	}

	return orders, err
}

func (repo *PGRepo) GetOrderByID(ctx context.Context, id int) (domain.Order, error) {
	var order domain.Order
	err := repo.pool.QueryRow(ctx, `
		SELECT
			id,
			created_at,
			deadline,
			st.id,
			st.name,
			design,
			quantnity,
			price,
			addr,
			employee_id,
			deadline,
			pt.id,
			pt.name,
			delivery_info,
			additional
		FROM orders
		INNER JOIN order_statuses st ON st.id = orders.status_id
		INNER JOIN payment_types pt ON pt.id = orders.payment_type_id
		WHERE id = $1
		`,
		id,
	).
		Scan(
			&order.ID,
			&order.CreatedAt,
			&order.Deadline,
			&order.Status.ID,
			&order.Status.Name,
			&order.Design,
			&order.Quantnity,
			&order.Price,
			&order.Addr,
			&order.Employee.ID,
			&order.Deadline,
			&order.PaymentType.ID,
			&order.PaymentType.Name,
			&order.DeliveryInfo,
			&order.Additional,
		)
	return order, err
}

func (repo *PGRepo) CreateOrder(ctx context.Context, order domain.Order) (int, error) {
	var id int
	err := repo.pool.QueryRow(ctx, `
		INSERT INTO orders (
			created_at,
			design,
			quantnity,
			status_id,
			price,
			addr,
			employee_id,
			deadline,
			payment_type_id,
			delivery_info,
			additional
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11
		) RETURNING id;
	`,
		&order.CreatedAt,
		&order.Design,
		&order.Quantnity,
		&order.Status.ID,
		&order.Price,
		&order.Addr,
		&order.Employee.ID,
		&order.Deadline,
		&order.PaymentType.ID,
		&order.DeliveryInfo,
		&order.Additional,
	).Scan(&id)

	return id, err
}

func (repo *PGRepo) UpdateOrder(ctx context.Context, order domain.Order) error {
	_, err := repo.pool.Exec(ctx, `
		UPDATE orders 
		SET
			created_at = $1,
			design = $2,
			quantnity = $3,
			status_id = $4,
			price = $5,
			addr = $6,
			employee_id = $7,
			deadline = $8,
			payment_type_id = $9,
			delivery_info = $10,
			additional =  $11
		WHERE id  =  $12
			`,
		&order.CreatedAt,
		&order.Design,
		&order.Quantnity,
		&order.Status.ID,
		&order.Price,
		&order.Addr,
		&order.Employee.ID,
		&order.Deadline,
		&order.PaymentType.ID,
		&order.DeliveryInfo,
		&order.Additional,
		&order.ID,
	)
	return err
}

func (repo *PGRepo) DeleteOrder(ctx context.Context, id int) error {
	_, err := repo.pool.Exec(ctx, `
		DELETE FROM orders
		WHERE id = $1
	`,
		id,
	)
	return err

}
