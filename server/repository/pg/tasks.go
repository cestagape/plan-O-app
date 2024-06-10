package pg

import (
	"context"
	"plan-O/domain"
)

func (repo *PGRepo) GetTaskByID(ctx context.Context, id int) (domain.Task, error) {
	var task domain.Task

	err := repo.pool.QueryRow(ctx, `
		SELECT
			id,
			name,
			deadline,
			description,
			employee_id,
			created_at,
			states.id,
			states.name
		FROM tasks
		INNER JOIN task_states states ON states.id = tasks.state_id
		WHERE id = $1;`,
		id,
	).
		Scan(
			&task.ID,
			&task.Name,
			&task.Deadline,
			&task.Description,
			&task.Employee.ID,
			&task.CreatedAt,
			&task.State.ID,
			&task.State.Name,
		)
	if err != nil {
		return domain.Task{}, err
	}

	u, err := repo.GetUserByID(ctx, task.Employee.ID)
	if err != nil {
		return domain.Task{}, err
	}

	task.Employee = u
	return task, err
}

func (repo *PGRepo) GetTasks(ctx context.Context) ([]domain.Task, error) {
	var tasks []domain.Task
	rows, err := repo.pool.Query(ctx, `
		SELECT
			id,
			name,
			deadline,
			description,
			employee_id,
			created_at,
			states.id,
			states.name
		FROM tasks
		INNER JOIN task_states states ON states.id = tasks.state_id
		`)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		var task domain.Task
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Deadline,
			&task.Description,
			&task.Employee.ID,
			&task.CreatedAt,
			&task.State.ID,
			&task.State.Name,
		)
		if err != nil {
			return tasks, err
		}

		u, err := repo.GetUserByID(ctx, task.Employee.ID)
		if err != nil {
			return tasks, err
		}

		task.Employee = u
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repo *PGRepo) CreateTask(ctx context.Context, task domain.Task) (int, error) {
	var id int
	err := repo.pool.QueryRow(ctx, `
		INSERT INTO tasks
			(name, deadline, description, employee_id, created_at, state_id)
		VALUES
			($1, $2, $3, $4, $5, $6)
		RETURNING id
		`,
		task.Name,
		task.Deadline,
		task.Description,
		task.Employee.ID,
		task.CreatedAt,
		task.State.ID,
	).Scan(&id)

	return id, err
}

func (repo *PGRepo) UpdateTask(ctx context.Context, task domain.Task) error {
	_, err := repo.pool.Exec(ctx, `
		UPDATE tasks
		SET
			name = $1,
			deadline = $2,
			description = $3,
			employee_id = $4,
			state_id = $5
		WHERE id = $6
		`,
		task.Name,
		task.Deadline,
		task.Description,
		task.Employee.ID,
		task.State.ID,
		task.ID,
	)
	return err
}

func (repo *PGRepo) DeleteTask(ctx context.Context, id int) error {
	_, err := repo.pool.Exec(ctx, `DELETE FROM tasks WHERE id = $1`, id)
	return err
}

func (repo *PGRepo) GetTasksByEmployeeID(ctx context.Context, id int) ([]domain.Task, error) {
	var tasks []domain.Task
	rows, err := repo.pool.Query(ctx, `
		SELECT
			id,
			name,
			deadline,
			description,
			employee_id,
			created_at,
			states.id,
			states.name
		FROM tasks
		INNER JOIN task_states states ON states.id = tasks.state_id
		WHERE employee_id = $1
		`,
		id,
	)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		var task domain.Task
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Deadline,
			&task.Description,
			&task.Employee.ID,
			&task.CreatedAt,
			&task.State.ID,
			&task.State.Name,
		)
		if err != nil {
			return tasks, err
		}

		u, err := repo.GetUserByID(ctx, task.Employee.ID)
		if err != nil {
			return tasks, err
		}

		task.Employee = u
		tasks = append(tasks, task)
	}

	return tasks, nil
}
