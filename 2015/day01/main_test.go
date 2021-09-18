package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		instructions string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{"(())"},
			want: 0,
		},
		{
			name: "Test 2",
			args: args{"()()"},
			want: 0,
		},
		{
			name: "Test 3",
			args: args{"((("},
			want: 3,
		},
		{
			name: "Test 4",
			args: args{"(()(()("},
			want: 3,
		},
		{
			name: "Test 5",
			args: args{"))((((("},
			want: 3,
		},
		{
			name: "Test 6",
			args: args{"())"},
			want: -1,
		},
		{
			name: "Test 7",
			args: args{"))("},
			want: -1,
		},
		{
			name: "Test 8",
			args: args{")))"},
			want: -3,
		},
		{
			name: "Test 9",
			args: args{")())())"},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.instructions); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		instructions string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Test 1",
			args:    args{")"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "Test 2",
			args:    args{"()())"},
			want:    5,
			wantErr: false,
		},
		{
			name:    "Test 3",
			args:    args{"((((((("},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.args.instructions)
			if (err != nil) != tt.wantErr {
				t.Errorf("part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("part2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
