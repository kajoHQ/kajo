package task

import "time"

type Resources struct {
	CPU    string `gorm:"type:varchar(255)"` // CPU limit (e.g., "1000m" for 1000 milliCPU units)
	Memory string `gorm:"type:varchar(255)"` // Memory limit (e.g., "500Mi" for 500 Mebibytes)
}

type TaskKind int

const (
	Job TaskKind = iota
	Cron
	Collection
)

// String method is used to print the string representation of the enum.
// This is useful when printing the value of the enum.
func (tk TaskKind) String() string {
	return [...]string{"Job", "Cron", "Collection"}[tk]
}

// Task is an abstraction for all different resources that can be created by KAJO.
type Task struct {
	// Identification
	ID               string `gorm:"primary_key"`
	Name             string `gorm:"type:varchar(255)"`
	UserFriendlyName string `gorm:"type:varchar(255)"`
	Namespace        string `gorm:"type:varchar(255)"`

	// User related metadata
	Kind           TaskKind          `gorm:"type:int"`
	APIVersion     string            `gorm:"type:varchar(255)"`
	ContainerImage string            `gorm:"type:varchar(255)"`
	Labels         map[string]string `gorm:"type:json"`

	// Kubernetes related metadata
	KLabels        map[string]string `gorm:"type:json"`
	KAnnotations   map[string]string `gorm:"type:json"`
	KStatus        string            `gorm:"type:varchar(255)"`
	KStatusMessage string            `gorm:"type:varchar(255)"`

	// Time related fields
	CreatedAt  time.Time `gorm:"type:timestamp"`
	UpdatedAt  time.Time `gorm:"type:timestamp"`
	StartedAt  time.Time `gorm:"type:timestamp"`
	FinishedAt time.Time `gorm:"type:timestamp"`

	// Configuration
	EnvVars   map[string]string `gorm:"type:json"`
	Resources Resources         `gorm:"embedded"`
	Timeout   int               `gorm:"type:int"` // Timeout in seconds

	// Status
	Status  string `gorm:"type:varchar(255)"`
	Error   string `gorm:"type:varchar(255)"`
	Retries int    `gorm:"type:int"`
}
