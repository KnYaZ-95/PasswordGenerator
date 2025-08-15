package passwordGenerator

import (
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name           string
		length         int
		special        bool
		expectedLength int
		error          string
	}{
		{"Length 10, no special", 10, false, 10, ""},
		{"Length 5, with special", 5, true, 5, ""},
		{"Length 0", 0, true, 0, "password length must be greater than zero"},
		{"Length -1", -1, true, -1, "password length must be greater than zero"}, // Должен вернуть ошибку
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GeneratePassword(tt.length, tt.special)

			if err != nil && err.Error() != tt.error {
				t.Errorf("Got error = %v, expected error = %v", err.Error(), tt.error)
				return
			}

			if tt.error == "" {
				if len(got) != tt.expectedLength {
					t.Errorf("Длина: got %d, want %d", len(got), tt.expectedLength)
				}
			}
		})
	}
}
