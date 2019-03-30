package base

import (
	"fmt"
	"testing"
)

func TestRandomDigit(t *testing.T) {
	fmt.Println(RandomDigit())
}

func TestNumberBetween(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{min: 0, max: 10}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberBetween(tt.args.min, tt.args.max); got > 10 {
				t.Errorf("NumberBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomDigitNot(t *testing.T) {
	type args struct {
		except int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomDigitNot(tt.args.except); got != tt.want {
				t.Errorf("RandomDigitNot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomElement(t *testing.T) {

	tests := []struct {
		name string
		args []string
	}{
		// TODO: Add test cases.
		{name: "1", args: []string{}},
		{name: "2", args: []string{"a", "b", "c", "d"}},
		{name: "3", args: []string{"a", "b", "c", "d"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomElement(tt.args)
			t.Log(got)
		})
	}

}
func TestRandomElements(t *testing.T) {

	tests := []struct {
		name  string
		args  []string
		count int
	}{
		// TODO: Add test cases.
		{name: "3", args: []string{"a", "b", "c", "d"}, count: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomElements(tt.args, tt.count)
			t.Log(got)
		})
	}

}

func TestDdl(t *testing.T) {
	res := make([]string, 10, 10)

	ok := res[8]
	t.Log(ok)
}

func TestNumerify(t *testing.T) {
	type args struct {
		format []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "1", args: args{}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Numerify(tt.args.format...); got != tt.want {
				t.Errorf("Numerify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexify(t *testing.T) {
	type args struct {
		format []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "1", args: args{}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lexify(tt.args.format...); got != tt.want {
				t.Errorf("Lexify() = %v, want %v", got, tt.want)
			}
		})
	}
}
