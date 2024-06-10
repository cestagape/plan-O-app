package pg

import (
	"context"
	"plan-O/domain"
)

func (repo *PGRepo) GetItems(ctx context.Context) ([]domain.Item, error) {
	rows, err := repo.pool.Query(ctx, `
		SELECT
			id,
			name,
			price,
			units.id,
			units.name
		FROM items
			INNER JOIN units
				ON units.id = items.unit_id;`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.Item
	for rows.Next() {
		var item domain.Item
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Price,
			&item.Unit.ID,
			&item.Unit.Name,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (repo *PGRepo) GetItemByID(ctx context.Context, id int) (domain.Item, error) {
	var item domain.Item
	err := repo.pool.QueryRow(ctx, `
		SELECT
			id,
			name,
			price,
			units.id,
			units.name,
		FROM items
			INNER JOIN units
				ON units.id = items.unit_id
			INNER JOIN product_types
				ON product_types.id = items.product_type_id
		WHERE id = $1;
		`,
		id,
	).Scan(
		&item.ID,
		&item.Name,
		&item.Price,
		&item.Unit.ID,
		&item.Unit.Name,
	)

	return item, err
}

func (repo *PGRepo) CreateItem(ctx context.Context, item domain.Item) (int, error) {
	var id int
	err := repo.pool.QueryRow(ctx, `
		INSERT INTO items
			(name, price, unit_id)
		VALUES
			($1, $2, $3)
		RETURNING id
		`,
		item.Name,
		item.Price,
		item.Unit.ID,
	).Scan(&id)

	return id, err
}

func (repo *PGRepo) UpdateItem(ctx context.Context, item domain.Item) error {
	_, err := repo.pool.Exec(ctx, `
		UPDATE items
		SET
			name = $1,
			price = $2,
			unit_id = $3
		WHERE id = $4
		`,
		item.Name,
		item.Price,
		item.Unit.ID,
		item.ID,
	)
	return err
}

func (repo *PGRepo) DeleteItem(ctx context.Context, id int) error {
	_, err := repo.pool.Exec(ctx, `DELETE FROM items WHERE id = $1`, id)
	return err
}
