package address

import (
	"testing"
)

func TestAddress_CityPrefix(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		a := &Address{}
		got := a.CityPrefix()
		t.Log(got)

	})

}

func TestAddress_CitySuffix(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			if got := a.CitySuffix(); got != tt.want {
				t.Errorf("Address.CitySuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_City(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		// TODO: Add test cases.
		{name: "a", a: &Address{}, want: ""},
		{name: "a", a: &Address{}, want: ""},
		{name: "a", a: &Address{}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			if got := a.City(); got != tt.want {
				t.Log(got)

			}
		})
	}
}

func TestAddress_Latitude(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want float32
	}{
		// TODO: Add test cases.
		{name: "1", a: &Address{}, want: 0.1},
		{name: "2", a: &Address{}, want: 0.1},
		{name: "3", a: &Address{}, want: 0.1},
		{name: "4", a: &Address{}, want: 0.1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			got := a.Latitude()
			t.Log(got)

		})
	}
}

func TestAddress_Longitude(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want float32
	}{
		// TODO: Add test cases.
		{name: "1", a: &Address{}, want: 0.1},
		{name: "2", a: &Address{}, want: 0.1},
		{name: "3", a: &Address{}, want: 0.1},
		{name: "4", a: &Address{}, want: 0.1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			got := a.Longitude()
			t.Log(got)

		})
	}
}

func TestAddress_BuildingNumber(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		// TODO: Add test cases.
		{name: "1", a: &Address{}, want: "#"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			got := a.BuildingNumber()
			t.Log(got)

		})
	}
}

func TestAddress_Postcode(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		// TODO: Add test cases.
		{name: "1", a: &Address{}, want: "123"},
		{name: "2", a: &Address{}, want: "123"},
		{name: "3", a: &Address{}, want: "123"},
		{name: "4", a: &Address{}, want: "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			got := a.Postcode()
			t.Log(got)

		})
	}
}

func TestAddress_State(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			if got := a.State(); got != tt.want {
				t.Errorf("Address.State() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_StateAbbr(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			if got := a.StateAbbr(); got != tt.want {
				t.Errorf("Address.StateAbbr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_StreetSuffix(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			if got := a.StreetSuffix(); got != tt.want {
				t.Errorf("Address.StreetSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_parse(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name    string
		a       *Address
		args    args
		wantStr string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			if gotStr := a.parse(tt.args.format); gotStr != tt.wantStr {
				t.Errorf("Address.parse() = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}

func TestAddress_LastName(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			if got := a.LastName(); got != tt.want {
				t.Errorf("Address.LastName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_StreetName(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			if got := a.StreetName(); got != tt.want {
				t.Errorf("Address.StreetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_StreetAddress(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			if got := a.StreetAddress(); got != tt.want {
				t.Errorf("Address.StreetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_Address(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		// TODO: Add test cases.
		{name: "1", a: &Address{}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			got := a.Address()
			t.Log(got)

		})
	}
}

func TestAddress_Country(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		// TODO: Add test cases.
		{name: "1", a: &Address{}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{}
			got := a.Country()
			t.Log(got)

		})
	}
}
