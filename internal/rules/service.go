package rules

import (
	"context"
	"errors"
	"log/slog"

	"github.com/julioshinoda/go-rest-exercise/internal/model"
)

type Service struct {
	OptiiClient Provider
}

func NewService(o Provider) *Service {
	return &Service{
		OptiiClient: o,
	}
}

func (s *Service) Evaluate(req model.Request) ([]string, error) {
	if req.Department != "" {
		dep, err := s.OptiiClient.GetDepartment(context.Background(), map[string]string{"displayName": req.Department})
		if err != nil {
			slog.Error(err.Error())
			return []string{}, err
		}
		if len(dep.List) == 0 {
			slog.Warn("department not found")
			return []string{}, errors.New("Bad request - department not found")
		}

	}
	if req.JobItem != "" {
		jl, err := s.OptiiClient.GetJobItem(context.Background(), map[string]string{"displayName": req.JobItem})
		if err != nil {
			slog.Error(err.Error())
			return []string{}, err
		}
		if len(jl.List) == 0 {
			slog.Warn("job item not found")
			return []string{}, errors.New("Bad request - job item not found")
		}
	}
	if IsHousekeeping(req) {
		s.OptiiClient.CreateJob(context.Background(), model.CreateJobRequest{})
	}
	return []string{}, nil
}

func IsHousekeeping(req model.Request) bool {
	if req.Department != "Housekeeping" {
		return false
	}
	for _, v := range req.Locations {
		if v == "Blanket" {
			return true
		}
		if v == "Sheets" {
			return true
		}
		if v == "Mattress" {
			return true
		}
	}
	return false
}
