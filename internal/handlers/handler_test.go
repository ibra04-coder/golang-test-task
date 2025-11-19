package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang-test-task/internal/handlers"
)

type MockUseCase struct {
	AddAndGetSortedFunc func(n int) ([]int, error)
}

func (m *MockUseCase) AddAndGetSorted(n int) ([]int, error) {
	return m.AddAndGetSortedFunc(n)
}

func TestNumberHandler_HandleAdd(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           string
		mockFunc       func(int) ([]int, error)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:   "Success",
			method: http.MethodPost,
			body:   `{"number": 10}`,
			mockFunc: func(n int) ([]int, error) {
				return []int{1, 5, 10}, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[1,5,10]`,
		},
		{
			name:           "Invalid Method",
			method:         http.MethodGet,
			body:           ``,
			mockFunc:       nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed",
		},
		{
			name:           "Invalid JSON",
			method:         http.MethodPost,
			body:           `{"number": "invalid"}`,
			mockFunc:       nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid JSON",
		},
		{
			name:   "Internal Error",
			method: http.MethodPost,
			body:   `{"number": 5}`,
			mockFunc: func(n int) ([]int, error) {
				return nil, errors.New("service error")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Internal Server Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUC := &MockUseCase{
				AddAndGetSortedFunc: tt.mockFunc,
			}
			handler := handlers.NewNumberHandler(mockUC)

			req := httptest.NewRequest(tt.method, "/numbers", strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			handler.HandleAdd(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			if tt.expectedStatus == http.StatusOK {
				var result []int
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Fatal(err)
				}
				expectedJSON := tt.expectedBody
				actualJSONBytes, _ := json.Marshal(result)
				if string(actualJSONBytes) != expectedJSON {
					t.Errorf("expected body %s, got %s", expectedJSON, string(actualJSONBytes))
				}
			} else {
				buf := new(bytes.Buffer)
				buf.ReadFrom(resp.Body)
				if !strings.Contains(strings.TrimSpace(buf.String()), tt.expectedBody) {
					t.Errorf("expected body to contain %q, got %q", tt.expectedBody, buf.String())
				}
			}
		})
	}
}