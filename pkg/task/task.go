package task

import "time"

type Resources struct {
	CPU    string // CPU limit (e.g., "1000m" for 1000 milliCPU units)
	Memory string // Memory limit (e.g., "500Mi" for 500 Mebibytes)
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
	ID               string
	Name             string
	UserFriendlyName string

	// User related metadata
	Kind           TaskKind
	APIVersion     string
	ContainerImage string
	Labels         map[string]string

	// Kubernetes related metadata
	KLabels        map[string]string
	KAnnotations   map[string]string
	KStatus        string
	KStatusMessage string

	// Time related fields
	CreatedAt  time.Time
	UpdatedAt  time.Time
	StartedAt  time.Time
	FinishedAt time.Time

	// Configuration
	EnvVars   map[string]string
	Args      []string
	Resources Resources
	Timeout   time.Duration

	// Status
	Status  string
	Error   string
	Retries int
}
