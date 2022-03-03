package cpu

import "fmt"

type Flag uint8

// Status Flags for the CPU
// bitmask iota enumerates thus:
// 0b001, 0b010, 0b100, etc.
const (
	C Flag = 1 << iota // sets Carry to 0b00000001
	Z                  // sets Zero to 0b00000010
	I                  // sets Interrupt Disable to 0b00000100
	D                  // sets Decimal Mode to 0b00001000
	B                  // sets Break to 0b00010000
	U                  // sets UNUSED to 0b00100000
	V                  // sets oVerflow Flag to 0b01000000
	N                  // sets Negative Flag to 0b10000000
)

func (cpu MOS6502) getFlag(f Flag) Flag {
	return cpu.Registers.P
}

func (cpu MOS6502) setFlag(f Flag, b bool) {
	if b {
		cpu.Registers.P = f
	} else {
		fmt.Sprintln("FIX ERROR HANDLING IN SET FLAG")
	}
}

type Registers struct {
	A  uint8 // Accumulator
	X  uint8
	Y  uint8
	P  Flag   // Processor Status Register
	SP uint8  // Stack Pointer
	PC uint16 // Program Counter
}

func InitializeRegisters() (r Registers) {
	r = Registers{}
	r.Reset()
	return
}

func (r *Registers) Reset() {
	r.A = 0
	r.X = 0
	r.Y = 0
	r.P = I | U
	r.SP = 0xFD
	r.PC = 0xFFFC
}

type Interrupt uint8

const (
	IRQ Interrupt = iota
	NMI
	RST
)

type MOS6502 struct {
	Registers   Registers
	decimalMode bool
	breakErr    bool
	Irq         bool
	Nmi         bool
	Rst         bool
	//Ram 	Ram
	//Opcodes Opcodes
}

type Instruction struct {
	Name           string
	OpcodeFunc     func()
	OpcodeAddrMode func()
	cycles         uint8
}

func (cpu MOS6502) Reset() {
	cpu.Registers.Reset()
	// cpu.Ram.Reset()
	// cpu.FullReset()
}
