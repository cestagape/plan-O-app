package pg

import (
	"context"
	"plan-O/domain"
)

func (repo *PGRepo) GetSupplierByID(ctx context.Context, id int) (domain.Supplier, error) {
	var supplier domain.Supplier
	err := repo.pool.QueryRow(ctx, `
		SELECT
			id,
			company_name,
			manager_name,
			email,
			phone,
			prt.id,
			prt.name,
			pt.id,
			pt.name,
			website_link,
			comments
		FROM suppliers
		INNER JOIN product_types prt ON prt.id = suppliers.product_type_id
		INNER JOIN payment_types pt ON pt.id = suppliers.payment_type_id
		WHERE id = $1;`,
		id,
	).
		Scan(
			&supplier.ID,
			&supplier.CompanyName,
			&supplier.ManagerName,
			&supplier.Email,
			&supplier.Phone,
			&supplier.ProductType.ID,
			&supplier.ProductType.Name,
			&supplier.PaymentType.ID,
			&supplier.PaymentType.Name,
			&supplier.Website,
			&supplier.Comments,
		)

	return supplier, err
}

func (repo *PGRepo) GetSuppliers(ctx context.Context) ([]domain.Supplier, error) {
	var suppliers []domain.Supplier
	rows, err := repo.pool.Query(ctx, `
		SELECT
			id,
			company_name,
			manager_name,
			email,
			phone,
			prt.id,
			prt.name,
			pt.id,
			pt.name,
			website_link,
			comments
		FROM suppliers
		INNER JOIN product_types prt ON prt.id = suppliers.product_type_id
		INNER JOIN payment_types pt ON pt.id = suppliers.payment_type_id;`,
	)
	if err != nil {
		return suppliers, err
	}
	defer rows.Close()

	for rows.Next() {
		var supplier domain.Supplier
		err := rows.Scan(
			&supplier.ID,
			&supplier.CompanyName,
			&supplier.ManagerName,
			&supplier.Email,
			&supplier.Phone,
			&supplier.ProductType.ID,
			&supplier.ProductType.Name,
			&supplier.PaymentType.ID,
			&supplier.PaymentType.Name,
			&supplier.Website,
			&supplier.Comments,
		)
		if err != nil {
			return suppliers, err
		}

		suppliers = append(suppliers, supplier)
	}

	return suppliers, err
}

func (repo *PGRepo) CreateSupplier(ctx context.Context, supplier domain.Supplier) (int, error) {
	var id int
	err := repo.pool.QueryRow(ctx, `
		INSERT INTO suppliers (
			company_name,
			manager_name,
			email,
			phone,
			product_type_id,
			payment_type_id,
			website_link,
			comments
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;`,
		supplier.CompanyName,
		supplier.ManagerName,
		supplier.Email,
		supplier.Phone,
		supplier.ProductType.ID,
		supplier.PaymentType.ID,
		supplier.Website,
		supplier.Comments,
	).Scan(&id)

	return id, err
}

func (repo *PGRepo) UpdateSupplier(ctx context.Context, supplier domain.Supplier) error {
	_, err := repo.pool.Exec(ctx, `
		UPDATE suppliers
		SET
			company_name = $1,
			manager_name = $2,
			email = $3,
			phone = $4,
			product_type_id = $5,
			payment_type_id = $6,
			website_link = $7,
			comments = $8
		WHERE id = $9
		RETURNING id;`,
		supplier.CompanyName,
		supplier.ManagerName,
		supplier.Email,
		supplier.Phone,
		supplier.ProductType.ID,
		supplier.PaymentType.ID,
		supplier.Website,
		supplier.Comments,
		supplier.ID,
	)

	return err
}

func (repo *PGRepo) DeleteSupplier(ctx context.Context, id int) error {
	_, err := repo.pool.Exec(ctx, `
		DELETE FROM suppliers
		WHERE id = $1;`,
		id,
	)

	return err
}
