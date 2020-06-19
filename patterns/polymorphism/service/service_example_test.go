package service

import (
	"reflect"
	"testing"
)

// This example trying test hit an "api web service" without actually hit the "web service"
func TestApi_GetData(t *testing.T) {
	expectedData := Data{
		"I'm from Mock data",
		4,
	}

	mockService := NewMockService()
	tests := []struct {
		name     string
		wantData Data
	}{
		{"Testing Get Data from Api", expectedData},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := mockService.GetData(); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("GetData() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
