package usecase_test

import (
	"errors"
	"reflect"
	"testing"

	"golang-test-task/internal/usecase"
)

type MockRepository struct {
	SaveFunc         func(n int) error
	GetAllSortedFunc func() ([]int, error)
}

func (m *MockRepository) Save(n int) error {
	return m.SaveFunc(n)
}

func (m *MockRepository) GetAllSorted() ([]int, error) {
	return m.GetAllSortedFunc()
}

func TestService_AddAndGetSorted(t *testing.T) {
	tests := []struct {
		name          string
		input         int
		mockSave      func(int) error
		mockGetAll    func() ([]int, error)
		expected      []int
		expectedError bool
	}{
		{
			name:  "Success",
			input: 5,
			mockSave: func(n int) error {
				if n != 5 {
					return errors.New("wrong number saved")
				}
				return nil
			},
			mockGetAll: func() ([]int, error) {
				return []int{1, 2, 5}, nil
			},
			expected:      []int{1, 2, 5},
			expectedError: false,
		},
		{
			name:  "Save Error",
			input: 10,
			mockSave: func(n int) error {
				return errors.New("db error")
			},
			mockGetAll: func() ([]int, error) {
				return nil, nil
			},
			expected:      nil,
			expectedError: true,
		},
		{
			name:  "Get Error",
			input: 3,
			mockSave: func(n int) error {
				return nil
			},
			mockGetAll: func() ([]int, error) {
				return nil, errors.New("read error")
			},
			expected:      nil,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockRepository{
				SaveFunc:         tt.mockSave,
				GetAllSortedFunc: tt.mockGetAll,
			}
			service := usecase.NewService(mockRepo)

			result, err := service.AddAndGetSorted(tt.input)

			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}

			if !tt.expectedError && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected result: %v, got: %v", tt.expected, result)
			}
		})
	}
}