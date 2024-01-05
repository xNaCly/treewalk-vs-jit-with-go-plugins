package eval

import (
	"fmt"
	"interpreter/ast"
	"interpreter/tokens"
)

type Operator uint8

const registerCount uint8 = 255

const (
	NOP      Operator = iota
	ADD               // add value in supplied register to value in register 0
	SUBTRACT          // subtract value supplied register from value in register 0
	DIVIDE            // divide value in register 0 with value in suplied register
	MULTIPLY          // multiply value in register 0 with value in suplied register
	NEGATE            // negate value in register 0
	STORE             // move value from register 0 to any register
	LOAD              // move value from constant pool to register 0
)

var operatorLookup = map[Operator]string{
	NOP:      "NOP",
	ADD:      "ADD",
	SUBTRACT: "SUBTRACT",
	DIVIDE:   "DIVIDE",
	MULTIPLY: "MULTIPLY",
	NEGATE:   "NEGATE",
	STORE:    "STORE",
	LOAD:     "LOAD",
}

type Instruction struct {
	Operator Operator
	Argument uint8
}

// all operator results are assumed to be stored in register 0
type Vm struct {
	constantPool []float64
	Registers    [registerCount]float64
	alloc        allocator
}

func (vm *Vm) makeConstant(value float64) uint8 {
	vm.constantPool = append(vm.constantPool, value)
	return uint8(len(vm.constantPool) - 1)
}

func (vm *Vm) Eval(instructions []Instruction) float64 {
	for _, instruction := range instructions {
		switch instruction.Operator {
		case NOP:
			continue
		case ADD:
			vm.Registers[0] += vm.Registers[instruction.Argument]
		case SUBTRACT:
			vm.Registers[0] -= vm.Registers[instruction.Argument]
		case DIVIDE:
			vm.Registers[0] /= vm.Registers[instruction.Argument]
		case MULTIPLY:
			vm.Registers[0] *= vm.Registers[instruction.Argument]
		case NEGATE:
			vm.Registers[0] *= -1
		case STORE:
			vm.Registers[instruction.Argument] = vm.Registers[0]
			vm.Registers[0] = 0
		case LOAD:
			vm.Registers[0] = vm.constantPool[instruction.Argument]
		}
	}
	return vm.Registers[0]
}

// simulates the vm structure for compile time correctness
type allocator struct {
	registers [registerCount]bool
}

func (a *allocator) alloc() uint8 {
	for i, v := range a.registers {
		if v == false {
			a.registers[i] = true
			return uint8(i + 1)
		}
	}
	panic("Out of free registers")
}

func (a *allocator) dealloc(index uint8) {
	if a.registers[index-1] {
		a.registers[index-1] = false
	} else {
		panic("Deallocating a not previously used register")
	}
}

func InspectBytecode(instructions []Instruction) []string {
	r := make([]string, len(instructions))
	for i, instruction := range instructions {
		r[i] = fmt.Sprintf("[%s:%d]", operatorLookup[instruction.Operator], instruction.Argument)
	}
	return r
}

func BytecodeCompile(vm *Vm, node ast.Node) []Instruction {

	r := make([]Instruction, 0)
	switch node := node.(type) {
	case *ast.Binary:
		r = append(r, BytecodeCompile(vm, node.Right)...)
		register := vm.alloc.alloc()
		r = append(r, Instruction{Operator: STORE, Argument: register})
		r = append(r, BytecodeCompile(vm, node.Left)...)
		switch node.Token.Type {
		case tokens.PLUS:
			r = append(r, Instruction{Operator: ADD, Argument: register})
		case tokens.MINUS:
			r = append(r, Instruction{Operator: SUBTRACT, Argument: register})
		case tokens.ASTERIKS:
			r = append(r, Instruction{Operator: MULTIPLY, Argument: register})
		case tokens.SLASH:
			r = append(r, Instruction{Operator: DIVIDE, Argument: register})
		}
		vm.alloc.dealloc(register)
	case *ast.Unary:
		r = append(r, BytecodeCompile(vm, node.Right)...)
		r = append(r, Instruction{Operator: NEGATE})
	case *ast.Number:
		r = append(r, Instruction{Operator: LOAD, Argument: vm.makeConstant(node.Value)})
	}
	return r
}
