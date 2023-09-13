package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Job struct {
	ID               string    `valid:"uuid"`
	OutputBucketPath string    `valid:"required"`
	Status           string    `valid:"required"`
	Video            *Video    `valid:"required"`
	Error            string    `valid:"-"`
	CreatedAt        time.Time `valid:"-"`
	UpdatedAt        time.Time `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewJob(path string, status string, video *Video) (*Job, error) {
	job := &Job{
		ID:               uuid.NewString(),
		OutputBucketPath: path,
		Status:           status,
		Video:            video,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err := job.Validate()
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (j *Job) Validate() error {
	_, err := govalidator.ValidateStruct(j)

	return err
}
