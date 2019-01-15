package chips

// Not16 applies Not to a 16-bit value
func Not16(in [16]int) [16]int {
	var out [16]int

	for i, n := range in {
		out[i] = Not(n)
	}

	return out
}

// Add16 applies an And to two 16-bit values
func And16(a, b [16]int) [16]int {
	var out [16]int

	for i, _ := range a {
		out[i] = And(a[i], b[i])
	}

	return out
}

// Or16 applies an Or to two 16-bite values
func Or16(a, b [16]int) [16]int {
	var out [16]int

	for i, _ := range a {
		out[i] = Or(a[i], b[i])
	}

	return out
}

// Mux16 applies a Mux to two 16-bit values
func Mux16(a, b [16]int, sel int) [16]int {
	var out [16]int

	for i, _ := range a {
		out[i] = Mux(a[i], b[i], sel)
	}

	return out
}

// Mux4Way16 applies a 4-way mux with 16-bit values
func Mux4Way16(a, b, c, d [16]int, sel [2]int) [16]int {
	c1 := Mux16(a, b, sel[1])
	c2 := Mux16(c, d, sel[1])

	return Mux16(c1, c2, sel[0])
}

// Mux8Way16 applies a 8-way mux with 16-bit values
func Mux8Way16(a, b, c, d, e, f, g, h [16]int, sel [3]int) [16]int {
	c1 := Mux16(a, b, sel[2])
	c2 := Mux16(c, d, sel[2])
	c3 := Mux16(e, f, sel[2])
	c4 := Mux16(g, h, sel[2])

	c5 := Mux16(c1, c2, sel[1])
	c6 := Mux16(c3, c4, sel[1])

	return Mux16(c5, c6, sel[0])
}
