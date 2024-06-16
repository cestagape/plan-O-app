package pg

import (
	"context"
	"plan-O/domain"
)

func (repo *PGRepo) GetCustomerByID(ctx context.Context, id int) (domain.Customer, error) {
	var customer domain.Customer
	err := repo.pool.QueryRow(ctx, `
		SELECT
			id,
			name,
			ct.id,
			ct.name,
			addr,
			email,
			phone,
			manager_name,
			notes
		INNER JOIN customer_types ct ON ct.id = customers.type_id
		FROM customers WHERE id = $1;`,
		id,
	).
		Scan(
			&customer.ID,
			&customer.Name,
			&customer.Type.ID,
			&customer.Type.Name,
			&customer.Addr,
			&customer.Email,
			&customer.Phone,
			&customer.ManagerName,
			&customer.Notes,
		)

	return customer, err
}

func (repo *PGRepo) GetCustomers(ctx context.Context) ([]domain.Customer, error) {
	var customers []domain.Customer
	rows, err := repo.pool.Query(ctx, `
		SELECT
			id,
			name,
			ct.id,
			ct.name,
			addr,
			email,
			phone,
			manager_name,
			notes
		INNER JOIN customer_types ct ON ct.id = customers.type_id
		FROM customers;`,
	)
	if err != nil {
		return customers, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer domain.Customer
		err = rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Type.ID,
			&customer.Type.Name,
			&customer.Addr,
			&customer.Email,
			&customer.Phone,
			&customer.ManagerName,
			&customer.Notes,
		)
		if err != nil {
			return customers, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (repo *PGRepo) CreateCustomer(ctx context.Context, customer domain.Customer) (int, error) {
	var id int
	err := repo.pool.QueryRow(ctx, `
		INSERT INTO customers
			(name, type_id, addr, email, phone, manager_name, notes)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
		`,
		customer.Name,
		customer.Type.ID,
		customer.Addr,
		customer.Email,
		customer.Phone,
		customer.ManagerName,
		customer.Notes,
	).Scan(&id)

	return id, err
}

func (repo *PGRepo) UpdateCustomer(ctx context.Context, customer domain.Customer) error {
	_, err := repo.pool.Exec(ctx, `
		UPDATE customers
		SET
			name = $1,
			type_id = $2,
			addr = $3,
			email = $4,
			phone = $5,
			manager_name = $6,
			notes = $7
		WHERE id = $8
		`,
		customer.Name,
		customer.Type.ID,
		customer.Addr,
		customer.Email,
		customer.Phone,
		customer.ManagerName,
		customer.Notes,
		customer.ID,
	)
	return err
}

func (repo *PGRepo) DeleteCustomer(ctx context.Context, id int) error {
	_, err := repo.pool.Exec(ctx, `DELETE FROM customers WHERE id = $1`, id)
	return err
}
