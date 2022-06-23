package geo

import (
	"math"
	"reflect"
	"testing"
)

func TestVector_Module(t *testing.T) {
	tests := []struct {
		fields Vector
		want   float64
	}{
		{Vector{x: 0, y: 0}, 0},
		{Vector{x: 0, y: 0}, 0},
		{Vector{x: 1, y: 1}, 1.4142135623730951},
		{Vector{x: 3, y: 4}, 5},
		{Vector{x: .3, y: .4}, .5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			v := &tt.fields
			if got := v.Module(); got != tt.want {
				t.Errorf("Vector.Module() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Mul(t *testing.T) {
	tests := []struct {
		fields Vector
		args   float64
		want   Vector
	}{
		{Vector{x: 0, y: 0}, 0, Vector{x: 0, y: 0}},
		{Vector{x: 1, y: 2}, 3, Vector{x: 3, y: 6}},
		{Vector{x: 1, y: 2}, 0.3, Vector{x: 0.3, y: 0.6}},
		{Vector{x: 0.1, y: 0.2}, 0.3, Vector{x: 0.03, y: 0.06}},
		{Vector{x: -1, y: 2}, 3, Vector{x: -3, y: 6}},
		{Vector{x: -1, y: 2}, 0.3, Vector{x: -0.3, y: 0.6}},
		{Vector{x: -0.1, y: 0.2}, 0.3, Vector{x: -0.03, y: 0.06}},
		{Vector{x: 1, y: -2}, 3, Vector{x: 3, y: -6}},
		{Vector{x: 1, y: -2}, 0.3, Vector{x: 0.3, y: -0.6}},
		{Vector{x: 0.1, y: -0.2}, 0.3, Vector{x: 0.03, y: -0.06}},
		{Vector{x: -1, y: -2}, 3, Vector{x: -3, y: -6}},
		{Vector{x: -1, y: -2}, 0.3, Vector{x: -0.3, y: -0.6}},
		{Vector{x: -0.1, y: -0.2}, 0.3, Vector{x: -0.03, y: -0.06}},
		{Vector{x: 1, y: 2}, -3, Vector{x: -3, y: -6}},
		{Vector{x: 1, y: 2}, -0.3, Vector{x: -0.3, y: -0.6}},
		{Vector{x: 0.1, y: 0.2}, -0.3, Vector{x: -0.03, y: -0.06}},
		{Vector{x: -1, y: 2}, -3, Vector{x: 3, y: -6}},
		{Vector{x: -1, y: 2}, -0.3, Vector{x: 0.3, y: -0.6}},
		{Vector{x: -0.1, y: 0.2}, -0.3, Vector{x: 0.03, y: -0.06}},
		{Vector{x: 1, y: -2}, -3, Vector{x: -3, y: 6}},
		{Vector{x: 1, y: -2}, -0.3, Vector{x: -0.3, y: 0.6}},
		{Vector{x: 0.1, y: -0.2}, -0.3, Vector{x: -0.03, y: 0.06}},
		{Vector{x: -1, y: -2}, -3, Vector{x: 3, y: 6}},
		{Vector{x: -1, y: -2}, -0.3, Vector{x: 0.3, y: 0.6}},
		{Vector{x: -0.1, y: -0.2}, -0.3, Vector{x: 0.03, y: 0.06}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			v := &tt.fields
			if got := v.Mul(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Div(t *testing.T) {
	tests := []struct {
		fields Vector
		args   float64
		want   Vector
	}{
		{Vector{x: 4, y: 2}, 2, Vector{x: 2, y: 1}},
		{Vector{x: 1, y: 1}, 0.5, Vector{x: 2, y: 2}},
		{Vector{x: 0.3333333333333333, y: float64(1) / float64(6)}, 0.3333333333333333, Vector{x: 1, y: 0.5}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			v := &tt.fields
			if got := v.Div(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Sub(t *testing.T) {
	tests := []struct {
		fields Vector
		args   Vector
		want   Vector
	}{
		{Vector{x: 1, y: 1}, Vector{x: 0, y: -2}, Vector{x: 1, y: 3}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			v := &tt.fields
			if got := v.Sub(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Add(t *testing.T) {
	tests := []struct {
		fields Vector
		args   Vector
		want   Vector
	}{
		{Vector{x: 1, y: 1}, Vector{x: 0, y: 2}, Vector{x: 1, y: 3}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			v := &tt.fields
			if got := v.Add(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Alpha(t *testing.T) {
	tests := []struct {
		fields Vector
		want   Angle
	}{
		{Vector{x: 1, y: 0}, Angle(0)},
		{Vector{x: 0, y: 1}, Angle(math.Pi / 2)},
		{Vector{x: -1, y: 0}, Angle(math.Pi)},
		{Vector{x: 0, y: -1}, Angle(-math.Pi / 2)},
		{Vector{x: 1, y: 1}, Angle(math.Pi / 4)},
		{Vector{x: -1, y: 1}, Angle(3 * math.Pi / 4)},
		{Vector{x: -1, y: -1}, Angle(-3 * math.Pi / 4)},
		{Vector{x: 1, y: -1}, Angle(-math.Pi / 4)},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			v := &tt.fields
			if got := v.Alpha(); got != tt.want {
				t.Errorf("Vector.Alpha() = %v, want %v", got, tt.want)
			}
		})
	}
}
