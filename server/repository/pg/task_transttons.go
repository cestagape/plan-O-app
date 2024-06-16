package pg

import (
	"context"
	"time"
)

func (repo *PGRepo) MoveTask(ctx context.Context, taskID int, stateID int) error {
	var prevStateID int
	err := repo.pool.QueryRow(ctx, `
		SELECT
			prev_state_id
		FROM
			task_transitions
		WHERE
			task_id = $1
		AND
			curr_state_id = $2
		`, taskID, stateID).
		Scan(
			&prevStateID,
		)
	if err != nil {
		return err
	}

	_, err = repo.pool.Exec(ctx, `UPDATE tasks SET state_id = $2 WHERE id = $1`, taskID, stateID)
	if err != nil {
		return err
	}

	_, err = repo.pool.Exec(ctx, `
		UPDATE task_transitions
		SET
			prev_state_id = $2,
			curr_state_id = $3,
			transition_time = $4
		WHERE id = $1;
		`, prevStateID, prevStateID, time.Now().UnixNano())

	return err
}
