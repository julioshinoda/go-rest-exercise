package rules_test

import (
	"errors"
	"testing"

	"github.com/julioshinoda/go-rest-exercise/internal/model"
	"github.com/julioshinoda/go-rest-exercise/internal/rules"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestService_Evaluate_Not_Found_department(t *testing.T) {
	type fields struct {
		OptiiClient rules.Provider
	}
	type args struct {
		req model.Request
	}
	mProvider := &rules.MockProvider{}
	mProvider.On("GetDepartment", mock.Anything, mock.Anything).Return(model.Departments{}, errors.New("Bad request"))
	s := &rules.Service{
		OptiiClient: mProvider,
	}
	_, err := s.Evaluate(model.Request{
		Department: "not-exist",
	})
	require.Error(t, err)

}

func TestService_Evaluate_not_found_job_item(t *testing.T) {
	type fields struct {
		OptiiClient rules.Provider
	}
	type args struct {
		req model.Request
	}
	mProvider := &rules.MockProvider{}
	mProvider.On("GetDepartment", mock.Anything, mock.Anything).Return(model.Departments{List: []model.Department{{ID: 1}}}, nil)
	mProvider.On("GetJobItem", mock.Anything, mock.Anything).Return(model.Departments{}, errors.New("errror"))
	s := &rules.Service{
		OptiiClient: mProvider,
	}
	_, err := s.Evaluate(model.Request{
		Department: "Enginnering",
		JobItem:    "error",
	})
	require.Error(t, err)

}
