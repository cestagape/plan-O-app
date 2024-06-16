package pg

import (
	"context"
	"plan-O/domain"
)

func (repo *PGRepo) CreateItemOrder(ctx context.Context, itemID, orderID int) error {
	_, err := repo.pool.Exec(ctx, `
		INSERT INTO orders_items
			(item_id, order_id)
		VALUES
			($1, $2)
		`,
		itemID,
		orderID,
	)

	return err
}

func (repo *PGRepo) GetItemsByOrderID(ctx context.Context, orderID int) ([]domain.Item, error) {
	var items []domain.Item

	rows, err := repo.pool.Query(ctx, `
		SELECT
			items.id,
			items.name,
			items.price,
			units.id,
			units.name,
			product_types.id,
			product_types.name
		FROM orders_items
			INNER JOIN items
				ON items.id = orders_items.item_id
			INNER JOIN units
				ON units.id = items.unit_id
			INNER JOIN product_types
				ON product_types.id = items.product_type_id
		WHERE orders_items.order_id = $1
		`,
		orderID,
	)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item domain.Item
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Price,
			&item.Unit.ID,
			&item.Unit.Name,
			&item.ProductType.ID,
			&item.ProductType.Name,
		)
		if err != nil {
			return items, err
		}

		items = append(items, item)
	}

	return items, err
}
