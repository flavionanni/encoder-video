package domain_test

import (
	"testing"
	"time"

	"github.com/flavionanni/encoder-video/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_NewJob(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewString()
	video.ResourceID = "12345"
	video.FilePath = "path/example"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path/other_example", "Completed", video)

	require.Nil(t, err)
	require.NotNil(t, job)
}

func Test_Job_ErrorIfEmptyJob(t *testing.T) {
	job, err := domain.NewJob("path/other_example", "Completed", nil)

	require.NotNil(t, err)
	require.Nil(t, job)
}
