package service

import (
	"github.com/stretchr/testify/assert"
	"prometheus_mate/pkg/model"
	"testing"
)

func TestAddStaticJob(t *testing.T) {
	jobReq := model.CreateJobReq{}
	jobReq.Job = "test-static"
	job, err := AddJob(jobReq)
	assert.Nil(t, err)
	assert.Equal(t, "test-static", job.Job)
}
