package chips

var FALSE [16]int = [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// Nand takes in two bits. The two bits are anded and then the not of that result is returned.
func Nand(a, b int) int {
	if (a == 1) && (b == 1) {
		return 0
	}
	return 1
}

// Not takes a bit and returns the opposite state of that bit
func Not(in int) int {
	return Nand(in, in)
}

// And takes two bits and checks to see if they are both 1
func And(a, b int) int {
	return Not(Nand(a, b))
}

// Or takes two bits and returns 1 if either of the bits are 1
func Or(a, b int) int {
	return Not(And(Not(a), Not(b)))
}

// Xor, or exclusive-or, returns 1 if the sum of the bits is odd.
func Xor(a, b int) int {
	return Or(And(a, Not(b)), And(Not(a), b))
}

// Mux (multiplexer) is similar to an if statement, reads like the following: if sel == 1 then b else a
func Mux(a, b, sel int) int {
	nsel := Not(sel)
	c1 := And(sel, b)
	c2 := And(nsel, a)
	return Or(c1, c2)
}

// Dmux is an inverted Mux, and reads like the following:
// if sel == 1 then (a = 0, b = 1) else (a = 1, b = 0)
func Dmux(in, sel int) (int, int) {
	nsel := Not(sel)
	a := And(nsel, in)
	b := And(sel, in)
	return a, b
}

// Or8Way applies an Or to the bits of an 8-bit value
func Or8Way(in [8]int) int {
	c1 := Or(in[0], in[1])
	c2 := Or(in[2], in[3])
	c3 := Or(in[4], in[5])
	c4 := Or(in[6], in[7])

	c5 := Or(c1, c2)
	c6 := Or(c3, c4)

	return Or(c5, c6)
}

// Dmux4Way takes in one bit and a 2-bit select and returns four bits with the in bit in the correct spot
func Dmux4Way(in int, sel [2]int) (int, int, int, int) {
	c1, c2 := Dmux(in, sel[0])
	a, b := Dmux(c1, sel[1])
	c, d := Dmux(c2, sel[1])

	return a, b, c, d
}

// Dmux8Way is just like Dmux4Way, with 8-bits out and 3-bits of select
func Dmux8Way(in int, sel [3]int) (int, int, int, int, int, int, int, int) {
	c1, c2 := Dmux(in, sel[0])

	c3, c4 := Dmux(c1, sel[1])
	c5, c6 := Dmux(c2, sel[1])

	a, b := Dmux(c3, sel[2])
	c, d := Dmux(c4, sel[2])
	e, f := Dmux(c5, sel[2])
	g, h := Dmux(c6, sel[2])

	return a, b, c, d, e, f, g, h
}
