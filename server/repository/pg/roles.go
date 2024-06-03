package pg

import (
	"context"
	"plan-O/domain"
)

func (repo *PGRepo) GetRoleByID(ctx context.Context, id int) (domain.Role, error) {
	var role domain.Role
	err := repo.pool.QueryRow(ctx, `SELECT id, name FROM roles WHERE id = $1;`, id).Scan(&role.ID, &role.Name)
	return role, err
}

func (repo *PGRepo) GetRoles(ctx context.Context) ([]domain.Role, error) {
	var roles []domain.Role
	rows, err := repo.pool.Query(ctx, `SELECT id, name FROM roles;`)
	if err != nil {
		return roles, err
	}
	defer rows.Close()

	for rows.Next() {
		var role domain.Role
		err = rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return roles, err
		}

		roles = append(roles, role)
	}

	return roles, err
}
