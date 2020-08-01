package types

type JobStatus int

const (
	JobStatusCreated = iota
	JobStatusActive
	JobStatusCompleted
	JobStatusFailed
	JobStatusArchived
)

type Job struct {
	FileName string
	Key string
	Created int64
	Status JobStatus
}
