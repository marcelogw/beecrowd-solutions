package main

import "testing"

func Test_formatName(t *testing.T) {
	type args struct {
		line string
		year int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return correct values for name",
			args: args{
				line: "   marcelo   gaioso   werneck     ",
				year: 1992,
			},
			want: "mgw1992",
		},
		{
			name: "should return correct values for strange values",
			args: args{
				line: "   kkk   asd  ww e d asdqwe zxc   123  a",
				year: 2000,
			},
			want: "kawedaza2000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatName(tt.args.line, tt.args.year); got != tt.want {
				t.Errorf("formatName() = %v, want %v", got, tt.want)
			}
		})
	}
}