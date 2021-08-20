package geo

import (
	"reflect"
	"testing"
)

func TestLineFrom(t *testing.T) {
	type args struct {
		m Point
		n Point
	}
	tests := []struct {
		args args
		want Line
	}{
		{args: args{m: Point{0, 0}, n: Point{1, 1}}, want: Line{a: -1, b: 1, c: 0}},
		{args: args{m: Point{1, 2}, n: Point{3, -1}}, want: Line{a: 3, b: 2, c: -7}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NewLine(tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineFrom() = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}
