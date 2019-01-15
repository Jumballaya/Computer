package chips

// DataFlipFlop is the basic clock flip-flop structure
type DataFlipFlop struct {
	previous int
}

// Tick sets the DFF to LOW
func (dff *DataFlipFlop) Tick() {
	dff.previous = 0
}

// Tock sets the DFF to HIGH
func (dff *DataFlipFlop) Tock() {
	dff.previous = 1
}

// Next sets the previous based on the incoming bit
func (dff *DataFlipFlop) Next(in int) int {
	// do something with in
	// then return the previous
	return dff.previous
}

// Bit is the representation of a single bit in memory
func Bit(dff *DataFlipFlop, in, load int) int {
	return 0
}

// Register is a 16-bit stored number
func Register(dff *DataFlipFlop, in [16]int, load int) [16]int {
	var out [16]int

	return out
}

// RAM8 is 8 bits of memory (3 address bits)
func RAM8(dff *DataFlipFlop, in [16]int, load int, address [3]int) [16]int {
	var out [16]int

	return out
}

// RAM64 is 64 bits of memory (6 address bits)
func RAM64(dff *DataFlipFlop, in [16]int, load int, address [6]int) [16]int {
	var out [16]int

	return out
}

// RAM512 is 512 bits of memory (9 address bits)
func RAM512(dff *DataFlipFlop, in [16]int, load int, address [9]int) [16]int {
	var out [16]int

	return out
}

// RAM4k is 4 kilobytes of memory (12 address bits)
func RAM4K(dff *DataFlipFlop, in [16]int, load int, address [12]int) [16]int {
	var out [16]int

	return out
}

// RAM16K is 16 kilobytes of memory (14 address bits)
func RAM16K(dff *DataFlipFlop, in [16]int, load int, address [14]int) [16]int {
	var out [16]int

	return out
}

// PC is the Program Counter that keeps track of what instruction to fetch
// reset sets the PC to 0
func PC(in [16]int, load, inc, reset int) [16]int {
	var out [16]int

	return out
}
