package chips

import "testing"

func TestHalfAdder(t *testing.T) {
	tests := []struct {
		a  int
		b  int
		ex [2]int
	}{
		{0, 0, [2]int{0, 0}},
		{0, 1, [2]int{1, 0}},
		{1, 0, [2]int{1, 0}},
		{1, 1, [2]int{0, 1}},
	}

	for _, tt := range tests {
		sum, carry := HalfAdder(tt.a, tt.b)
		got := [2]int{sum, carry}
		if tt.ex != got {
			t.Errorf("HalfAdder gate failed, expected: %v, got: %v", tt.ex, got)
		}
	}
}

func TestFullAdder(t *testing.T) {
	tests := []struct {
		a  int
		b  int
		c  int
		ex [2]int
	}{
		{0, 0, 0, [2]int{0, 0}},
		{0, 0, 1, [2]int{1, 0}},
		{0, 1, 0, [2]int{1, 0}},
		{0, 1, 1, [2]int{0, 1}},

		{1, 0, 0, [2]int{1, 0}},
		{1, 0, 1, [2]int{0, 1}},
		{1, 1, 0, [2]int{0, 1}},
		{1, 1, 1, [2]int{1, 1}},
	}

	for _, tt := range tests {
		sum, carry := FullAdder(tt.a, tt.b, tt.c)
		got := [2]int{sum, carry}
		if tt.ex != got {
			t.Errorf("HalfAdder gate failed, expected: %v, got: %v", tt.ex, got)
		}
	}
}

func TestAdd16(t *testing.T) {
	tests := []struct {
		a  [16]int
		b  [16]int
		ex [16]int
	}{
		{
			[16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			[16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			[16]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			[16]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			[16]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			[16]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		},
		{
			[16]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
			[16]int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			[16]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			[16]int{0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1},
			[16]int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
			[16]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1},
		},
		{
			[16]int{0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1, 0, 1, 0, 0},
			[16]int{1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0},
			[16]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		},
	}

	for _, tt := range tests {
		got := Add16(tt.a, tt.b)
		if tt.ex != got {
			t.Errorf("Add16 gate failed, expected: %v, got: %v", tt.ex, got)
		}
	}
}

func TestInc16(t *testing.T) {
	tests := []struct {
		in [16]int
		ex [16]int
	}{
		{
			[16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			[16]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			[16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			[16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
			[16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0},
		},
		{
			[16]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1},
			[16]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		},
	}

	for _, tt := range tests {
		got := Inc16(tt.in)
		if tt.ex != got {
			t.Errorf("Inc16 gate failed, expected: %v, got: %v", tt.ex, got)
		}
	}
}

type aluExpected struct {
	out [16]int
	zr  int
	ng  int
}

func TestALU(t *testing.T) {
	tests := []struct {
		x [16]int
		y [16]int

		zx int
		nx int
		zy int
		ny int
		f  int
		no int

		ex aluExpected
	}{
		{
			[16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			1, 0, 1, 0, 1, 0,
			aluExpected{
				[16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				1, 0,
			},
		},
		{
			[16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			1, 1, 1, 1, 1, 1,
			aluExpected{
				[16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
				0, 0,
			},
		},
	}

	for _, tt := range tests {
		out, zr, ng := ALU(tt.x, tt.y, tt.zx, tt.nx, tt.zy, tt.ny, tt.f, tt.no)
		got := aluExpected{out, zr, ng}
		if tt.ex != got {
			t.Errorf("Inc16 gate failed, expected: %v, got: %v", tt.ex, got)
		}
	}
}
