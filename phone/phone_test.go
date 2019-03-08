package phone

import "testing"

func TestPhoneNumber(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "1", want: "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PhoneNumber(); got != tt.want {
				t.Errorf("PhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
