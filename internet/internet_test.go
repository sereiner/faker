package internet

import (
	"testing"
)

func TestEmail(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "1", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Email(); got != tt.want {
				t.Errorf("Email() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "1", want: "gmail.com"},
		{name: "2", want: "jjad.com"},
		{name: "3", want: "llkkf.cc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DomainName(); got != tt.want {
				t.Errorf("DomainName() = %v, want %v", got, tt.want)
			}
		})
	}
}
