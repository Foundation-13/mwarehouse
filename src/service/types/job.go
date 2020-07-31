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

type JobStatusDTO struct {
	Key		string		`json:"key"`
	Status	JobStatus	`json:"status"`
}

func NewJobStatusDTO(job Job) JobStatusDTO {
	return JobStatusDTO{
		Key: job.Key,
		Status: job.Status,
	}
}