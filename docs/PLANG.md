# P Language

High-level C-like programming language for the computer. Plang is modeled after C, C++ and Go


## Ideas --

Keywords
  * `import`
  * `from`
  * `int`, `int*`
  * `float`, `float*`
  * `char`, `char*`
  * `bool`, `bool*`
  * `true`
  * `false`
  * `fn`
  * `struct`
  * `this`
  * `if`
  * `elif`
  * `else`
  * `for`
  * `break`
  * `case`
  * `switch`
  * `default`


- Comments:
  * `/* */` - multi-line
  * `//` - single-line
- Imports
  * `import math`
  * `import './local_file.p'`
  * `import PI from math`
  * `import object from './local_file.p'`
- Types
  * Numbers
    - `int`
    - `float`
  * Chars
    - `char`
    - `[]char`
    - `char*`
  * Booleans
    - `true`
    - `false`
    - Truthy: `int/float > 0`
    - Falsey: `int/floay <= 0`
  * Arrays
    - `[]int`
    - `[5]char`
  * Pointers
    - `char* foo = "bar"`
    - `[]char foo = ['c', 'a', 't']` `*foo == 'c'`
    - `char foo` `fill_word(&foo)`
  * Structs
    - `type Point struct { int x, int y }` `Point p = new Point(10, 8)`
    - Struct methods `fn Point::incX(int i) void { this.x += i }`
    - Struct constructor methods `fn Point::Point(int x, int y) void { this. x = x; this.y = y }`
  * First class functions
    - Function uses the `fn` keyword
    - Function signatures `fn add(int x, int y) int { return x + y }` turns into `(int, int) -> int`
    - This will be a little tough, keep it as simple as possible
    - Look in examples for using functions as parameters
    - Closures
  * For Loops and recursion
    - Basic - `for() { fmt.println("infinite loop") }`
    - Boolean - `for(bool running = true) {}` as long as the expression in the parenthesis is bool or truthy/falsey
    - Traditional - `for (int i = 0; i < 10; i++) { fmt.println("%d", i) }`
  * If statements
    - `if`, `elif`, `else`
  * Switch/Case
    - `switch (foo) { case 0: return 10 }`
  * Module system
    - Standard library (for the target computer)
    - `import PI from math`
    - `import prompt from io`
    - `import foo from './bar.p'`
  * Boolean logic
    - `>`, `<`, `>=`, `<=`, `||`, `&&`
  * Bitwise operations
