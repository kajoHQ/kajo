package store

import (
	"github.com/adhaamehab/kajo/pkg/task"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SqliteStore is a store that uses bbolt as a backend.
// it implements the Store interface.
// This is more suitable for lightweight deployments where you need a single
// instance of KAJO Server.
type SqliteStore struct {
	db *gorm.DB
}

func NewSqliteStore() (*SqliteStore, error) {
	db, err := gorm.Open(sqlite.Open("kajo.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&task.Task{})

	return &SqliteStore{
		db: db,
	}, nil
}

func (s *SqliteStore) Create(task *task.Task) error {
	return s.db.Create(task).Error
}

func (s *SqliteStore) Update(task *task.Task) error {
	return s.db.Save(task).Error
}

func (s *SqliteStore) Delete(task *task.Task) error {
	return s.db.Delete(task).Error
}

func (s *SqliteStore) Get(id string) (*task.Task, error) {
	var task task.Task
	err := s.db.First(&task, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *SqliteStore) List() ([]*task.Task, error) {
	var tasks []*task.Task
	err := s.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
