package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		dimensions []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[]string{"2x3x4"}},
			want: 58,
		},
		{
			name: "Test 2",
			args: args{[]string{"1x1x10"}},
			want: 43,
		},
		{
			name: "Test 3",
			args: args{[]string{"2x3x4", "1x1x10"}},
			want: 101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.dimensions); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		dimensions []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[]string{"2x3x4"}},
			want: 34,
		},
		{
			name: "Test 2",
			args: args{[]string{"1x1x10"}},
			want: 14,
		},
		{
			name: "Test 3",
			args: args{[]string{"2x3x4", "1x1x10"}},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.dimensions); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
