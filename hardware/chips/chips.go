package chips

type RAM struct{}

type ROM struct{}

// Memory is a basic implementation of a Memory structure
func Memory(in, address [16]int, load int) [16]int {
	var out [16]int

	return out
}

// CPU fetches the current instruction and executes it
func CPU(inM, instruction [16]int, reset int) (outM, addressM, pc [16]int, writeM int) {
	outM = [16]int{}
	addressM = [16]int{}
	pc = [16]int{}
	writeM = 0

	return outM, addressM, pc, writeM
}
