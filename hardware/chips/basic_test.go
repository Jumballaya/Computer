package chips

import "testing"

func TestAnd(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	for _, tt := range tests {
		got := And(tt.a, tt.b)
		if got != tt.expected {
			t.Errorf("And gate failed, expected: %d, got: %d", tt.expected, got)
		}
	}
}

func TestOr(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, 1},
	}

	for _, tt := range tests {
		got := Or(tt.a, tt.b)
		if got != tt.expected {
			t.Errorf("Or gate failed, expected: %d, got: %d", tt.expected, got)
		}
	}
}

func TestNand(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{0, 0, 1},
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, 0},
	}

	for _, tt := range tests {
		got := Nand(tt.a, tt.b)
		if got != tt.expected {
			t.Errorf("Or gate failed, expected: %d, got: %d", tt.expected, got)
		}
	}
}

func TestNot(t *testing.T) {
	tests := []struct {
		in       int
		expected int
	}{
		{0, 1},
		{1, 0},
	}

	for _, tt := range tests {
		got := Not(tt.in)
		if got != tt.expected {
			t.Errorf("Or gate failed, expected: %d, got: %d", tt.expected, got)
		}
	}
}

func TestXor(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, 0},
	}

	for _, tt := range tests {
		got := Xor(tt.a, tt.b)
		if got != tt.expected {
			t.Errorf("Xor gate failed, expected: %d, got: %d", tt.expected, got)
		}
	}
}

func TestMux(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		sel      int
		expected int
	}{
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 1, 0, 0},
		{0, 1, 1, 1},
		{1, 0, 0, 1},
		{1, 0, 1, 0},
		{1, 1, 0, 1},
		{1, 1, 1, 1},
	}

	for _, tt := range tests {
		got := Mux(tt.a, tt.b, tt.sel)
		if got != tt.expected {
			t.Errorf("Or gate failed, expected: %d, got: %d", tt.expected, got)
		}
	}
}

func TestDmux(t *testing.T) {
	tests := []struct {
		in  int
		sel int
		a   int
		b   int
	}{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 1},
	}

	for _, tt := range tests {
		a, b := Dmux(tt.in, tt.sel)
		if !((a == tt.a) && (b == tt.b)) {
			t.Errorf("Dmux gate failed, expected: (%d, %d), got: (%d, %d)", tt.a, tt.b, a, b)
		}
	}
}

func TestOr8Way(t *testing.T) {
	tests := []struct {
		in       [8]int
		expected int
	}{
		{
			[8]int{0, 0, 0, 0, 0, 0, 0, 0},
			0,
		},
		{
			[8]int{1, 1, 1, 1, 1, 1, 1, 1},
			1,
		},
		{
			[8]int{0, 0, 0, 1, 0, 0, 0, 0},
			1,
		},
		{
			[8]int{0, 0, 0, 0, 0, 0, 0, 1},
			1,
		},
		{
			[8]int{0, 0, 1, 0, 0, 1, 1, 0},
			1,
		},
	}

	for _, tt := range tests {
		got := Or8Way(tt.in)
		if got != tt.expected {
			t.Errorf("Or8Way gate failed, expected: %d, got: %d", tt.expected, got)
		}
	}
}

func TestDmux4Way(t *testing.T) {
	tests := []struct {
		in  int
		sel [2]int
		ex  [4]int
	}{
		{0, [2]int{0, 0}, [4]int{0, 0, 0, 0}},
		{0, [2]int{0, 1}, [4]int{0, 0, 0, 0}},
		{0, [2]int{1, 0}, [4]int{0, 0, 0, 0}},
		{0, [2]int{1, 1}, [4]int{0, 0, 0, 0}},
		{1, [2]int{0, 0}, [4]int{1, 0, 0, 0}},
		{1, [2]int{0, 1}, [4]int{0, 1, 0, 0}},
		{1, [2]int{1, 0}, [4]int{0, 0, 1, 0}},
		{1, [2]int{1, 1}, [4]int{0, 0, 0, 1}},
	}

	for _, tt := range tests {
		a, b, c, d := Dmux4Way(tt.in, tt.sel)
		got := [4]int{a, b, c, d}
		if tt.ex != got {
			t.Errorf("Dmux4Way gate failed, expected: %v, got: %v", tt.ex, got)
		}
	}
}

func TestDmux8Way(t *testing.T) {
	tests := []struct {
		in  int
		sel [3]int
		ex  [8]int
	}{
		{0, [3]int{0, 0, 0}, [8]int{0, 0, 0, 0, 0, 0, 0, 0}},
		{0, [3]int{0, 0, 1}, [8]int{0, 0, 0, 0, 0, 0, 0, 0}},
		{0, [3]int{0, 1, 0}, [8]int{0, 0, 0, 0, 0, 0, 0, 0}},
		{0, [3]int{0, 1, 1}, [8]int{0, 0, 0, 0, 0, 0, 0, 0}},
		{0, [3]int{1, 0, 0}, [8]int{0, 0, 0, 0, 0, 0, 0, 0}},
		{0, [3]int{1, 0, 1}, [8]int{0, 0, 0, 0, 0, 0, 0, 0}},
		{0, [3]int{1, 1, 0}, [8]int{0, 0, 0, 0, 0, 0, 0, 0}},
		{0, [3]int{1, 1, 1}, [8]int{0, 0, 0, 0, 0, 0, 0, 0}},

		{1, [3]int{0, 0, 0}, [8]int{1, 0, 0, 0, 0, 0, 0, 0}},
		{1, [3]int{0, 0, 1}, [8]int{0, 1, 0, 0, 0, 0, 0, 0}},
		{1, [3]int{0, 1, 0}, [8]int{0, 0, 1, 0, 0, 0, 0, 0}},
		{1, [3]int{0, 1, 1}, [8]int{0, 0, 0, 1, 0, 0, 0, 0}},
		{1, [3]int{1, 0, 0}, [8]int{0, 0, 0, 0, 1, 0, 0, 0}},
		{1, [3]int{1, 0, 1}, [8]int{0, 0, 0, 0, 0, 1, 0, 0}},
		{1, [3]int{1, 1, 0}, [8]int{0, 0, 0, 0, 0, 0, 1, 0}},
		{1, [3]int{1, 1, 1}, [8]int{0, 0, 0, 0, 0, 0, 0, 1}},
	}

	for _, tt := range tests {
		a, b, c, d, e, f, g, h := Dmux8Way(tt.in, tt.sel)
		got := [8]int{a, b, c, d, e, f, g, h}
		if tt.ex != got {
			t.Errorf("Dmux4Way gate failed, expected: %v, got: %v", tt.ex, got)
		}
	}
}
