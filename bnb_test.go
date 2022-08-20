package go_bnb_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCurrentCurrency(t *testing.T) {
	tests := []struct {
		name             string
		expectedResponse string
		want             string
		wantErr          bool
	}{
		{
			name:    "ok",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCurrentCurrency(451)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
			}
		})
	}
}
