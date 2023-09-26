package task

import "time"

// KTask is an abstraction for all different resources that can be created by KAJO.
type KTask struct {
	// Identification
	ID               string
	Name             string
	UserFriendlyName string
	Kind             string
	APIVersion       string
	ContainerImage   string

	// Metadata
	Labels        map[string]string
	Annotations   map[string]string
	Status        string
	StatusMessage string

	// Time related fields
	CreatedAt  time.Time
	UpdatedAt  time.Time
	StartedAt  time.Time
	FinishedAt time.Time

	// Configuration
	EnvVars map[string]string
	Args    []string
}
