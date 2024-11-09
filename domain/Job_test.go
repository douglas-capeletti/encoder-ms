package domain_test

import (
	"encoder-ms/domain"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()

	job, err := domain.NewJob("output", "Converted", video)

	require.NotNil(t, job)
	require.Nil(t, err)
}
