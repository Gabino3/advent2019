package main

import "testing"

func Test_hasDoubleDigits(t *testing.T) {

	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "1",
			input: 654321,
			want:  false,
		},
		{
			name:  "2",
			input: 111234,
			want:  false,
		},
		{
			name:  "3",
			input: 111223,
			want:  true,
		},
		{
			name:  "4",
			input: 123345,
			want:  true,
		},
		{
			name:  "5",
			input: 123455,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasDoubleDigits(tt.input); got != tt.want {
				t.Errorf("hasDoubleDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
