package rules

import (
	"context"

	"github.com/julioshinoda/go-rest-exercise/internal/model"
)

type Rules interface {
	Evaluate() int
}

type Provider interface {
	GetDepartment(ctx context.Context, p map[string]string) (model.Departments, error)
	GetJobItem(ctx context.Context, p map[string]string) (model.Departments, error)
	CreateJob(ctx context.Context, p model.CreateJobRequest) error
}
