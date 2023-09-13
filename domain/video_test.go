package domain_test

import (
	"testing"
	"time"

	"github.com/flavionanni/encoder-video/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_Video_IfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func Test_Video_InvalidUuid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "12345"
	video.ResourceID = "12345"
	video.FilePath = "path/example"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

func Test_Video_NewVideoSuccessfully(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewString()
	video.ResourceID = "12345"
	video.FilePath = "path/example"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}
