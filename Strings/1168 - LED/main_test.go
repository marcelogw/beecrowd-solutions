package main

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return correct 0 for incorrect input",
			args: args{
				input: "invalid",
			},
			want: 0,
		},
		{
			name: "should return correct number for 115380",
			args: args{
				input: "115380",
			},
			want: 27,
		},
		{
			name: "should return correct number for 2819311",
			args: args{
				input: "2819311",
			},
			want: 29,
		},
		{
			name: "should return correct number for 1234567890",
			args: args{
				input: "1234567890",
			},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.input); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
