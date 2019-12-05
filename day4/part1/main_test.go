package main

import "testing"

func Test_getDigitAtPosition(t *testing.T) {
	type args struct {
		input    int
		position int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base",
			args: args{
				input:    654321,
				position: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDigitAtPosition(tt.args.input, tt.args.position); got != tt.want {
				t.Errorf("getDigitAtPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
