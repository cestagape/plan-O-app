package pg

import (
	"context"
	"plan-O/domain"
)

func (repo *PGRepo) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	err := repo.pool.QueryRow(ctx, `
		SELECT
			id,
			name,
			r.id,
			r.name,
			login,
			password
		FROM users
			INNER JOIN roles r ON r.id = users.role_id
		WHERE id = $1;`, id,
	).
		Scan(
			&user.ID,
			&user.Name,
			&user.Role.ID,
			&user.Role.Name,
			&user.Login,
			&user.Password,
		)
	return user, err
}

func (repo *PGRepo) GetUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	rows, err := repo.pool.Query(ctx, `
		SELECT
			id,
			name,
			r.id,
			r.name,
			login,
			password
		FROM users
			INNER JOIN roles r ON r.id = users.role_id
	`)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Role.ID,
			&user.Role.Name,
			&user.Login,
			&user.Password,
		)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (repo *PGRepo) CreateUser(ctx context.Context, user domain.User) (int, error) {
	var id int
	err := repo.pool.QueryRow(ctx, `
		INSERT INTO users
			(name, role_id, login, password)
		VALUES
			($1, $2, $3, $4)
		RETURNING id
		`,
		user.Name,
		user.Role.ID,
		user.Login,
		user.Password,
	).Scan(&id)

	return id, err
}

func (repo *PGRepo) UpdateUser(ctx context.Context, user domain.User) error {
	_, err := repo.pool.Exec(ctx, `
		UPDATE users
		SET
			name = $1,
			role_id = $2,
			login = $3,
			password = $4
		WHERE id = $5
		`,
		user.Name,
		user.Role.ID,
		user.Login,
		user.Password,
		user.ID,
	)
	return err
}

func (repo *PGRepo) DeleteUser(ctx context.Context, id int) error {
	_, err := repo.pool.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	return err
}
