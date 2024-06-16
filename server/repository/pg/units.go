package pg

import (
	"context"
	"plan-O/domain"
)

func (repo *PGRepo) GetUnits(ctx context.Context) ([]domain.Unit, error) {
	var units []domain.Unit
	rows, err := repo.pool.Query(ctx, `SELECT id, name FROM units;`)
	if err != nil {
		return units, err
	}
	defer rows.Close()

	for rows.Next() {
		var unit domain.Unit
		err = rows.Scan(&unit.ID, &unit.Name)
		if err != nil {
			return units, err
		}

		units = append(units, unit)
	}

	return units, err
}

func (repo *PGRepo) GetUnitByID(ctx context.Context, id int) (domain.Unit, error) {
	var unit domain.Unit
	err := repo.pool.QueryRow(ctx, `SELECT id, name FROM units WHERE id = $1;`, id).Scan(&unit.ID, &unit.Name)
	return unit, err
}
