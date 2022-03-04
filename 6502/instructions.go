package cpu

import (
	"fmt"

	"6502/cpu"
)

// HERE IS WHERE ALL THE OPCODE FUNCTIONS WILL GO AS WELL AS THE MAP MAYBE

type OpcodeFunc func()

type OpcodeAddrMode func()

type Instruction struct {
	Name           string
	OpcodeFunc     OpcodeFunc     // we're going to see if this works since we will be pointing to the struct most likely
	OpcodeAddrMode OpcodeAddrMode // double pointers all the way across the sky!
	cycles         uint8
}

func Brk(fp OpcodeFunc) {
	var count = 0
	count++
	fmt.Printf("%s", count)
	cpu.Registers.P
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

var BRK = Instruction{"BRK", Brk(), Imm(), 7}
var BPL = Instruction{"BPL", Bpl(), Rel(), 2}

func CreateInstructionMap() map[string]Instruction {
	opcodes := make(map[string]Instruction)

	opcodes["BRK"] = BRK{}
	opcodes["BPL"] = BPL{}

	return opcodes
}
