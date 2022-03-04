package cpu

// HERE IS WHERE ALL THE OPCODE FUNCTIONS WILL GO AS WELL AS THE MAP MAYBE
import (
	"fmt"
)

// HERE IS WHERE ALL THE OPCODE FUNCTIONS WILL GO AS WELL AS THE MAP MAYBE

type OpcodeFunc func()

type OpcodeAddrMode func()

type Instruction struct {
	Name string
	//OpcodeFunc     OpcodeFunc     // we're going to see if this works since we will be pointing to the struct most likely
	//OpcodeAddrMode OpcodeAddrMode // double pointers all the way across the sky!
	cycles uint8
}

func Brk(fp OpcodeFunc) uint8 {
	var r Registers
	fmt.Println(r.X)
	return r.X
}

func Imm() {
	var count = 0
	count++
}

func Bpl() {
	var count = 0
	count++
}

func Rel() {
	var count = 0
	count++
}

var BRK = Instruction{"BRK", 7}
var BPL = Instruction{"BPL", 2}

func CreateInstructionMap() map[string]*Instruction {
	opcodes := make(map[string]*Instruction)

	opcodes["BRK"] = &BRK
	opcodes["BPL"] = &BPL

	fmt.Println(BRK)
	return opcodes
}

func main() {
	CreateInstructionMap()
}
