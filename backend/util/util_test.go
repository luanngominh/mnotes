package util

import "testing"

func TestGenerateToken(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Random",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateToken()
			if len(got) != 6 {
				t.Errorf("Token error")
			}
		})
	}
}
