# Assembler Spec

The Hack Assembly Language is minimal, it mainly consists of 3 types of instructions. It ignores whitespace and allows programs to declare symbols with a single symbol declaration instruction. Symbols can either be labels or variables. It also allows the programmer to write comments in the source code, for example: `// this is a single line comment`.

### Predefined Symbols

  * **A**: Address Register.
  * **D**: Data Register.
  * **M**: Refers to the register in Main Memory whose address is currently stored in **A**.
  * **SP**: RAM address 0.
  * **LCL**: RAM address 1.
  * **ARG**: RAM address 2.
  * **THIS**: RAM address 3.
  * **THAT**: RAM address 4.
  * **R0**-**R15**: Addresses of 16 RAM Registers, mapped from 0 to 15.
  * **SCREEN**: Base address of the Screen Map in Main Memory, which is equal to 16384.
  * **KBD**: Keyboard Register address in Main Memory, which is equal to 24576.

### Types of Instructions:

  1. A-Instruction: Addressing instructions.
  2. C-Instruction: Computation instructions.
  3. L-Instruction: Labels (Symbols) declaration instructions.

#### A-INSTRUCTIONS:

##### Symbolic Syntax:

`@value`, where value is either a decimal non-negative number or a Symbol.

Examples:

  * `@21`
  * `@R0`
  * `@SCREEN`

##### Binary Syntax:

`0xxxxxxxxxxxxxxx`, where `x` is a bit, either 0 or 1. A-Instructions always have their MSB set to 0.

Examples:

  * `000000000001010`
  * `011111111111111`

##### Effects:

Sets the contents of the **A** register to the specified value. The value is either a non-negative number (i.e. 21) or a Symbol. If the value is a Symbol, then the contents of the **A** register is set to the value that the Symbol refers to but not the actual data in that Register or Memory Location.

#### L-INSTRUCTIONS:

Symbols can be either variables or lables. Variables are symbolic names for memory addresses to make remembering these addresses easier. Labels are instructions addresses that allow multiple jumps in the program easier to handle. Symbols declaration is not a machine instruction because machine code doesn't operate on the level of abstraction of that of labels and variables, and hence it is considered a pseudo-instruction.

##### Declaring Variables:

Declaring variables is a straight forward A-Instruction, example:

```
@i
M=0
```

The instruction `@i` declares a variable "i", and the instruction `M=0` sets the memory location of "i" in Main Memory to 0, the address "i" was automatically generated and stored in **A** Register by the instruction.

##### Declaring Labels:

To declare a label we need to use the command `(LABEL_NAME)`, where "LABEL_NAME" can be any name we desire to have for the label, as long as it's wraped between parentheses. For example:

```
(LOOP)
  // ...
  // instruction 1
  // instruction 2
  // instruction 3
  // ...
  @LOOP
  0;JMP
```

The instruction `(LOOP)` declares a new label called "LOOP", the assembler will resolve this label to the address of the next instruction (A or C instruction) on the following line.

The instruction `@LOOP` is a straight-forward A-Instruction that sets the contents of **A** Register to the instruction address the label refers to, whereas the `0;JMP` instruction causes an unconditional jump to the address in **A** Register causing the program to execute the set of instructions between `(LOOP)` and `0;JMP` infinitely.

#### C-INSTRUCTIONS:

##### Symbolic Syntax:

*dest* = *comp* ; *jmp*, where:

  1. *dest*: Destination register in which the result of computation will be stored.
  2. *comp*: Computation code.
  3. *jmp*: The jump directive.

Examples:
  * `D=0`
  * `M=1`
  * `D=D+1;JMP`
  * `M=M-D;JEQ`

##### Binary Syntax:

`1 1 1 a c1 c2 c3 c4 c5 c6 d1 d2 d3 j1 j2 j3`, where:

  * `111` bits: C-Instructions always begin with bits `111`.
  * `a` bit: Chooses to load the contents of either **A** register or **M** (Main Memory register addressed by **A**) into the ALU for computation.
  * Bits `c1` through `c6`: Control bits expected by the ALU to perform arithmetic or bit-wise logic operations.
  * Bits `d1` through `d3`: Specify which memory location to store the result of ALU computation into: **A**, **D** or **M**.
  * Bits `j1` through `j3`: Specify which JUMP directive to execute (either conditional or uncoditional).

##### Effects:

Performs a computation on the CPU (arithmetic or bit-wise logic) and stores it into a destination register or memory location, and then (optionally) JUMPS to an instruction memory location that is usually addressed by a value or a Symbol (label).

##### Reference:

![C-Instructions Reference](assets/c_instructions_reference.png "C-Instructions Reference")

*This C-Instruction reference image is taken from the [nand2tetris Coursera MOOC](https://www.coursera.org/learn/build-a-computer).*
