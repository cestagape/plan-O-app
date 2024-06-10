package pg

import (
	"context"
	"plan-O/domain"
)

func (repo *PGRepo) GetPaymentTypeByID(ctx context.Context, id int) (domain.PaymentType, error) {
	var paymentType domain.PaymentType
	err := repo.pool.QueryRow(ctx,
		`SELECT id, name FROM payment_types WHERE id = $1;`,
		id,
	).Scan(&paymentType.ID, &paymentType.Name)

	return paymentType, err
}

func (repo *PGRepo) GetPaymentTypes(ctx context.Context) ([]domain.PaymentType, error) {
	var paymentTypes []domain.PaymentType
	rows, err := repo.pool.Query(ctx, `SELECT id, name FROM payment_types;`)
	if err != nil {
		return paymentTypes, err
	}
	defer rows.Close()

	for rows.Next() {
		var paymentType domain.PaymentType
		err = rows.Scan(&paymentType.ID, &paymentType.Name)
		if err != nil {
			return paymentTypes, err
		}

		paymentTypes = append(paymentTypes, paymentType)
	}

	return paymentTypes, nil
}

func (repo *PGRepo) CreatePaymentType(ctx context.Context, paymentType domain.PaymentType) (int, error) {
	var id int
	err := repo.pool.QueryRow(ctx, `
		INSERT INTO payment_types
			(name)
		VALUES
			($1)
		RETURNING id
		`,
		paymentType.Name,
	).Scan(&id)

	return id, err
}
