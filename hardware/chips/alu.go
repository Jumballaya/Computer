package chips

// HalfAdder adds two bits together returning the sum and carry
func HalfAdder(a, b int) (int, int) {
	sum := Xor(a, b)
	carry := And(a, b)

	return sum, carry
}

// FullAdder adds three bits together returning the sum and carry
func FullAdder(a, b, c int) (int, int) {
	s1, c1 := HalfAdder(a, b)
	sum, c2 := HalfAdder(s1, c)

	carry := Or(c1, c2)

	return sum, carry
}

// Add16 adds 2 16-bit numbers
func Add16(a, b [16]int) [16]int {
	var out [16]int

	s1, c1 := HalfAdder(a[15], b[15])
	s2, c2 := FullAdder(a[14], b[14], c1)
	s3, c3 := FullAdder(a[13], b[13], c2)
	s4, c4 := FullAdder(a[12], b[12], c3)
	s5, c5 := FullAdder(a[11], b[11], c4)
	s6, c6 := FullAdder(a[10], b[10], c5)
	s7, c7 := FullAdder(a[9], b[9], c6)
	s8, c8 := FullAdder(a[8], b[8], c7)
	s9, c9 := FullAdder(a[7], b[7], c8)
	s10, c10 := FullAdder(a[6], b[6], c9)
	s11, c11 := FullAdder(a[5], b[5], c10)
	s12, c12 := FullAdder(a[4], b[4], c11)
	s13, c13 := FullAdder(a[3], b[3], c12)
	s14, c14 := FullAdder(a[2], b[2], c13)
	s15, c15 := FullAdder(a[1], b[1], c14)
	s16, _ := FullAdder(a[0], b[0], c15)

	out[15] = s1
	out[14] = s2
	out[13] = s3
	out[12] = s4
	out[11] = s5
	out[10] = s6
	out[9] = s7
	out[8] = s8
	out[7] = s9
	out[6] = s10
	out[5] = s11
	out[4] = s12
	out[3] = s13
	out[2] = s14
	out[1] = s15
	out[0] = s16

	return out
}

// Inc16 increments a 16-bit value by 1
func Inc16(in [16]int) [16]int {
	one := [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	return Add16(in, one)
}

/**
 * The ALU (Arithmetic Logic Unit).
 * Computes one of the following functions:
 * x+y, x-y, y-x, 0, 1, -1, x, y, -x, -y, !x, !y,
 * x+1, y+1, x-1, y-1, x&y, x|y on two 16-bit inputs,
 * according to 6 input bits denoted zx,nx,zy,ny,f,no.
 * In addition, the ALU computes two 1-bit outputs:
 * if the ALU output == 0, zr is set to 1; otherwise zr is set to 0;
 * if the ALU output < 0, ng is set to 1; otherwise ng is set to 0.
 */

/**
 * Implementation: the ALU logic manipulates the x and y inputs
 * and operates on the resulting values, as follows:
 * if (zx == 1) set x = 0        // 16-bit constant
 * if (nx == 1) set x = !x       // bitwise not
 * if (zy == 1) set y = 0        // 16-bit constant
 * if (ny == 1) set y = !y       // bitwise not
 * if (f == 1)  set out = x + y  // integer 2's complement addition
 * if (f == 0)  set out = x & y  // bitwise and
 * if (no == 1) set out = !out   // bitwise not
 * if (out == 0) set zr = 1
 * if (out < 0) set ng = 1
 */

// ALU computes mathematical functions for the CPU
func ALU(x, y [16]int, zx, nx, zy, ny, f, no int) ([16]int, int, int) {
	x1 := Mux16(x, FALSE, zx)
	notx := Not16(x1)
	x2 := Mux16(x1, notx, nx)

	y1 := Mux16(y, FALSE, zy)
	noty := Not16(y1)
	y2 := Mux16(y1, noty, ny)

	addxy := Add16(x2, y2)
	andxy := And16(x2, y2)

	posf := Mux16(andxy, addxy, f)
	negf := Not16(posf)

	out := Mux16(posf, negf, no)

	outlow := [8]int{out[7], out[6], out[5], out[4], out[3], out[2], out[1], out[0]}
	outhi := [8]int{out[15], out[14], out[13], out[12], out[11], out[10], out[9], out[8]}

	lowor := Or8Way(outlow)
	hior := Or8Way(outhi)

	nzr := Or(lowor, hior)

	zr := Not(nzr)
	ng := out[0]

	return out, zr, ng
}
