// stack pointer vs base pointer

package main

import "fmt"

func add(x int, y int) int {
	var result int
	result = x + y
	return result
}

func main() {
	var a int = 10
	var sum = add(a, 4)
	fmt.Println("The sum is:", sum)
}

/*
1. Prothome amar pc run/on korlam. Run/On korar por operating system er code RAM e load holo.
2. Tarpor Control Unit Program Counter k bolbe -> prothom j address ache oitak point kor. Suppose 
first address is 0. Then 0 Program Counter e load holo.
3. Tarpor Control Unit 0 address theke instruction fetch kore ene Instruction Register e ene rakhbe.
4. Fetch kore rakhar por Control Unit Instruction Register theke read korbe. Read kore instruction 
decode korbe.
Suppose decode kore pailo 3, 5 multiplication kora lagbe
5. 3 k niye rakhlo Accumulator e, 5 k niye rakhlo Base e, Multiplication sign k niye rakhlo Counter e. (It's not obvious just to understand the concept.)
6. Erpor Control Unit Arithmetic Logic Unit (ALU) k bolbe eigula Multiplication koro.

#7. Tarpor Program Counter er value 1 bere jabe. And Repeat (1 -6).

For 8-bit, registers are - 
AL, BL, CL, DL

For 16-bit, registers are - 
AX, BX, CX, DX

For 32-bit, registers are - 
EAX, EBX, ECX, EDX

For 64-bit, registers are -
RAX, RBX, RCX, RDX

* A for - Accumulator
* B for - Base
* C for - Counter
* D for - Data


### SP value is always less than BP
*/

/*
Ekhon ami ei go file ta run korbo. Run korar jonno RAM e ekta Process create hbe. Tarpor OS tar jaiga 
theke shore ekhon process k Point korbe. Let's start the procedure:
1. Suppose process er memory start hoy 24 address theke. Tahole Program counter ekhon point korbe Address
24 k.
2. So compilation phase e main & add ei 2 ta function code segment e giye load hbe. 
3. Stack segment e function er jonno stack frame create hoy last er dik theke.
.
.
.
.

*/