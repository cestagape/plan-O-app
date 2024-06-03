package pg

import (
	"context"
	"plan-O/domain"
)

func (repo *PGRepo) GetProductTypes(ctx context.Context) ([]domain.ProductType, error) {
	var productTypes []domain.ProductType
	rows, err := repo.pool.Query(ctx, `SELECT id, name FROM product_types;`)
	if err != nil {
		return productTypes, err
	}
	defer rows.Close()

	for rows.Next() {
		var productType domain.ProductType
		err = rows.Scan(&productType.ID, &productType.Name)
		if err != nil {
			return productTypes, err
		}
	}

	return productTypes, nil
}

func (repo *PGRepo) GetProductTypeByID(ctx context.Context, id int) (domain.ProductType, error) {
	var productType domain.ProductType
	err := repo.pool.QueryRow(ctx,
		`SELECT id, name FROM product_types WHERE id = $1;`,
		id,
	).Scan(&productType.ID, &productType.Name)

	return productType, err
}
