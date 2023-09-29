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
	ID               string `gorm:"primary_key" json:"id"`
	Name             string `gorm:"type:varchar(255)" json:"name"`
	UserFriendlyName string `gorm:"type:varchar(255)" json:"userFriendlyName"`
	Namespace        string `gorm:"type:varchar(255)" json:"namespace"`

	// User related metadata
	Kind           TaskKind          `gorm:"type:int" json:"kind"`
	APIVersion     string            `gorm:"type:varchar(255)" json:"apiVersion"`
	ContainerImage string            `gorm:"type:varchar(255)" json:"containerImage"`
	Labels         map[string]string `gorm:"type:json" json:"labels"`

	// Kubernetes related metadata
	KLabels        map[string]string `gorm:"type:json" json:"kLabels"`
	KAnnotations   map[string]string `gorm:"type:json" json:"kAnnotations"`
	KStatus        string            `gorm:"type:varchar(255)" json:"kStatus"`
	KStatusMessage string            `gorm:"type:varchar(255)" json:"kStatusMessage"`

	// Time related fields
	CreatedAt  time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"type:timestamp" json:"updatedAt"`
	StartedAt  time.Time `gorm:"type:timestamp" json:"startedAt"`
	FinishedAt time.Time `gorm:"type:timestamp" json:"finishedAt"`

	// Configuration
	EnvVars   map[string]string `gorm:"type:json" json:"envVars"`
	Resources Resources         `gorm:"embedded" json:"resources"`
	Timeout   int               `gorm:"type:int" json:"timeout"` // Timeout in seconds

	// Status
	Status  string `gorm:"type:varchar(255)" json:"status"`
	Error   string `gorm:"type:varchar(255)" json:"error"`
	Retries int    `gorm:"type:int" json:"retries"`
}
