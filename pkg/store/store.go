package store

import "github.com/adhaamehab/kajo/pkg/task"

// Store is an abstraction for all different task store that KAJO can use.
// this allow us to use different storage backend for our tasks. Which gives
// the user the freedom to choose the one that fits their needs.
type Store interface {
	// Create creates a new task in the store.
	Create(task *task.Task) error

	// Update updates an existing task in the store.
	Update(task *task.Task) error

	// Delete deletes an existing task from the store.
	Delete(task *task.Task) error

	// Get gets an existing task from the store.
	Get(id string) (*task.Task, error)

	// List lists all the tasks in the store.
	List() ([]*task.Task, error)
}
