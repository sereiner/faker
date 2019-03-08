package company

import "testing"

func TestCompany(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "1", want: "111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Company(); got != tt.want {
				t.Errorf("Company() = %v, want %v", got, tt.want)
			}
		})
	}
}
