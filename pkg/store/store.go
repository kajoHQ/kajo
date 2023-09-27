package store

import (
	. "github.com/adhaamehab/kajo/pkg/task"
)

// Store is an abstraction for all different task store that KAJO can use.
// this allow us to use different storage backend for our tasks. Which gives
// the user the freedom to choose the one that fits their needs.
type Store interface {
	// Create creates a new task in the store.
	Create(task *Task) error

	// Update updates an existing task in the store.
	Update(task *Task) error

	// Delete deletes an existing task from the store.
	Delete(task *Task) error

	// Get gets an existing task from the store.
	Get(id string) (*Task, error)

	// List lists all the tasks in the store.
	List() ([]*Task, error)

	// ListByStatus lists all the tasks in the store with the given status.
	ListByStatus(status string) ([]*Task, error)

	// ListByKind lists all the tasks in the store with the given kind.
	ListByKind(kind string) ([]*Task, error)

	// ListByKindAndStatus lists all the tasks in the store with the given kind and status.
	ListByKindAndStatus(kind, status string) ([]*Task, error)

	// ListByLabel lists all the tasks in the store with the given label.
	ListByLabel(label string) ([]*Task, error)

	// GetSummary gets the summary of the tasks in the store.
	GetSummary() (map[string]int, error)
}
