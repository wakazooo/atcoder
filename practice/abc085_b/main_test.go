package main

import (
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
	type tArgs struct {
		args  *args
		input int
	}
	tests := []struct {
		name string
		args tArgs
		want *args
	}{
		{
			name: "middle insert pattern",
			args: tArgs{
				args: &args{
					arr: []int{3, 1},
				},
				input: 2,
			},
			want: &args{
				arr: []int{3, 2, 1},
			},
		},
		{
			name: "first insert pattern",
			args: tArgs{
				args: &args{
					arr: []int{3, 1},
				},
				input: 4,
			},
			want: &args{
				arr: []int{4, 3, 1},
			},
		},
		{
			name: "last insert pattern",
			args: tArgs{
				args: &args{
					arr: []int{4, 3},
				},
				input: 2,
			},
			want: &args{
				arr: []int{4, 3, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort(tt.args.args, tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
