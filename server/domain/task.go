package domain

type Task struct {
	ID          int
	Name        string
	Deadline    int64
	Description string
	Employee    User
	CreatedAt   int64
	State       TaskState
}

type TaskState struct {
	ID   int
	Name string
}
